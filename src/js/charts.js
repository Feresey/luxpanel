import { Chart, registerables } from "chart.js";
import ChartDataLabels from "chartjs-plugin-datalabels";
import { getTimeRangeJSON } from "./timeline.js";

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

let graphViewMode = 'pie';
let graphToolbarRefresh = null;

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

    let html = '<table class="pair-matrix-table"><thead><tr>';
    html += `<th class="corner">${escapeHtml(corner)}</th>`;
    for (let c = 0; c < headerCols.length; c++) {
        html += `<th title="${escapeHtml(headerCols[c])}">${escapeHtml(headerCols[c])}</th>`;
    }
    html += '<th class="matrix-total-head" title="Total">Total</th>';
    html += '</tr></thead><tbody>';
    for (let r = 0; r < bodyRows.length; r++) {
        html += '<tr>';
        html += `<th class="row-head" title="${escapeHtml(bodyRows[r])}">${escapeHtml(bodyRows[r])}</th>`;
        for (let c = 0; c < headerCols.length; c++) {
            const v = matrixCellAt(mat, c, r);
            html += `<td>${escapeHtml(formatMatrixCell(v, mode))}</td>`;
        }
        html += `<td class="matrix-total">${escapeHtml(formatMatrixCell(rowTotals[r], mode))}</td>`;
        html += '</tr>';
    }
    html += '<tr class="matrix-total-row">';
    html += '<th class="row-head matrix-total-label" scope="row">Total</th>';
    for (let c = 0; c < headerCols.length; c++) {
        html += `<td class="matrix-total">${escapeHtml(formatMatrixCell(colTotals[c], mode))}</td>`;
    }
    html += `<td class="matrix-total matrix-grand">${escapeHtml(formatMatrixCell(grandTotal, mode))}</td>`;
    html += '</tr>';
    html += '</tbody></table>';
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
    const heal = metric === 'heal';
    const text = heal
        ? 'Внутри команды: строки — к получателю, столбцы — кто лечит.'
        : 'Строки — игроки этой команды (источник), столбцы — вражеская команда (цель).';
    if (h0) {
        h0.textContent = text;
        h0.hidden = false;
    }
    if (h1) {
        h1.textContent = text;
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

    if (graphViewMode === 'table') {
        applyMatrixTables(levelIndex, mode);
    } else {
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
