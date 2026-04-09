import { Chart, registerables } from "chart.js";
import ChartDataLabels from "chartjs-plugin-datalabels";
import { getTimeRangeJSON } from "./timeline.js";
import { deferAfterPaint, hideChartPreloader, showChartPreloader } from "./chart_preloader.js";
import { gaEvent } from "./analytics.js";

Chart.register(...registerables, ChartDataLabels);

export { CreateCharts, ApplyParsedCharts, setupGraphViewToolbar, clearTeamSwapCookieState, applyTeamSwapState };

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
let graphDamageType = 'all';
let graphIncludeBots = false;
let teamsSwapped = false;
let currentLevelIndex = 0;
const swapCookieName = 'lux_team_swap_by_match';
let swapByMatch = readSwapCookieMap();
let swapBtnEl = null;
const intFmt = new Intl.NumberFormat('ru-RU');

function readSwapCookieMap() {
    if (typeof document === 'undefined') {
        return {};
    }
    const row = document.cookie
        .split('; ')
        .find((x) => x.startsWith(`${swapCookieName}=`));
    if (!row) {
        return {};
    }
    try {
        const raw = decodeURIComponent(row.slice(swapCookieName.length + 1));
        const v = JSON.parse(raw);
        if (!v || typeof v !== 'object') {
            return {};
        }
        return v;
    } catch (_) {
        return {};
    }
}

function writeSwapCookieMap() {
    if (typeof document === 'undefined') {
        return;
    }
    const raw = encodeURIComponent(JSON.stringify(swapByMatch));
    // 30 дней достаточно, при новой пачке логов очищаем вручную.
    document.cookie = `${swapCookieName}=${raw}; path=/; max-age=${60 * 60 * 24 * 30}; samesite=lax`;
}

function swapStateForMatch(levelIndex) {
    const k = String(Number(levelIndex) || 0);
    return swapByMatch[k] === 1;
}

function setSwapStateForMatch(levelIndex, swapped) {
    const k = String(Number(levelIndex) || 0);
    if (swapped) {
        swapByMatch[k] = 1;
    } else {
        delete swapByMatch[k];
    }
    writeSwapCookieMap();
}

function syncSwapStateForLevel(levelIndex) {
    currentLevelIndex = Number(levelIndex) || 0;
    teamsSwapped = swapStateForMatch(currentLevelIndex);
}

function syncSwapButtonUi() {
    if (!swapBtnEl) {
        return;
    }
    swapBtnEl.classList.toggle('is-swapped', teamsSwapped);
    swapBtnEl.setAttribute('aria-pressed', teamsSwapped ? 'true' : 'false');
    swapBtnEl.title = teamsSwapped
        ? 'Сейчас слева Team 2. Нажмите, чтобы вернуть Team 1 слева.'
        : 'Сейчас слева Team 1. Нажмите, чтобы поменять местами.';
}

function formatWholeNumber(v) {
    const n = typeof v === 'number' ? v : Number(v);
    if (!Number.isFinite(n)) {
        return '—';
    }
    return intFmt.format(Math.round(n));
}

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
                    const s = formatWholeNumber(v);
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
    const opts = JSON.stringify({
        damage_type: graphDamageType,
        include_bots: graphIncludeBots,
    });
    return fn(levelIndex, String(mode), tr, opts);
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
    if (v === 0) {
        return '0';
    }
    return formatWholeNumber(v);
}

function matrixCellAt(mat, c, r) {
    const row = Array.isArray(mat[c]) ? mat[c] : [];
    return row[r];
}

function renderPlayerLabelWithNewbieBonus(name, newbieBonusMax) {
    const bonus = Number(newbieBonusMax && newbieBonusMax[name]) || 0;
    const safeName = escapeHtml(name);
    if (bonus <= 0) {
        return safeName;
    }
    const hint = `Бонус новичка (Normalizer): максимум за матч = x${bonus}`;
    return `${safeName} <span class="newbie-bonus-badge" title="${escapeHtml(hint)}">NB x${bonus}</span>`;
}

function isSimpleUiModeActive() {
    return typeof document !== 'undefined'
        && !!document.body
        && document.body.classList.contains('simple-ui-mode');
}

