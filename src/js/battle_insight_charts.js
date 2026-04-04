/**
 * Линия жизни + интенсивность боя — синхронизированы с окном времени getTimeRangeJSON().
 */
import { Chart, registerables } from 'chart.js';
import { getRelativePosition } from 'chart.js/helpers';
import {
    commitZoomFromBrush,
    getTimeRangeBounds,
    getTimeRangeJSON,
    getTimelineCursorSec,
    setTimelineCursorSec,
    subscribeTimelineCursor,
} from './timeline.js';
import { deferAfterPaint, hideChartPreloader, showChartPreloader } from './chart_preloader.js';
import { gaEvent } from './analytics.js';

Chart.register(...registerables);

const battleSyncOverlayPlugin = {
    id: 'battleSyncOverlay',
    afterDraw(chart) {
        const xScale = chart.scales.x;
        if (!xScale || xScale.type !== 'linear') {
            return;
        }
        const { ctx, chartArea } = chart;
        if (!chartArea) {
            return;
        }
        const sec = getTimelineCursorSec();
        if (sec !== null && typeof sec === 'number' && Number.isFinite(sec)) {
            const x = xScale.getPixelForValue(sec);
            if (x >= chartArea.left - 1 && x <= chartArea.right + 1) {
                ctx.save();
                ctx.strokeStyle = 'rgba(248, 250, 252, 0.92)';
                ctx.lineWidth = 1;
                ctx.shadowColor = 'rgba(56, 189, 248, 0.45)';
                ctx.shadowBlur = 5;
                ctx.beginPath();
                ctx.moveTo(x, chartArea.top);
                ctx.lineTo(x, chartArea.bottom);
                ctx.stroke();
                ctx.restore();
            }
        }
    },
};

Chart.register(battleSyncOverlayPlugin);

let lifeChart = null;
let intensityChart = null;
let battleInsightLoadGen = 0;
let chartPointerUnsubs = [];
let cursorUnsub = null;

/** Пока тянем кисть на любом графике боя — не сбрасывать курсор по pointerleave. */
let battleChartBrushActive = false;

function ensureBattleChartBrushPreview(wrap) {
    let el = wrap.querySelector('.battle-chart-brush-preview');
    if (!el) {
        el = document.createElement('div');
        el.className = 'battle-chart-brush-preview';
        el.setAttribute('aria-hidden', 'true');
        wrap.insertBefore(el, wrap.firstChild);
    }
    return el;
}

function updateBattleChartBrushPreview(chart, wrap, previewEl, a, b) {
    const xScale = chart.scales.x;
    const ca = chart.chartArea;
    const canvas = chart.canvas;
    if (!xScale || !ca || !canvas || !wrap || !previewEl) {
        return;
    }
    const lo = Math.min(a, b);
    const hi = Math.max(a, b);
    const x1 = xScale.getPixelForValue(lo);
    const x2 = xScale.getPixelForValue(hi);
    const rect = canvas.getBoundingClientRect();
    const wrapRect = wrap.getBoundingClientRect();
    const sx = rect.width / chart.width;
    const sy = rect.height / chart.height;
    const px1 = x1 * sx;
    const px2 = x2 * sx;
    const left = rect.left - wrapRect.left + Math.min(px1, px2);
    const width = Math.max(Math.abs(px2 - px1), 0);
    const top = rect.top - wrapRect.top + ca.top * sy;
    const height = (ca.bottom - ca.top) * sy;
    previewEl.style.display = 'block';
    previewEl.style.left = `${left}px`;
    previewEl.style.width = `${width}px`;
    previewEl.style.top = `${top}px`;
    previewEl.style.height = `${height}px`;
}

