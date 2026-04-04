import { Chart, registerables } from "chart.js";
import ChartDataLabels from "chartjs-plugin-datalabels";
import { getTimeRangeJSON } from "./timeline.js";
import { deferAfterPaint, hideChartPreloader, showChartPreloader } from "./chart_preloader.js";

Chart.register(...registerables, ChartDataLabels);

export { CreateCharts, ApplyParsedCharts, setupGraphViewToolbar };

const pieBorder = 'rgba(24, 24, 28, 0.55)';
const pieBorderWidth = 1.5;

const pieColors = [
    'rgba(56, 189, 248, 0.92)',
    'rgba(167, 139, 250, 0.92)',
    'rgba(52, 211, 153, 0.92)',
    'rgba(251, 191, 36, 0.92)',
    'rgba(251, 113, 133, 0.92)',
    'rgba(45, 212, 191, 0.92)',
    'rgba(251, 146, 60, 0.92)',
    'rgba(232, 121, 249, 0.92)',
    'rgba(148, 163, 184, 0.92)',
    'rgba(96, 165, 250, 0.92)',
];

let graphViewMode = 'table';
let graphToolbarRefresh = null;
let matrixLoadGen = 0;

function chartDevicePixelRatio() {
    if (typeof window === "undefined" || !window.devicePixelRatio) {
        return 1;
    }
    return Math.min(Math.max(window.devicePixelRatio, 1), 2.5);
}

const pieOptions = {
    responsive: true,
    maintainAspectRatio: true,
    aspectRatio: 1,
    devicePixelRatio: chartDevicePixelRatio(),
    interaction: {
        mode: 'nearest',
        intersect: true,
    },
    plugins: {
        legend: {
            display: false,
        },
        datalabels: {
            display: (ctx) => {
                const v = ctx.dataset.data[ctx.dataIndex];
                const n = typeof v === "number" ? v : Number(v);
                return Number.isFinite(n) && n > 0;
            },
            formatter: (_value, ctx) => {
                const label = ctx.chart.data.labels[ctx.dataIndex];
                return label != null ? String(label) : "";
            },
            color: "#f8fafc",
            font: {
                size: 12,
                weight: "600",
                family: 'system-ui, "Segoe UI", Roboto, sans-serif',
            },
            textStrokeColor: "rgba(15, 23, 42, 0.4)",
            textStrokeWidth: 1.25,
        },
        tooltip: {
            enabled: true,
            backgroundColor: 'rgba(15, 23, 42, 0.92)',
            titleColor: '#e2e8f0',
            bodyColor: '#f1f5f9',
            borderColor: 'rgba(148, 163, 184, 0.35)',
            borderWidth: 1,
            padding: 10,
            displayColors: true,
            callbacks: {
                title() {
                    return '';
                },
                label(ctx) {
                    const name = ctx.label || '';
                    const v = ctx.raw;
                    const n = typeof v === 'number' ? v : Number(v);
                    const s = Number.isInteger(n) ? String(n) : n.toFixed(1);
                    return `${name}: ${s}`;
                },
            },
        },
    },
};

function createPieWithData(ctx, data) {
    const pie_config = {
        type: 'pie',
        data: data,
        options: pieOptions,
    };
    return new Chart(ctx, pie_config);
}

let pieChart1 = null;
let pieChart2 = null;

function getLastValue(curve) {
    if (!curve || !Array.isArray(curve.data) || curve.data.length === 0) {
        return 0;
    }
    return Number(curve.data[curve.data.length - 1]) || 0;
}

function updateChartDataset(chart, labels, values, datasetLabel) {
    if (!chart || !chart.data || !chart.data.datasets || chart.data.datasets.length === 0) {
        return;
    }
    const colors = labels.map((_, i) => pieColors[i % pieColors.length]);
    const borders = labels.map(() => pieBorder);
    chart.data.labels = labels;
    chart.data.datasets[0].label = datasetLabel;
    chart.data.datasets[0].data = values;
    chart.data.datasets[0].backgroundColor = colors;
    chart.data.datasets[0].borderColor = borders;
    chart.data.datasets[0].borderWidth = pieBorderWidth;
    chart.data.datasets[0].hoverBorderWidth = 2;
    chart.update('none');
}