function renderSimplePanelTable(panel, mode, newbieBonusMax) {
    const colPlayers = Array.isArray(panel.col_players) ? panel.col_players : [];
    const mat = Array.isArray(panel.matrix) ? panel.matrix : [];
    const sums = colPlayers.map((name, colIdx) => {
        let total = 0;
        for (let r = 0; r < mat.length; r++) {
            const row = Array.isArray(mat[r]) ? mat[r] : [];
            const v = Number(row[colIdx]) || 0;
            total += v;
        }
        return { name, total };
    });
    sums.sort((a, b) => b.total - a.total || String(a.name).localeCompare(String(b.name)));
    const label = mode === 'heal' ? 'Лечение' : mode === 'kill' ? 'Киллы' : 'Урон';
    let html = '<table class="simple-team-table"><thead><tr>';
    html += '<th>Игрок</th>';
    html += `<th>${escapeHtml(label)}</th>`;
    html += '</tr></thead><tbody>';
    sums.forEach((x) => {
        const v = mode === 'kill' ? String(Math.round(x.total || 0)) : formatWholeNumber(x.total);
        html += '<tr>';
        html += `<th class="row-head">${renderPlayerLabelWithNewbieBonus(x.name, newbieBonusMax)}</th>`;
        html += `<td>${escapeHtml(v)}</td>`;
        html += '</tr>';
    });
    html += '</tbody></table>';
    html += '<p class="simple-team-foot">Простой режим: вклад игроков без матрицы связей.</p>';
    return html;
}

function currentFocusedPlayerName() {
    const el = document.getElementById('graph_player_focus_select');
    return el ? String(el.value || '').trim() : '';
}

function moveNameFirst(names, focus) {
    if (!focus || !Array.isArray(names) || names.length < 2) {
        return names.slice();
    }
    const canon = String(focus).trim().toLowerCase();
    const out = names.slice();
    const i = out.findIndex((n) => String(n || '').trim().toLowerCase() === canon);
    if (i > 0) {
        const [v] = out.splice(i, 1);
        out.unshift(v);
    }
    return out;
}