function attachChartBrushAndCursor(chart, canvas, wrap) {
    const previewEl = ensureBattleChartBrushPreview(wrap);
    let docMove = null;
    let docUp = null;

    const onMove = (e) => {
        if (battleChartBrushActive) {
            return;
        }
        const pos = getRelativePosition(e, chart);
        const xScale = chart.scales.x;
        if (!xScale) {
            return;
        }
        const v = xScale.getValueForPixel(pos.x);
        if (typeof v === 'number' && Number.isFinite(v)) {
            setTimelineCursorSec(v);
        }
    };
    const onLeave = () => {
        if (!battleChartBrushActive) {
            setTimelineCursorSec(null);
        }
    };
    const onDown = (e) => {
        if (e.button !== 0) {
            return;
        }
        e.preventDefault();
        const xScale = chart.scales.x;
        if (!xScale) {
            return;
        }
        const pos0 = getRelativePosition(e, chart);
        const anchorSec = xScale.getValueForPixel(pos0.x);
        if (typeof anchorSec !== 'number' || !Number.isFinite(anchorSec)) {
            return;
        }
        let lastSec = anchorSec;
        battleChartBrushActive = true;
        updateBattleChartBrushPreview(chart, wrap, previewEl, anchorSec, anchorSec);

        docMove = (ev) => {
            const pos = getRelativePosition(ev, chart);
            const cur = xScale.getValueForPixel(pos.x);
            if (typeof cur === 'number' && Number.isFinite(cur)) {
                lastSec = cur;
                setTimelineCursorSec(cur);
            }
            updateBattleChartBrushPreview(chart, wrap, previewEl, anchorSec, lastSec);
        };
        docUp = () => {
            document.removeEventListener('pointermove', docMove);
            document.removeEventListener('pointerup', docUp);
            document.removeEventListener('pointercancel', docUp);
            docMove = null;
            docUp = null;
            previewEl.style.display = 'none';
            battleChartBrushActive = false;
            commitZoomFromBrush(anchorSec, lastSec);
        };
        document.addEventListener('pointermove', docMove);
        document.addEventListener('pointerup', docUp);
        document.addEventListener('pointercancel', docUp);
    };

    canvas.addEventListener('pointermove', onMove);
    canvas.addEventListener('pointerleave', onLeave);
    canvas.addEventListener('pointerdown', onDown);
    return () => {
        canvas.removeEventListener('pointermove', onMove);
        canvas.removeEventListener('pointerleave', onLeave);
        canvas.removeEventListener('pointerdown', onDown);
        if (docMove) {
            document.removeEventListener('pointermove', docMove);
        }
        if (docUp) {
            document.removeEventListener('pointerup', docUp);
            document.removeEventListener('pointercancel', docUp);
        }
        battleChartBrushActive = false;
    };
}

function chartDpr() {
    if (typeof window === 'undefined' || !window.devicePixelRatio) {
        return 1;
    }
    return Math.min(Math.max(window.devicePixelRatio, 1), 2.5);
}

function formatAxisSec(sec) {
    if (typeof sec !== 'number' || Number.isNaN(sec)) {
        return '';
    }
    const m = Math.floor(sec / 60);
    const s = sec - m * 60;
    const mm = String(m).padStart(2, '0');
    const ss = s < 10 ? `0${s.toFixed(1)}` : s.toFixed(1);
    return `${mm}:${ss}`;
}

/** Ширина оси Y = --timeline-plot-gutter-left в .match-timeline-section (метки времени ровно с таймлайном). */
function readTimelinePlotGutterLeftPx() {
    const el = document.querySelector('.match-timeline-section');
    if (!el) {
        return 80;
    }
    const raw = getComputedStyle(el).getPropertyValue('--timeline-plot-gutter-left').trim();
    const n = parseFloat(raw);
    return Number.isFinite(n) ? n : 80;
}

/** Домен оси X = окно таймлайна, иначе min/max по точкам. */
function xDomainSec(tArr) {
    const b = getTimeRangeBounds();
    if (
        b
        && typeof b.from === 'number'
        && typeof b.to === 'number'
        && b.to > b.from + 1e-9
    ) {
        return { min: b.from, max: b.to };
    }
    if (!tArr.length) {
        return { min: undefined, max: undefined };
    }
    return { min: Math.min(...tArr), max: Math.max(...tArr) };
}

const axisStyle = {
    grid: { color: 'rgba(63, 63, 70, 0.55)' },
    ticks: { color: '#a1a1aa', maxTicksLimit: 10 },
    border: { color: 'rgba(63, 63, 70, 0.6)' },
};

const commonLineOptions = {
    responsive: true,
    maintainAspectRatio: false,
    devicePixelRatio: chartDpr(),
    layout: {
        padding: {
            left: 0,
            right: 6,
            top: 4,
            bottom: 28,
        },
    },
    interaction: { mode: 'index', intersect: false },
    plugins: {
        legend: {
            display: false,
        },
        datalabels: {
            display: false,
        },
        tooltip: {
            backgroundColor: 'rgba(15, 23, 42, 0.94)',
            titleColor: '#e2e8f0',
            bodyColor: '#f1f5f9',
            borderColor: 'rgba(148, 163, 184, 0.35)',
            borderWidth: 1,
            padding: 10,
            displayColors: true,
            callbacks: {
                title(items) {
                    if (!items.length) {
                        return '';
                    }
                    const x = items[0].parsed.x;
                    return typeof x === 'number' ? `Время ${formatAxisSec(x)}` : '';
                },
                label(ctx) {
                    const ds = ctx.dataset.label || '';
                    const y = ctx.parsed.y;
                    const v = typeof y === 'number' && Number.isFinite(y)
                        ? (Number.isInteger(y) ? String(y) : y.toFixed(2))
                        : '—';
                    return ds ? `${ds}: ${v}` : v;
                },
            },
        },
    },
    scales: {
        x: {
            type: 'linear',
            title: {
                display: true,
                text: 'Время от начала матча (с)',
                color: '#9ca3af',
                font: { size: 11 },
            },
            ticks: {
                ...axisStyle.ticks,
                callback: (v) => formatAxisSec(Number(v)),
            },
            grid: axisStyle.grid,
            border: axisStyle.border,
        },
        y: {
            ticks: { color: '#a1a1aa' },
            grid: axisStyle.grid,
            border: axisStyle.border,
            afterFit(scale) {
                if (typeof scale.isHorizontal === 'function' && !scale.isHorizontal()) {
                    const minW = readTimelinePlotGutterLeftPx();
                    // Нельзя ужимать ширину ниже расчёта Chart.js — иначе chartArea ломается и график не рисуется.
                    scale.width = Math.max(scale.width, minW);
                }
            },
        },
    },
};