function fetchChartsJSON(levelIndex, mode) {
    const tr = typeof getTimeRangeJSON === 'function' ? getTimeRangeJSON() : '{}';
    const g = globalThis.getChartsJSON;
    if (typeof g === 'function') {
        const raw = g(levelIndex, String(mode), tr);
        if (raw && raw !== 'null') {
            return raw;
        }
    }
    const d = globalThis.getDamageChartsJSON;
    if (mode === 'damage' && typeof d === 'function') {
        return d(levelIndex, tr);
    }
    return null;
}

function fetchMatricesJSON(levelIndex, mode) {
    const tr = typeof getTimeRangeJSON === 'function' ? getTimeRangeJSON() : '{}';
    const fn = globalThis.getChartMatricesJSON;
    if (typeof fn !== 'function') {
        return null;
    }
    return fn(levelIndex, String(mode), tr);
}

function escapeHtml(s) {
    return String(s)
        .replace(/&/g, '&amp;')
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;')
        .replace(/"/g, '&quot;');
}

function formatMatrixCell(v, mode) {
    if (typeof v !== 'number' || Number.isNaN(v)) {
        return '—';
    }
    if (mode === 'kill') {
        return String(Math.round(v));
    }
    if (v === 0) {
        return '0';
    }
    return v.toFixed(1);
}

function matrixCellAt(mat, c, r) {
    const row = Array.isArray(mat[c]) ? mat[c] : [];
    return row[r];
}

/** Фон ячейки: насыщенность по доле от максимума в строке (фокус источника по целям). */
function matrixCellHeatStyle(v, maxInRow, mode) {
    if (typeof v !== 'number' || Number.isNaN(v) || v <= 0 || maxInRow <= 0) {
        return '';
    }
    const t = Math.min(1, v / maxInRow);
    const a = 0.07 + t * 0.42;
    if (mode === 'heal') {
        return `background-color:rgba(52,211,153,${a})`;
    }
    if (mode === 'kill') {
        return `background-color:rgba(251,191,36,${a})`;
    }
    return `background-color:rgba(56,189,248,${a})`;
}

function matrixCellTooltip(sourceName, targetName, v, rowTotal, grandTotal, mode) {
    const parts = [];
    const label = mode === 'heal' ? 'лечение' : mode === 'kill' ? 'киллы' : 'урон';
    parts.push(`${sourceName} → ${targetName}: ${formatMatrixCell(v, mode)} (${label})`);
    if (typeof rowTotal === 'number' && rowTotal > 1e-9 && typeof v === 'number' && Number.isFinite(v)) {
        parts.push(`${((v / rowTotal) * 100).toFixed(1)}% от строки (вклад этого источника по целям)`);
    }
    if (typeof grandTotal === 'number' && grandTotal > 1e-9 && typeof v === 'number' && Number.isFinite(v)) {
        parts.push(`${((v / grandTotal) * 100).toFixed(1)}% от суммы таблицы команды`);
    }
    return parts.join('\n');
}

function argmaxIndex(arr) {
    if (!arr.length) {
        return -1;
    }
    let best = 0;
    for (let i = 1; i < arr.length; i++) {
        if (arr[i] > arr[best]) {
            best = i;
        }
    }
    return best;
}

function renderMatrixInsight(mode, headerCols, bodyRows, rowTotals, colTotals, grandTotal) {
    if (typeof grandTotal !== 'number' || grandTotal <= 1e-9 || !bodyRows.length || !headerCols.length) {
        return '';
    }
    const ri = argmaxIndex(rowTotals);
    const ci = argmaxIndex(colTotals);
    if (ri < 0 || ci < 0) {
        return '';
    }
    const src = bodyRows[ri];
    const tgt = headerCols[ci];
    const rowPct = (rowTotals[ri] / grandTotal) * 100;
    const colPct = (colTotals[ci] / grandTotal) * 100;
    let line = '';
    if (mode === 'heal') {
        line = `Главный вклад по объёму лечения: ${escapeHtml(src)} (${rowPct.toFixed(1)}% от суммы таблицы). Больше всего получено: ${escapeHtml(tgt)} (${colPct.toFixed(1)}% от суммы по получателям).`;
    } else if (mode === 'kill') {
        line = `Больше всего киллов: ${escapeHtml(src)} (${rowPct.toFixed(1)}% от суммы таблицы). Чаще всего убивали: ${escapeHtml(tgt)} (${colPct.toFixed(1)}% от суммы по целям).`;
    } else {
        line = `Главный вклад по урону: ${escapeHtml(src)} (${rowPct.toFixed(1)}% от суммы команды). Больше всего урона получил: ${escapeHtml(tgt)} (${colPct.toFixed(1)}% от суммы по целям).`;
    }
    return `<p class="graph-matrix-insight" role="status">${line}</p>`;
}

function renderMatrixTable(panel, mode) {
    const colPlayers = Array.isArray(panel.col_players) ? panel.col_players : [];
    const rowPlayers = Array.isArray(panel.row_players) ? panel.row_players : [];
    const mat = Array.isArray(panel.matrix) ? panel.matrix : [];
    // Транспонированная таблица: заголовки столбцов = бывшие строки, строки = бывшие столбцы.
    const headerCols = rowPlayers;
    const bodyRows = colPlayers;
    const corner = mode === 'heal' ? 'Леч. \\ Получ.' : 'Источн. \\ Цель';

    const rowTotals = bodyRows.map((_, r) => {
        let s = 0;
        for (let c = 0; c < headerCols.length; c++) {
            const v = matrixCellAt(mat, c, r);
            if (typeof v === 'number' && !Number.isNaN(v)) {
                s += v;
            }
        }
        return s;
    });
    const colTotals = headerCols.map((_, c) => {
        let s = 0;
        for (let r = 0; r < bodyRows.length; r++) {
            const v = matrixCellAt(mat, c, r);
            if (typeof v === 'number' && !Number.isNaN(v)) {
                s += v;
            }
        }
        return s;
    });
    const grandTotal = rowTotals.reduce((a, b) => a + b, 0);

    const maxPerRow = bodyRows.map((_, r) => {
        let m = 0;
        for (let c = 0; c < headerCols.length; c++) {
            const v = matrixCellAt(mat, c, r);
            if (typeof v === 'number' && !Number.isNaN(v) && v > m) {
                m = v;
            }
        }
        return m;
    });

    let html = '<table class="pair-matrix-table"><thead><tr>';
    html += `<th class="corner">${escapeHtml(corner)}</th>`;
    for (let c = 0; c < headerCols.length; c++) {
        html += `<th title="${escapeHtml(headerCols[c])}">${escapeHtml(headerCols[c])}</th>`;
    }
    html += '<th class="matrix-total-head" title="Сумма по строке (источник)">Total</th>';
    html += '</tr></thead><tbody>';
    for (let r = 0; r < bodyRows.length; r++) {
        html += '<tr>';
        html += `<th class="row-head" title="${escapeHtml(bodyRows[r])}">${escapeHtml(bodyRows[r])}</th>`;
        const maxR = maxPerRow[r];
        for (let c = 0; c < headerCols.length; c++) {
            const v = matrixCellAt(mat, c, r);
            const heat = matrixCellHeatStyle(v, maxR, mode);
            const peak = typeof v === 'number' && maxR > 0 && Math.abs(v - maxR) < 1e-6;
            const tip = matrixCellTooltip(bodyRows[r], headerCols[c], v, rowTotals[r], grandTotal, mode);
            const cls = `matrix-cell-inner${peak ? ' matrix-cell-peak' : ''}`;
            const styleAttr = heat ? ` style="${heat}"` : '';
            html += `<td class="${cls}" title="${escapeHtml(tip)}"${styleAttr}>${escapeHtml(formatMatrixCell(v, mode))}</td>`;
        }
        html += `<td class="matrix-total" title="Сумма по строке">${escapeHtml(formatMatrixCell(rowTotals[r], mode))}</td>`;
        html += '</tr>';
    }
    html += '<tr class="matrix-total-row">';
    html += '<th class="row-head matrix-total-label" scope="row">Total</th>';
    for (let c = 0; c < headerCols.length; c++) {
        html += `<td class="matrix-total" title="Сумма по столбцу (цель)">${escapeHtml(formatMatrixCell(colTotals[c], mode))}</td>`;
    }
    html += `<td class="matrix-total matrix-grand" title="Общая сумма таблицы">${escapeHtml(formatMatrixCell(grandTotal, mode))}</td>`;
    html += '</tr>';
    html += '</tbody></table>';
    html += renderMatrixInsight(mode, headerCols, bodyRows, rowTotals, colTotals, grandTotal);
    return html;
}

function setGraphsRootView(mode) {
    const root = document.getElementById('graphs_root');
    if (!root) {
        return;
    }
    root.classList.toggle('graphs-view-pie', mode === 'pie');
    root.classList.toggle('graphs-view-table', mode === 'table');
}

function updateMatrixHints(metric) {
    const h0 = document.getElementById('graph_matrix_hint_0');
    const h1 = document.getElementById('graph_matrix_hint_1');
    let html = '';
    if (metric === 'heal') {
        html = 'Данные за выбранный на таймлайне интервал времени.<br/>'
            + 'Строки — получатель лечения, столбцы — кто лечит (внутри команды).<br/>'
            + 'Цвет ячейки: насыщенность по доле от максимума в этой строке (на кого/чей пул лечения приходится больше всего). Наведите на ячейку — доли от строки и от суммы таблицы.';
    } else if (metric === 'kill') {
        html = 'Данные за выбранный на таймлайне интервал времени.<br/>'
            + 'Строки — игроки этой команды (источник килла), столбцы — вражеская команда (цель).<br/>'
            + 'Цвет ячейки: насыщенность по доле от максимума в строке (главные «пары» киллов). Наведите на ячейку — доли от строки и от суммы таблицы.';
    } else {
        html = 'Данные за выбранный на таймлайне интервал времени.<br/>'
            + 'Строки — игроки этой команды (источник урона), столбцы — вражеская команда (цель).<br/>'
            + 'Цвет ячейки: насыщенность по доле от максимума в строке (на кого этот источник тратил урон сильнее всего). Наведите на ячейку — доли от строки и от суммы таблицы.';
    }
    if (h0) {
        h0.innerHTML = html;
        h0.hidden = false;
    }
    if (h1) {
        h1.innerHTML = html;
        h1.hidden = false;
    }
}

function hideMatrixHints() {
    ['graph_matrix_hint_0', 'graph_matrix_hint_1'].forEach((id) => {
        const el = document.getElementById(id);
        if (el) {
            el.hidden = true;
        }
    });
}

function applyMatrixTables(levelIndex, mode) {
    const raw = fetchMatricesJSON(levelIndex, mode);
    const c0 = document.getElementById('chart_matrix_0');
    const c1 = document.getElementById('chart_matrix_1');
    if (!c0 || !c1) {
        return;
    }
    if (!raw || raw === 'null') {
        c0.innerHTML = '<p class="graph-matrix-empty">Нет данных</p>';
        c1.innerHTML = '<p class="graph-matrix-empty">Нет данных</p>';
        return;
    }
    let data;
    try {
        data = JSON.parse(raw);
    } catch (_) {
        c0.innerHTML = '';
        c1.innerHTML = '';
        return;
    }
    const panels = Array.isArray(data.panels) ? data.panels : [];
    const metric = data.metric || mode;
    const p0 = panels[0] || {};
    const p1 = panels[1] || {};
    c0.innerHTML = renderMatrixTable(p0, metric);
    c1.innerHTML = renderMatrixTable(p1, metric);
    const t0 = document.getElementById('graph_total_0');
    const t1 = document.getElementById('graph_total_1');
    if (t0) {
        t0.textContent = formatTotal(p0.total, metric);
    }
    if (t1) {
        t1.textContent = formatTotal(p1.total, metric);
    }
    updateMatrixHints(metric);
}

function formatTotal(total, mode) {
    if (typeof total !== 'number' || Number.isNaN(total)) {
        return '—';
    }
    const s = mode === 'kill' ? String(Math.round(total)) : total.toFixed(1);
    return `Σ ${s}`;
}

function setPieTotalsFixed(values1, values2, mode) {
    const sum = (arr) => arr.reduce((a, b) => a + (Number(b) || 0), 0);
    const t0 = document.getElementById('graph_total_0');
    const t1 = document.getElementById('graph_total_1');
    const m = mode === 'kill' ? 'kill' : mode === 'heal' ? 'heal' : 'damage';
    if (t0) {
        t0.textContent = formatTotal(sum(values1), m);
    }
    if (t1) {
        t1.textContent = formatTotal(sum(values2), m);
    }
}

function ApplyParsedCharts(levelIndex = 0, mode = 'damage') {
    const raw = fetchChartsJSON(levelIndex, mode);
    if (!raw || raw === 'null') {
        return;
    }

    let parsed;
    try {
        parsed = JSON.parse(raw);
    } catch (_) {
        return;
    }
    if (!Array.isArray(parsed) || parsed.length < 2) {
        return;
    }

    const team1 = Array.isArray(parsed[0]) ? parsed[0] : [];
    const team2 = Array.isArray(parsed[1]) ? parsed[1] : [];
    const t1PlayerCurves = team1.slice(1);
    const t2PlayerCurves = team2.slice(1);

    const labels1 = t1PlayerCurves.map((c) => c.name || "Unknown");
    const values1 = t1PlayerCurves.map((c) => getLastValue(c));
    const labels2 = t2PlayerCurves.map((c) => c.name || "Unknown");
    const values2 = t2PlayerCurves.map((c) => getLastValue(c));

    const label = mode === 'heal' ? 'Heal' : mode === 'kill' ? 'Kills' : 'Damage';
    updateChartDataset(pieChart1, labels1, values1, label);
    updateChartDataset(pieChart2, labels2, values2, label);

    const c0 = document.getElementById('chart_matrix_0');
    const c1 = document.getElementById('chart_matrix_1');

    if (graphViewMode === 'table') {
        const myGen = ++matrixLoadGen;
        showChartPreloader(c0, { label: 'Таблица Team 1…' });
        showChartPreloader(c1, { label: 'Таблица Team 2…' });
        deferAfterPaint(() => {
            try {
                if (myGen !== matrixLoadGen) {
                    return;
                }
                applyMatrixTables(levelIndex, mode);
            } finally {
                if (myGen === matrixLoadGen) {
                    hideChartPreloader(c0);
                    hideChartPreloader(c1);
                }
            }
        });
    } else {
        matrixLoadGen += 1;
        hideChartPreloader(c0);
        hideChartPreloader(c1);
        setPieTotalsFixed(values1, values2, mode);
        hideMatrixHints();
    }
}

function setupGraphViewToolbar(onViewChange) {
    const root = document.getElementById('graphs_root');
    if (!root) {
        return;
    }
    graphToolbarRefresh = typeof onViewChange === 'function' ? onViewChange : null;
    setGraphsRootView(graphViewMode);
    root.querySelectorAll('.graph-view-btn').forEach((btn) => {
        btn.addEventListener('click', () => {
            const v = btn.dataset.graphView;
            if (v !== 'pie' && v !== 'table') {
                return;
            }
            graphViewMode = v;
            root.querySelectorAll('.graph-view-btn').forEach((b) => {
                b.classList.toggle('active', b === btn);
            });
            setGraphsRootView(v);
            if (graphToolbarRefresh) {
                graphToolbarRefresh();
            }
        });
    });
}

function makePiePlaceholder() {
    const bg = pieColors.slice(0, 3);
    return {
        labels: ['—', '—', '—'],
        datasets: [
            {
                label: 'Damage',
                backgroundColor: bg,
                borderColor: bg.map(() => pieBorder),
                borderWidth: pieBorderWidth,
                hoverOffset: 6,
                data: [1, 1, 1],
            },
        ],
    };
}

function CreateCharts() {
    pieChart1 = createPieWithData(
        document.getElementById('pieChart1').getContext('2d'),
        makePiePlaceholder()
    );
    pieChart2 = createPieWithData(
        document.getElementById('pieChart2').getContext('2d'),
        makePiePlaceholder()
    );
}