/** Фон ячейки: насыщенность по доле от максимума всей таблицы (глобальная нормализация). */
function matrixCellHeatStyle(v, maxInTable, mode) {
    if (typeof v !== 'number' || Number.isNaN(v) || v <= 0 || maxInTable <= 0) {
        return '';
    }
    const t = Math.min(1, v / maxInTable);
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

function renderMatrixTable(panel, mode, opts = {}) {
    const colPlayers = Array.isArray(panel.col_players) ? panel.col_players : [];
    const rowPlayers = Array.isArray(panel.row_players) ? panel.row_players : [];
    const newbieBonusMax =
        panel && panel.newbie_bonus_max && typeof panel.newbie_bonus_max === 'object'
            ? panel.newbie_bonus_max
            : {};
    if (isSimpleUiModeActive()) {
        return renderSimplePanelTable(panel, mode, newbieBonusMax);
    }
    const mat = Array.isArray(panel.matrix) ? panel.matrix : [];
    // Транспонированная таблица: заголовки столбцов = бывшие строки, строки = бывшие столбцы.
    let headerCols = rowPlayers.slice();
    let bodyRows = colPlayers.slice();
    const focus = String(opts.focusName || '').trim();
    const focusPlace = opts.focusPlace === 'col' ? 'col' : 'row';
    if (focus) {
        if (focusPlace === 'row') {
            bodyRows = moveNameFirst(bodyRows, focus);
        } else {
            headerCols = moveNameFirst(headerCols, focus);
        }
    }
    const corner = mode === 'heal' ? 'Леч. \\ Получ.' : 'Источн. \\ Цель';
    const rowIdxByName = new Map(colPlayers.map((n, i) => [String(n), i]));
    const colIdxByName = new Map(rowPlayers.map((n, i) => [String(n), i]));

    const rowTotals = bodyRows.map((_, r) => {
        let s = 0;
        const srcIdx = rowIdxByName.get(String(bodyRows[r]));
        for (let c = 0; c < headerCols.length; c++) {
            const dstIdx = colIdxByName.get(String(headerCols[c]));
            const v = matrixCellAt(mat, dstIdx, srcIdx);
            if (typeof v === 'number' && !Number.isNaN(v)) {
                s += v;
            }
        }
        return s;
    });
    const colTotals = headerCols.map((_, c) => {
        let s = 0;
        const dstIdx = colIdxByName.get(String(headerCols[c]));
        for (let r = 0; r < bodyRows.length; r++) {
            const srcIdx = rowIdxByName.get(String(bodyRows[r]));
            const v = matrixCellAt(mat, dstIdx, srcIdx);
            if (typeof v === 'number' && !Number.isNaN(v)) {
                s += v;
            }
        }
        return s;
    });
    const grandTotal = rowTotals.reduce((a, b) => a + b, 0);

    let maxInTable = 0;
    for (let r = 0; r < bodyRows.length; r++) {
        for (let c = 0; c < headerCols.length; c++) {
            const srcIdx = rowIdxByName.get(String(bodyRows[r]));
            const dstIdx = colIdxByName.get(String(headerCols[c]));
            const v = matrixCellAt(mat, dstIdx, srcIdx);
            if (typeof v === 'number' && !Number.isNaN(v) && v > maxInTable) {
                maxInTable = v;
            }
        }
    }

    let html = '<table class="pair-matrix-table"><thead><tr>';
    html += `<th class="corner">${escapeHtml(corner)}</th>`;
    for (let c = 0; c < headerCols.length; c++) {
        html += `<th title="${escapeHtml(headerCols[c])}">${renderPlayerLabelWithNewbieBonus(headerCols[c], newbieBonusMax)}</th>`;
    }
    html += '<th class="matrix-total-head" title="Сумма по строке (источник)">Total</th>';
    html += '</tr></thead><tbody>';
    for (let r = 0; r < bodyRows.length; r++) {
        html += '<tr>';
        html += `<th class="row-head" title="${escapeHtml(bodyRows[r])}">${renderPlayerLabelWithNewbieBonus(bodyRows[r], newbieBonusMax)}</th>`;
        for (let c = 0; c < headerCols.length; c++) {
            const srcIdx = rowIdxByName.get(String(bodyRows[r]));
            const dstIdx = colIdxByName.get(String(headerCols[c]));
            const v = matrixCellAt(mat, dstIdx, srcIdx);
            const heat = matrixCellHeatStyle(v, maxInTable, mode);
            const peak = typeof v === 'number' && maxInTable > 0 && Math.abs(v - maxInTable) < 1e-6;
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

function pairPanelsForView(panels) {
    const p0 = panels[0] || {};
    const p1 = panels[1] || {};
    return [p0, p1];
}

function updateTeamPanelTitles() {
    const titles = document.querySelectorAll('[data-team-panel] .graph-box-title');
    if (titles.length < 2) {
        return;
    }
    titles[0].textContent = 'Team 1';
    titles[1].textContent = 'Team 2';
}

function clearTeamSwapCookieState() {
    teamsSwapped = false;
    currentLevelIndex = 0;
    updateTeamPanelTitles();
}

function applyTeamSwapState(levelIndex, swapped) {
    void levelIndex;
    void swapped;
    teamsSwapped = false;
    updateTeamPanelTitles();
    return true;
}

/** @returns {{ matrices_wasm_ms: number, matrices_json_ms: number, matrix_dom_ms: number }} */
function applyMatrixTables(levelIndex, mode) {
    const timing = { matrices_wasm_ms: 0, matrices_json_ms: 0, matrix_dom_ms: 0 };
    const tw0 = performance.now();
    const raw = fetchMatricesJSON(levelIndex, mode);
    timing.matrices_wasm_ms = Math.round(performance.now() - tw0);
    const c0 = document.getElementById('chart_matrix_0');
    const c1 = document.getElementById('chart_matrix_1');
    if (!c0 || !c1) {
        return timing;
    }
    if (!raw || raw === 'null') {
        const td0 = performance.now();
        c0.innerHTML = '<p class="graph-matrix-empty">Нет данных</p>';
        c1.innerHTML = '<p class="graph-matrix-empty">Нет данных</p>';
        timing.matrix_dom_ms = Math.round(performance.now() - td0);
        return timing;
    }
    let data;
    const tj0 = performance.now();
    try {
        data = JSON.parse(raw);
    } catch (_) {
        timing.matrices_json_ms = Math.round(performance.now() - tj0);
        c0.innerHTML = '';
        c1.innerHTML = '';
        return timing;
    }
    timing.matrices_json_ms = Math.round(performance.now() - tj0);
    const panels = Array.isArray(data.panels) ? data.panels : [];
    const metric = data.metric || mode;
    const [p0, p1] = pairPanelsForView(panels);
    const focusName = currentFocusedPlayerName();
    const td0 = performance.now();
    c0.innerHTML = renderMatrixTable(p0, metric, { focusName, focusPlace: 'row' });
    c1.innerHTML = renderMatrixTable(p1, metric, { focusName, focusPlace: 'col' });
    const t0 = document.getElementById('graph_total_0');
    const t1 = document.getElementById('graph_total_1');
    if (t0) {
        t0.textContent = formatTotal(p0.total, metric);
    }
    if (t1) {
        t1.textContent = formatTotal(p1.total, metric);
    }
    updateMatrixHints(metric);
    timing.matrix_dom_ms = Math.round(performance.now() - td0);
    return timing;
}

function formatTotal(total, mode) {
    if (typeof total !== 'number' || Number.isNaN(total)) {
        return '—';
    }
    return `Σ ${formatWholeNumber(total)}`;
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

function emitChartsTiming(mode, viewMode, chartsWasmMs, chartsJsonMs, pieRenderMs, matrixTiming) {
    const m = matrixTiming || { matrices_wasm_ms: 0, matrices_json_ms: 0, matrix_dom_ms: 0 };
    const dataPrep =
        chartsWasmMs + chartsJsonMs + m.matrices_wasm_ms + m.matrices_json_ms;
    const renderMs = pieRenderMs + m.matrix_dom_ms;
    gaEvent('lux_charts_timing', {
        metric: mode,
        view: viewMode,
        get_charts_wasm_ms: Math.round(chartsWasmMs),
        charts_json_ms: Math.round(chartsJsonMs),
        get_matrices_wasm_ms: m.matrices_wasm_ms,
        matrices_json_ms: m.matrices_json_ms,
        charts_data_prep_ms: Math.round(dataPrep),
        pie_render_ms: Math.round(pieRenderMs),
        matrix_dom_ms: m.matrix_dom_ms,
        charts_render_ms: Math.round(renderMs),
    });
}

function ApplyParsedCharts(levelIndex = 0, mode = 'damage') {
    syncSwapStateForLevel(levelIndex);
    syncSwapButtonUi();
    updateTeamPanelTitles();
    const chartsWasmMs = 0;
    const chartsJsonMs = 0;
    const pieRenderMs = 0;
    const c0 = document.getElementById('chart_matrix_0');
    const c1 = document.getElementById('chart_matrix_1');
    const myGen = ++matrixLoadGen;
    showChartPreloader(c0, { label: 'Таблица Team 1…' });
    showChartPreloader(c1, { label: 'Таблица Team 2…' });
    deferAfterPaint(() => {
        try {
            if (myGen !== matrixLoadGen) {
                return;
            }
            const mt = applyMatrixTables(levelIndex, mode);
            if (myGen === matrixLoadGen) {
                emitChartsTiming(mode, 'table', chartsWasmMs, chartsJsonMs, pieRenderMs, mt);
            }
        } finally {
            if (myGen === matrixLoadGen) {
                hideChartPreloader(c0);
                hideChartPreloader(c1);
            }
        }
    });
}

function setupGraphViewToolbar(onViewChange) {
    const root = document.getElementById('graphs_root');
    if (!root) {
        return;
    }
    graphToolbarRefresh = typeof onViewChange === 'function' ? onViewChange : null;
    setGraphsRootView('table');
    const dmgTypeSel = document.getElementById('graph_damage_type_select');
    if (dmgTypeSel) {
        dmgTypeSel.addEventListener('change', () => {
            graphDamageType = String(dmgTypeSel.value || 'all');
            if (graphToolbarRefresh) {
                graphToolbarRefresh();
            }
        });
    }
    const botsCb = document.getElementById('graph_include_bots_cb');
    if (botsCb) {
        botsCb.addEventListener('change', () => {
            graphIncludeBots = !!botsCb.checked;
            if (graphToolbarRefresh) {
                graphToolbarRefresh();
            }
        });
    }

    swapBtnEl = document.getElementById('graph_swap_teams_btn');
    if (swapBtnEl) {
        swapBtnEl.hidden = true;
    }
    updateTeamPanelTitles();
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
    pieChart1 = null;
    pieChart2 = null;
}
