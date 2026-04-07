/**
 * Линия жизни + интенсивность боя — синхронизированы с окном времени getTimeRangeJSON().
 */
import { Chart, registerables } from 'chart.js';
import { getRelativePosition } from 'chart.js/helpers';
import {
    commitZoomFromBrush,
    getTimeRangeBounds,
    getTimeRangeJSON,
    getTimelineStartUnixMs,
    getTimelineCursorSec,
    setTimelineCursorSec,
    subscribeTimelineCursor,
} from './timeline.js';
import { deferAfterPaint, hideChartPreloader, showChartPreloader } from './chart_preloader.js';
import { gaEvent } from './analytics.js';
import { getFocusedPlayer } from './player_focus.js';

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
const swapCookieName = 'lux_team_swap_by_match';
const intFmt = new Intl.NumberFormat('ru-RU');

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

function formatAxisClock(sec) {
    if (typeof sec !== 'number' || Number.isNaN(sec)) {
        return '';
    }
    const startMs = getTimelineStartUnixMs();
    if (!startMs) {
        return formatAxisSec(sec);
    }
    const d = new Date(startMs + sec * 1000);
    return d.toLocaleTimeString('ru-RU', {
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false,
    });
}

function formatWholeNumber(v) {
    const n = typeof v === 'number' ? v : Number(v);
    if (!Number.isFinite(n)) {
        return '—';
    }
    return intFmt.format(Math.round(n));
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
            right: 12,
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
                    return typeof x === 'number' ? `Время ${formatAxisClock(x)}` : '';
                },
                label(ctx) {
                    const ds = ctx.dataset.label || '';
                    const y = ctx.parsed.y;
                    const v = formatWholeNumber(y);
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
                text: 'Текущее время',
                color: '#9ca3af',
                font: { size: 11 },
            },
            ticks: {
                ...axisStyle.ticks,
                callback: (v) => formatAxisClock(Number(v)),
            },
            grid: axisStyle.grid,
            border: axisStyle.border,
        },
        y: {
            ticks: {
                color: '#a1a1aa',
                callback: (v) => formatWholeNumber(v),
            },
            grid: axisStyle.grid,
            border: axisStyle.border,
            afterFit(scale) {
                if (typeof scale.isHorizontal === 'function' && !scale.isHorizontal()) {
                    const minW = readTimelinePlotGutterLeftPx();
                    // Фиксируем ширину оси Y, чтобы таймлайн и оба графика совпадали по X-позициям.
                    scale.width = minW;
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
    const focused = getFocusedPlayer(levelIndex);
    return fn(levelIndex, tr, focused || '');
}

function seriesHead(points, max = 5) {
    if (!Array.isArray(points)) {
        return [];
    }
    return points.slice(0, max).map((p) => ({
        t: p && p.t,
        allies: p && p.allies,
        enemies: p && p.enemies,
        ally: p && p.ally,
        enemy: p && p.enemy,
        player: p && p.player,
        player_out: p && p.player_out,
        player_in: p && p.player_in,
    }));
}

function fetchTimelineMarkers(levelIndex, focusedPlayer) {
    const fn = globalThis.getTimelineJSON;
    if (typeof fn !== 'function') {
        return { markers: [], allyTeamID: 0, enemyTeamID: 0 };
    }
    try {
        const raw = fn(levelIndex, focusedPlayer || '');
        const data = JSON.parse(raw || '{}');
        return {
            markers: Array.isArray(data && data.markers) ? data.markers : [],
            allyTeamID: Number(data && data.ally_team_id) || 0,
            enemyTeamID: Number(data && data.enemy_team_id) || 0,
        };
    } catch (_) {
        return { markers: [], allyTeamID: 0, enemyTeamID: 0 };
    }
}

function selectedTeamIDs(levelIndex, allyTeamID, enemyTeamID) {
    // Для семантики событий используем реальные teamID из backend без инверсии.
    // Swap влияет на раскладку/вид, но не должен переопределять "кто союзник".
    void levelIndex;
    return {
        ally: allyTeamID || 0,
        enemy: enemyTeamID || 0,
    };
}

function markerRoleByTeams(marker, allyTeamID, enemyTeamID) {
    const kind = String((marker && marker.kind) || '');
    const teamID = Number(marker && marker.team_id) || 0;
    const victimTeamID = Number(marker && marker.victim_team_id) || 0;
    if (kind === 'spawn' || kind === 'enemy_spawn') {
        if (teamID && teamID === allyTeamID) return 'spawn';
        if (teamID && teamID === enemyTeamID) return 'enemy_spawn';
        return kind;
    }
    if (kind === 'kill' || kind === 'death') {
        if (victimTeamID && victimTeamID === allyTeamID) return 'death';
        if (victimTeamID && victimTeamID === enemyTeamID) return 'kill';
        return kind;
    }
    return kind;
}

function markerRelatedToPlayer(m, focusedPlayer) {
    const fp = String(focusedPlayer || '').trim().toLowerCase();
    if (!fp || !m || typeof m !== 'object') {
        return false;
    }
    const has = (v) => String(v || '').trim().toLowerCase() === fp;
    if (has(m.player) || has(m.killer) || has(m.victim)) {
        return true;
    }
    if (Array.isArray(m.assists) && m.assists.some((a) => has(a))) {
        return true;
    }
    return false;
}

function relatedEventTimes(markers, focusedPlayer) {
    return (Array.isArray(markers) ? markers : [])
        .filter((m) => markerRelatedToPlayer(m, focusedPlayer) && typeof m.time_sec === 'number')
        .map((m) => m.time_sec)
        .sort((a, b) => a - b);
}

function hasRelatedEventAtSec(sec, times) {
    if (typeof sec !== 'number' || Number.isNaN(sec) || !Array.isArray(times) || !times.length) {
        return false;
    }
    // Небольшой допуск на случай погрешности float при сериализации/отрисовке.
    const tol = 1e-4;
    for (let i = 0; i < times.length; i++) {
        if (Math.abs(times[i] - sec) <= tol) {
            return true;
        }
    }
    return false;
}

function formatMarkerTooltipLine(m, allyTeamID, enemyTeamID, focusedPlayer) {
    if (!m || typeof m !== 'object') {
        return '';
    }
    const related = markerRelatedToPlayer(m, focusedPlayer);
    const mark = related ? '● ' : '· ';
    const kind = markerRoleByTeams(m, allyTeamID, enemyTeamID);
    const killer = String(m.killer || '').trim();
    const victim = String(m.victim || m.player || '').trim();
    const player = String(m.player || '').trim();
    const ship = String(m.player_ship || '').trim();
    const killerShip = String(m.killer_ship || '').trim();
    const victimShip = String(m.victim_ship || '').trim();
    const weapon = String(m.weapon || '').trim();
    const assists = Array.isArray(m.assists) ? m.assists.filter(Boolean).map(String) : [];
    if (kind === 'spawn') {
        return `${mark}Спавн (союзники): ${player}${ship ? ` · ${ship}` : ''}`;
    }
    if (kind === 'enemy_spawn') {
        return `${mark}Спавн (враги): ${player}${ship ? ` · ${ship}` : ''}`;
    }
    if (kind === 'kill') {
        const k = killerShip ? `${killer} (${killerShip})` : killer;
        const v = victimShip ? `${victim} (${victimShip})` : victim;
        const extra = [weapon ? `Оружие: ${weapon}` : '', assists.length ? `Помогали: ${assists.join(', ')}` : ''].filter(Boolean).join(' · ');
        return extra ? `${mark}Убийство: ${k} → ${v} · ${extra}` : `${mark}Убийство: ${k} → ${v}`;
    }
    if (kind === 'death') {
        const k = killerShip ? `${killer} (${killerShip})` : killer;
        const v = victimShip ? `${victim} (${victimShip})` : victim;
        const extra = [weapon ? `Оружие: ${weapon}` : '', assists.length ? `Помогали: ${assists.join(', ')}` : ''].filter(Boolean).join(' · ');
        return extra ? `${mark}Смерть союзника: ${k} → ${v} · ${extra}` : `${mark}Смерть союзника: ${k} → ${v}`;
    }
    return `${mark}${String(m.label || '').trim()}`;
}

function markerLinesForSec(markers, sec, allyTeamID, enemyTeamID, focusedPlayer) {
    const tol = 0.26;
    const near = (Array.isArray(markers) ? markers : [])
        .filter((m) => typeof m.time_sec === 'number' && Math.abs(m.time_sec - sec) <= tol)
        .sort((a, b) => {
            const ar = markerRelatedToPlayer(a, focusedPlayer) ? 0 : 1;
            const br = markerRelatedToPlayer(b, focusedPlayer) ? 0 : 1;
            if (ar !== br) {
                return ar - br;
            }
            return Math.abs(a.time_sec - sec) - Math.abs(b.time_sec - sec);
        })
        .slice(0, 5);
    return near.map((m) => formatMarkerTooltipLine(m, allyTeamID, enemyTeamID, focusedPlayer)).filter(Boolean);
}

function readSwapStateForMatch(levelIndex) {
    if (typeof document === 'undefined') {
        return false;
    }
    const row = document.cookie
        .split('; ')
        .find((x) => x.startsWith(`${swapCookieName}=`));
    if (!row) {
        return false;
    }
    try {
        const raw = decodeURIComponent(row.slice(swapCookieName.length + 1));
        const map = JSON.parse(raw);
        if (!map || typeof map !== 'object') {
            return false;
        }
        return map[String(Number(levelIndex) || 0)] === 1;
    } catch (_) {
        return false;
    }
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
                console.debug('[battle] raw payload', {
                    match_index: idx,
                    raw_len: typeof raw === 'string' ? raw.length : 0,
                    wasm_ms: Math.round(battleWasmMs),
                });
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
                const focusedPlayer = String(data.focused_player || '').trim();
                const timelineData = fetchTimelineMarkers(idx, '');
                const baseAllyTeamID = Number(timelineData.allyTeamID) || Number(data.ally_team_id) || 0;
                const baseEnemyTeamID = Number(timelineData.enemyTeamID) || Number(data.enemy_team_id) || 0;
                const selectedTeams = selectedTeamIDs(idx, baseAllyTeamID, baseEnemyTeamID);
                // Для подсказок на линии жизни нужны все события, а не только фильтр выбранного игрока.
                const timelineMarkers = timelineData.markers;
                const focusedEventTimes = relatedEventTimes(timelineMarkers, focusedPlayer);
                console.debug('[battle] parsed series', {
                    match_index: idx,
                    focused_player: focusedPlayer,
                    ally_label: allyLabel,
                    enemy_label: enemyLabel,
                    ally_team_id: selectedTeams.ally,
                    enemy_team_id: selectedTeams.enemy,
                    life_points: life.length,
                    intensity_points: intensity.length,
                    life_head: seriesHead(life),
                    intensity_head: seriesHead(intensity),
                    timeline_markers: timelineMarkers.length,
                    focused_events: focusedEventTimes.length,
                    json_ms: Math.round(battleJsonMs),
                });

                const lifeT = life.map((p) => p.t);
                const lifeAllies = life.map((p) => p.allies);
                const lifeEnemies = life.map((p) => p.enemies);
                const lifeFocused = life.map((p) => (p && p.player ? 1 : 0));
                const lifeX = xDomainSec(lifeT);
                const intT = intensity.map((p) => p.t);
                const intAllies = intensity.map((p) => p.ally);
                const intEnemies = intensity.map((p) => p.enemy);
                const intFocusedOut = intensity.map((p) => p.player_out || 0);
                const intFocusedIn = intensity.map((p) => p.player_in || 0);
                const intX = xDomainSec(intT);
                const allyIsBaseAlly = !selectedTeams.ally || selectedTeams.ally === baseAllyTeamID;
                // Жестко маппим по выбранным teamID: союзники всегда "ally" (зеленый), противники всегда "enemy" (красный).
                const lifeAllySeries = allyIsBaseAlly ? lifeAllies : lifeEnemies;
                const lifeEnemySeries = allyIsBaseAlly ? lifeEnemies : lifeAllies;
                const intAllySeries = allyIsBaseAlly ? intAllies : intEnemies;
                const intEnemySeries = allyIsBaseAlly ? intEnemies : intAllies;
                const battleSeriesPrepMs = performance.now() - tSeries0;

                const tChart0 = performance.now();
                if (lifeT.length > 0) {
                    const lifeDatasets = [
                        {
                            label: `${allyLabel} (живых)`,
                            data: lifeAllySeries.map((y, i) => ({ x: lifeT[i], y })),
                            borderColor: 'rgba(52, 211, 153, 0.95)',
                            backgroundColor: 'rgba(52, 211, 153, 0.12)',
                            stepped: 'before',
                            fill: false,
                            tension: 0,
                            borderWidth: 2,
                            pointRadius: 0,
                            pointHoverRadius: 4,
                        },
                        {
                            label: `${enemyLabel} (живых)`,
                            data: lifeEnemySeries.map((y, i) => ({ x: lifeT[i], y })),
                            borderColor: 'rgba(248, 113, 113, 0.95)',
                            backgroundColor: 'rgba(248, 113, 113, 0.1)',
                            stepped: 'before',
                            fill: false,
                            tension: 0,
                            borderWidth: 2,
                            pointRadius: 0,
                            pointHoverRadius: 4,
                        },
                    ];
                    if (focusedPlayer) {
                        lifeDatasets.push({
                            label: `${focusedPlayer} (жив/мертв)`,
                            data: lifeFocused.map((y, i) => ({ x: lifeT[i], y })),
                            borderColor: 'rgba(250, 204, 21, 0.95)',
                            backgroundColor: 'rgba(250, 204, 21, 0.08)',
                            stepped: 'before',
                            fill: false,
                            tension: 0,
                            borderWidth: 2,
                            // В ключевых точках, связанных с выбранным игроком, рисуем заметные маркеры.
                            pointRadius(ctx) {
                                const raw = ctx && ctx.raw;
                                const x = raw && typeof raw.x === 'number' ? raw.x : NaN;
                                return hasRelatedEventAtSec(x, focusedEventTimes) ? 5 : 0;
                            },
                            pointHoverRadius(ctx) {
                                const raw = ctx && ctx.raw;
                                const x = raw && typeof raw.x === 'number' ? raw.x : NaN;
                                return hasRelatedEventAtSec(x, focusedEventTimes) ? 7 : 4;
                            },
                            pointBackgroundColor: 'rgba(250, 204, 21, 1)',
                            pointBorderColor: 'rgba(255, 255, 255, 0.92)',
                            pointBorderWidth: 1.5,
                            borderDash: [7, 5],
                        });
                    }
                    lifeChart = new Chart(lifeCanvas.getContext('2d'), {
                        type: 'line',
                        data: {
                            datasets: lifeDatasets,
                        },
                        options: {
                            ...commonLineOptions,
                            parsing: false,
                            plugins: {
                                ...commonLineOptions.plugins,
                                tooltip: {
                                    ...commonLineOptions.plugins.tooltip,
                                    callbacks: {
                                        ...commonLineOptions.plugins.tooltip.callbacks,
                                        afterBody(items) {
                                            if (!items.length) {
                                                return [];
                                            }
                                            const x = items[0].parsed && items[0].parsed.x;
                                            if (typeof x !== 'number' || Number.isNaN(x)) {
                                                return [];
                                            }
                                            const lines = markerLinesForSec(
                                                timelineMarkers,
                                                x,
                                                selectedTeams.ally,
                                                selectedTeams.enemy,
                                                focusedPlayer,
                                            );
                                            return lines.length ? ['События:'].concat(lines) : [];
                                        },
                                    },
                                },
                            },
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
                    const intDatasets = [
                        {
                            label: `${allyLabel}, урон/с (окно 10 с)`,
                            data: intAllySeries.map((y, i) => ({ x: intT[i], y })),
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
                            data: intEnemySeries.map((y, i) => ({ x: intT[i], y })),
                            borderColor: 'rgba(251, 146, 60, 0.95)',
                            backgroundColor: 'rgba(251, 146, 60, 0.08)',
                            fill: false,
                            tension: 0.15,
                            borderWidth: 2,
                            pointRadius: 0,
                            pointHoverRadius: 3,
                        },
                    ];
                    if (focusedPlayer) {
                        intDatasets.push({
                            label: `${focusedPlayer}, исходящий урон/с`,
                            data: intFocusedOut.map((y, i) => ({ x: intT[i], y })),
                            borderColor: 'rgba(250, 204, 21, 0.95)',
                            backgroundColor: 'rgba(250, 204, 21, 0.08)',
                            fill: false,
                            tension: 0.15,
                            borderWidth: 2,
                            pointRadius: 0,
                            pointHoverRadius: 3,
                            borderDash: [7, 5],
                        });
                        intDatasets.push({
                            label: `${focusedPlayer}, входящий урон/с`,
                            data: intFocusedIn.map((y, i) => ({ x: intT[i], y })),
                            borderColor: 'rgba(244, 114, 182, 0.95)',
                            backgroundColor: 'rgba(244, 114, 182, 0.08)',
                            fill: false,
                            tension: 0.15,
                            borderWidth: 2,
                            pointRadius: 0,
                            pointHoverRadius: 3,
                            borderDash: [4, 4],
                        });
                    }
                    intensityChart = new Chart(intCanvas.getContext('2d'), {
                        type: 'line',
                        data: {
                            datasets: intDatasets,
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
                console.debug('[battle] rendered', {
                    match_index: idx,
                    focused_player: focusedPlayer,
                    life_points: lifeT.length,
                    intensity_points: intT.length,
                    data_prep_ms: Math.round(prepTotal),
                    render_ms: Math.round(battleChartRenderMs),
                });
                gaEvent('lux_battle_charts_timing', {
                    swapped: allyIsBaseAlly ? '0' : '1',
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

    window.addEventListener('lux-team-swap-changed', () => {
        refresh();
    });

    return { refresh };
}