function fetchBattleInsightJSON(levelIndex) {
    const tr = typeof getTimeRangeJSON === 'function' ? getTimeRangeJSON() : '{}';
    const fn = globalThis.getBattleInsightJSON;
    if (typeof fn !== 'function') {
        return null;
    }
    return fn(levelIndex, tr);
}

function destroyCharts() {
    chartPointerUnsubs.forEach((fn) => {
        try {
            fn();
        } catch (_) {
            /* ignore */
        }
    });
    chartPointerUnsubs = [];
    if (lifeChart) {
        lifeChart.destroy();
        lifeChart = null;
    }
    if (intensityChart) {
        intensityChart.destroy();
        intensityChart = null;
    }
}

export function setupBattleInsightCharts(getMatchIndex) {
    const lifeCanvas = document.getElementById('battle_life_chart');
    const intCanvas = document.getElementById('battle_intensity_chart');
    const lifeWrap = lifeCanvas ? lifeCanvas.closest('.battle-insight-canvas-wrap') : null;
    const intWrap = intCanvas ? intCanvas.closest('.battle-insight-canvas-wrap') : null;
    if (!lifeCanvas || !intCanvas || !lifeWrap || !intWrap) {
        return { refresh: () => {} };
    }

    if (!cursorUnsub) {
        cursorUnsub = subscribeTimelineCursor(() => {
            if (lifeChart) {
                lifeChart.update('none');
            }
            if (intensityChart) {
                intensityChart.update('none');
            }
        });
    }

    function refresh() {
        const idx = typeof getMatchIndex === 'function' ? getMatchIndex() : 0;
        const myGen = ++battleInsightLoadGen;
        showChartPreloader(lifeWrap, { label: 'Линия жизни…' });
        showChartPreloader(intWrap, { label: 'Интенсивность боя…' });
        deferAfterPaint(() => {
            try {
                if (myGen !== battleInsightLoadGen) {
                    return;
                }
                const tbw0 = performance.now();
                const raw = fetchBattleInsightJSON(idx);
                const battleWasmMs = performance.now() - tbw0;
                destroyCharts();
                if (!raw || raw === 'null') {
                    return;
                }
                let data;
                const tbj0 = performance.now();
                try {
                    data = JSON.parse(raw);
                } catch (_) {
                    return;
                }
                const battleJsonMs = performance.now() - tbj0;
                const tSeries0 = performance.now();
                const life = Array.isArray(data.life) ? data.life : [];
                const intensity = Array.isArray(data.intensity) ? data.intensity : [];
                const allyLabel = data.ally_team_label || 'Союзники';
                const enemyLabel = data.enemy_team_label || 'Противники';

                const lifeT = life.map((p) => p.t);
                const lifeAlly = life.map((p) => p.allies);
                const lifeEnemy = life.map((p) => p.enemies);
                const lifeX = xDomainSec(lifeT);
                const intT = intensity.map((p) => p.t);
                const intAlly = intensity.map((p) => p.ally);
                const intEnemy = intensity.map((p) => p.enemy);
                const intX = xDomainSec(intT);
                const battleSeriesPrepMs = performance.now() - tSeries0;

                const tChart0 = performance.now();
                if (lifeT.length > 0) {
                    lifeChart = new Chart(lifeCanvas.getContext('2d'), {
                        type: 'line',
                        data: {
                            datasets: [
                                {
                                    label: `${allyLabel} (живых)`,
                                    data: lifeAlly.map((y, i) => ({ x: lifeT[i], y })),
                                    borderColor: 'rgba(52, 211, 153, 0.95)',
                                    backgroundColor: 'rgba(52, 211, 153, 0.12)',
                                    stepped: 'after',
                                    fill: false,
                                    tension: 0,
                                    borderWidth: 2,
                                    pointRadius: 0,
                                    pointHoverRadius: 4,
                                },
                                {
                                    label: `${enemyLabel} (живых)`,
                                    data: lifeEnemy.map((y, i) => ({ x: lifeT[i], y })),
                                    borderColor: 'rgba(248, 113, 113, 0.95)',
                                    backgroundColor: 'rgba(248, 113, 113, 0.1)',
                                    stepped: 'after',
                                    fill: false,
                                    tension: 0,
                                    borderWidth: 2,
                                    pointRadius: 0,
                                    pointHoverRadius: 4,
                                },
                            ],
                        },
                        options: {
                            ...commonLineOptions,
                            parsing: false,
                            scales: {
                                ...commonLineOptions.scales,
                                x: {
                                    ...commonLineOptions.scales.x,
                                    type: 'linear',
                                    min: lifeX.min,
                                    max: lifeX.max,
                                },
                                y: {
                                    ...commonLineOptions.scales.y,
                                    title: {
                                        display: true,
                                        text: 'Игроков в живых',
                                        color: '#9ca3af',
                                        font: { size: 11 },
                                    },
                                    beginAtZero: true,
                                    ticks: { ...axisStyle.ticks, stepSize: 1 },
                                },
                            },
                        },
                    });
                    chartPointerUnsubs.push(attachChartBrushAndCursor(lifeChart, lifeCanvas, lifeWrap));
                }

                if (intT.length > 0) {
                    intensityChart = new Chart(intCanvas.getContext('2d'), {
                        type: 'line',
                        data: {
                            datasets: [
                                {
                                    label: `${allyLabel}, урон/с (окно 10 с)`,
                                    data: intAlly.map((y, i) => ({ x: intT[i], y })),
                                    borderColor: 'rgba(56, 189, 248, 0.95)',
                                    backgroundColor: 'rgba(56, 189, 248, 0.08)',
                                    fill: false,
                                    tension: 0.15,
                                    borderWidth: 2,
                                    pointRadius: 0,
                                    pointHoverRadius: 3,
                                },
                                {
                                    label: `${enemyLabel}, урон/с (окно 10 с)`,
                                    data: intEnemy.map((y, i) => ({ x: intT[i], y })),
                                    borderColor: 'rgba(251, 146, 60, 0.95)',
                                    backgroundColor: 'rgba(251, 146, 60, 0.08)',
                                    fill: false,
                                    tension: 0.15,
                                    borderWidth: 2,
                                    pointRadius: 0,
                                    pointHoverRadius: 3,
                                },
                            ],
                        },
                        options: {
                            ...commonLineOptions,
                            parsing: false,
                            scales: {
                                ...commonLineOptions.scales,
                                x: {
                                    ...commonLineOptions.scales.x,
                                    type: 'linear',
                                    min: intX.min,
                                    max: intX.max,
                                },
                                y: {
                                    ...commonLineOptions.scales.y,
                                    title: {
                                        display: true,
                                        text: 'Урон в секунду (среднее за 10 с)',
                                        color: '#9ca3af',
                                        font: { size: 11 },
                                    },
                                    beginAtZero: true,
                                },
                            },
                        },
                    });
                    chartPointerUnsubs.push(attachChartBrushAndCursor(intensityChart, intCanvas, intWrap));
                }

                const battleChartRenderMs = performance.now() - tChart0;
                const prepTotal = battleWasmMs + battleJsonMs + battleSeriesPrepMs;
                gaEvent('lux_battle_charts_timing', {
                    battle_wasm_ms: Math.round(battleWasmMs),
                    battle_json_ms: Math.round(battleJsonMs),
                    battle_series_prep_ms: Math.round(battleSeriesPrepMs),
                    battle_data_prep_ms: Math.round(prepTotal),
                    battle_chart_render_ms: Math.round(battleChartRenderMs),
                    life_points: lifeT.length,
                    intensity_points: intT.length,
                });

                const sec = getTimelineCursorSec();
                if (sec !== null && (lifeChart || intensityChart)) {
                    setTimelineCursorSec(sec);
                }
            } finally {
                if (myGen === battleInsightLoadGen) {
                    hideChartPreloader(lifeWrap);
                    hideChartPreloader(intWrap);
                }
            }
        });
    }

    window.addEventListener('resize', () => {
        if (lifeChart) {
            lifeChart.options.devicePixelRatio = chartDpr();
            lifeChart.resize();
        }
        if (intensityChart) {
            intensityChart.options.devicePixelRatio = chartDpr();
            intensityChart.resize();
        }
    });

    return { refresh };
}
