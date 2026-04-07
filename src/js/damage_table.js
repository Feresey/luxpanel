/**
 * Damage table panel:
 * - defaults mode: multiple rows (modifier presets)
 * - single mode: one row (user-picked modifier tri-state)
 * - columns: source, targets (multi if recipient empty), modifiers (no nil), summary
 */

import { getTimeRangeJSON } from './timeline.js';
import { deferAfterPaint, hideChartPreloader, showChartPreloader } from './chart_preloader.js';

const emptyOptionValue = '';

const MODIFIER_ORDER_LEFT = ['EMP', 'THERMAL', 'KINETIC'];
const MODIFIER_ORDER_RIGHT = ['PRIMARY_WEAPON', 'EXPLOSION', 'CRIT'];
const MODIFIER_ORDER_KNOWN = [...MODIFIER_ORDER_LEFT, ...MODIFIER_ORDER_RIGHT];

function escapeHtml(s) {
    return String(s)
        .replace(/&/g, '&amp;')
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;')
        .replace(/\"/g, '&quot;');
}

function formatNum(n) {
    if (typeof n !== 'number' || Number.isNaN(n)) {
        return '—';
    }
    if (Number.isInteger(n)) {
        return String(n);
    }
    return n.toFixed(2);
}

function formatTime(sec) {
    if (typeof sec !== 'number' || Number.isNaN(sec)) {
        return '—';
    }
    const m = Math.floor(sec / 60);
    const s = sec - m * 60;
    const mm = String(m).padStart(2, '0');
    const ss = s < 10 ? `0${s.toFixed(2)}` : s.toFixed(2);
    return `${mm}:${ss}`;
}

function getMetaFn() {
    return globalThis.getDamageFilterMetaJSON;
}

function getDefaultTableFn() {
    return globalThis.getDamageDefaultTableJSON;
}

function getSingleRowFn() {
    return globalThis.getDamageTableJSON;
}

function fillSelect(select, values, emptyLabel = '— выберите —') {
    select.innerHTML = '';
    const oEmpty = document.createElement('option');
    oEmpty.value = emptyOptionValue;
    oEmpty.textContent = emptyLabel;
    select.appendChild(oEmpty);

    if (Array.isArray(values)) {
        for (let i = 0; i < values.length; i++) {
            const v = values[i];
            const o = document.createElement('option');
            o.value = v;
            o.textContent = v;
            select.appendChild(o);
        }
    }
    select.value = emptyOptionValue;
}

function buildModifiersCell(modifiers) {
    // Modifiers map contains only non-nil keys (value exists => nil omitted in JSON).
    // We still must respect requested display order.
    if (!modifiers || typeof modifiers !== 'object') {
        return '—';
    }

    const lines = [];
    const addIfPresent = (k) => {
        if (Object.prototype.hasOwnProperty.call(modifiers, k)) {
            lines.push(`${k}: ${modifiers[k] ? 'true' : 'false'}`);
        }
    };

    for (let i = 0; i < MODIFIER_ORDER_LEFT.length; i++) {
        addIfPresent(MODIFIER_ORDER_LEFT[i]);
    }
    for (let i = 0; i < MODIFIER_ORDER_RIGHT.length; i++) {
        addIfPresent(MODIFIER_ORDER_RIGHT[i]);
    }

    // Add remaining present modifiers (not nil, meaning present in the map)
    const knownSet = new Set(MODIFIER_ORDER_KNOWN);
    const others = Object.keys(modifiers).filter((k) => !knownSet.has(k));
    others.sort();
    for (let i = 0; i < others.length; i++) {
        const k = others[i];
        if (Object.prototype.hasOwnProperty.call(modifiers, k)) {
            lines.push(`${k}: ${modifiers[k] ? 'true' : 'false'}`);
        }
    }

    if (!lines.length) {
        return '—';
    }
    return lines.map((l) => `<div>${escapeHtml(l)}</div>`).join('');
}

function renderTargetsCell(targets) {
    if (!Array.isArray(targets) || targets.length === 0) {
        return '—';
    }
    return targets.map((t) => `<div>${escapeHtml(t)}</div>`).join('');
}

function appendDetailBreakdown(tbody, events) {
    const trDetail = document.createElement('tr');
    trDetail.className = 'df-row-detail df-row-detail-hidden';
    const td = document.createElement('td');
    td.colSpan = 8;
    const innerWrap = document.createElement('div');
    innerWrap.className = 'df-detail-inner';
    const table = document.createElement('table');
    table.className = 'df-inner-table';
    const thead = document.createElement('thead');
    thead.innerHTML = '<tr>'
        + '<th>Время</th><th>Источник</th><th>Цель</th><th>Оружие</th><th>Модификаторы</th><th>Урон</th>'
        + '</tr>';
    table.appendChild(thead);
    const itbody = document.createElement('tbody');
    for (let j = 0; j < events.length; j++) {
        const e = events[j] || {};
        const itr = document.createElement('tr');
        const modsHtml = buildModifiersCell(e.modifiers || {});
        itr.innerHTML = `
            <td>${escapeHtml(formatTime(e.time_sec))}</td>
            <td>${escapeHtml(e.initiator || '')}</td>
            <td>${escapeHtml(e.recipient || '')}</td>
            <td>${escapeHtml(e.weapon || '')}</td>
            <td class="df-cell-mods">${modsHtml}</td>
            <td>${escapeHtml(formatNum(e.amount))}</td>
        `;
        itbody.appendChild(itr);
    }
    table.appendChild(itbody);
    innerWrap.appendChild(table);
    td.appendChild(innerWrap);
    trDetail.appendChild(td);
    tbody.appendChild(trDetail);
}

function renderTableBody(tbody, rows) {
    tbody.innerHTML = '';
    if (!Array.isArray(rows) || rows.length === 0) {
        const tr = document.createElement('tr');
        const td = document.createElement('td');
        td.colSpan = 8;
        td.className = 'df-empty';
        td.textContent = 'Нет данных';
        tr.appendChild(td);
        tbody.appendChild(tr);
        return;
    }

    for (let i = 0; i < rows.length; i++) {
        const r = rows[i] || {};
        const isEvent = !!r.__event;
        const s = r.summary || {};
        const hits = isEvent ? 1 : (typeof s.hits === 'number' ? s.hits : 0);
        const dmg = isEvent ? r.amount : s.damage;
        const incoming = panelEls && panelEls.perspective && panelEls.perspective.value === 'recipient';
        let source = isEvent ? r.initiator : r.source;
        let targets = isEvent ? [r.recipient] : r.targets;
        if (incoming && !isEvent) {
            // Для входящего урона хотим видеть список источников именно в колонке "Источник урона".
            const srcList = Array.isArray(r.targets) ? r.targets : [];
            source = srcList.length ? srcList.join(', ') : '—';
            targets = r && r.source ? [r.source] : [];
        }
        const weapon = isEvent ? (r.weapon || '') : '—';
        const time = isEvent ? formatTime(r.time_sec) : '—';

        const evs = r.events;
        const expandable = !isEvent && Array.isArray(evs) && evs.length > 0;

        let expandCell = '<td class="df-cell-expand"></td>';
        if (!isEvent && expandable) {
            expandCell = '<td class="df-cell-expand">'
                + '<button type="button" class="df-expand-btn" aria-expanded="false" '
                + 'aria-label="Показать исходные строки урона">▶</button>'
                + '</td>';
        }

        const tr = document.createElement('tr');
        tr.className = 'df-row-main';
        tr.innerHTML = expandCell
            + `<td>${escapeHtml(time)}</td>`
            + `<td>${escapeHtml(source || '')}</td>`
            + `<td>${renderTargetsCell(targets)}</td>`
            + `<td>${escapeHtml(weapon)}</td>`
            + `<td>${buildModifiersCell(r.modifiers || {})}</td>`
            + `<td>${escapeHtml(String(hits))}</td>`
            + `<td>${escapeHtml(formatNum(dmg))}</td>`;
        tbody.appendChild(tr);

        if (expandable) {
            appendDetailBreakdown(tbody, evs);
        }
    }
}

function applyDefaultsTable(levelIndex, filters) {
    const tableFn = getDefaultTableFn();
    if (typeof tableFn !== 'function') {
        return null;
    }
    const raw = tableFn(levelIndex, JSON.stringify(filters));
    if (!raw || raw === 'null') {
        return null;
    }
    const parsed = JSON.parse(raw);
    if (parsed && parsed.error) {
        throw new Error(parsed.error);
    }
    return parsed.rows || [];
}

function applySingleRowTable(levelIndex, filters) {
    const fn = getSingleRowFn();
    if (typeof fn !== 'function') {
        return null;
    }
    const raw = fn(levelIndex, JSON.stringify(filters));
    if (!raw || raw === 'null') {
        return null;
    }
    const parsed = JSON.parse(raw);
    if (parsed && parsed.error) {
        throw new Error(parsed.error);
    }
    if (!parsed) {
        return [];
    }
    const aggregate = filters.aggregate !== false;
    if (aggregate && parsed.row) {
        return [{
            ...parsed.row,
            events: Array.isArray(parsed.events) ? parsed.events : [],
        }];
    }
    if (Array.isArray(parsed.events) && parsed.events.length) {
        return parsed.events.map((e) => ({ ...e, __event: true }));
    }
    if (parsed.row) {
        return [parsed.row];
    }
    return [];
}

function buildTriStateSelect(mod) {
    const label = document.createElement('label');
    label.className = 'df-mod-item';

    const span = document.createElement('span');
    span.textContent = mod;

    const select = document.createElement('select');
    select.className = 'df-mod-select';
    select.dataset.mod = mod;

    const oNil = document.createElement('option');
    oNil.value = '';
    oNil.textContent = 'nil';

    const oTrue = document.createElement('option');
    oTrue.value = 'true';
    oTrue.textContent = 'true';

    const oFalse = document.createElement('option');
    oFalse.value = 'false';
    oFalse.textContent = 'false';

    select.appendChild(oNil);
    select.appendChild(oTrue);
    select.appendChild(oFalse);
    select.value = '';

    label.appendChild(span);
    label.appendChild(select);
    return label;
}

function populateCustomModifiers(modifiersAvailable) {
    const wrap = document.getElementById('df_custom_modifiers_wrap');
    if (!wrap) {
        return;
    }
    wrap.innerHTML = '';

    const available = new Set(Array.isArray(modifiersAvailable) ? modifiersAvailable : []);

    const left = MODIFIER_ORDER_LEFT.filter((m) => available.has(m));
    const right = MODIFIER_ORDER_RIGHT.filter((m) => available.has(m));
    const knownSet = new Set([...MODIFIER_ORDER_LEFT, ...MODIFIER_ORDER_RIGHT]);
    const others = Array.from(available).filter((m) => !knownSet.has(m)).sort();

    const columns = document.createElement('div');
    columns.className = 'df-mod-columns';

    const makeColumn = (list, compact) => {
        const col = document.createElement('div');
        col.className = 'df-mod-column';
        if (compact) {
            col.classList.add('compact');
        }
        for (let i = 0; i < list.length; i++) {
            col.appendChild(buildTriStateSelect(list[i]));
        }
        return col;
    };

    columns.appendChild(makeColumn(left, true));
    columns.appendChild(makeColumn(right, false));

    wrap.appendChild(columns);

    if (others.length) {
        const otherWrap = document.createElement('div');
        otherWrap.className = 'df-modifiers-other';
        for (let i = 0; i < others.length; i++) {
            otherWrap.appendChild(buildTriStateSelect(others[i]));
        }
        wrap.appendChild(otherWrap);
    }
}

function collectCustomModifiers() {
    const wrap = document.getElementById('df_custom_modifiers_wrap');
    if (!wrap) {
        return {};
    }
    const selects = wrap.querySelectorAll('select.df-mod-select');
    const out = {};
    selects.forEach((sel) => {
        const mod = sel.dataset.mod;
        const v = sel.value;
        if (!mod || v === '') {
            return; // nil => omit from map
        }
        out[mod] = v === 'true';
    });
    return out;
}

let panelEls = null;
let damageTableLoadGen = 0;

function runDamageTableLoad(work) {
    if (!panelEls || !panelEls.tbody) {
        return;
    }
    const scroll = panelEls.tbody.closest('.df-table-scroll');
    if (!scroll) {
        work();
        return;
    }
    const myGen = ++damageTableLoadGen;
    showChartPreloader(scroll, { label: 'Таблица урона…' });
    deferAfterPaint(() => {
        try {
            if (myGen !== damageTableLoadGen) {
                return;
            }
            work();
        } finally {
            if (myGen === damageTableLoadGen) {
                hideChartPreloader(scroll);
            }
        }
    });
}

function getLevelIndex(getMatchIndex) {
    return typeof getMatchIndex === 'function' ? getMatchIndex() : 0;
}

function mergeTimeRange(base) {
    let tr = {};
    try {
        if (typeof getTimeRangeJSON === 'function') {
            tr = JSON.parse(getTimeRangeJSON());
        }
    } catch (_) {
        tr = {};
    }
    return { ...base, ...tr };
}

function refreshModifiersMeta(levelIndex) {
    if (!panelEls) {
        return;
    }
    const metaFn = getMetaFn();
    if (typeof metaFn !== 'function') {
        return;
    }
    const initiator = panelEls.initiator.value.trim();
    const tr = typeof getTimeRangeJSON === 'function' ? getTimeRangeJSON() : '{}';
    const raw = metaFn(levelIndex, initiator, tr);
    if (!raw || raw === 'null') {
        populateCustomModifiers([]);
        return;
    }
    const meta = JSON.parse(raw);
    const modifiers = Array.isArray(meta.modifiers) ? meta.modifiers : [];
    populateCustomModifiers(modifiers);
}

function refreshTableSync() {
    if (!panelEls) {
        return;
    }
    const levelIndex = getLevelIndex(panelEls.getMatchIndex);
    const hint = panelEls.hint;
    const tbody = panelEls.tbody;

    const initiator = panelEls.initiator.value.trim();
    const recipient = panelEls.recipient.value.trim();
    const weapon = panelEls.weapon.value.trim();
    const mode = panelEls.mode.value || 'defaults';
    const perspective = panelEls.perspective && panelEls.perspective.value === 'recipient'
        ? 'recipient'
        : 'initiator';

    if (!initiator) {
        if (hint) {
            hint.textContent = perspective === 'recipient'
                ? 'Выберите игрока (по которому получен урон)'
                : 'Выберите игрока (источник урона)';
        }
        renderTableBody(tbody, []);
        return;
    }

    if (hint) hint.textContent = '';

    const baseFilters = mergeTimeRange({
        initiator,
        recipient: recipient || '',
        weapon: weapon || '',
        perspective,
    });

    try {
        if (mode === 'single') {
            const modifiers = collectCustomModifiers();
            const rows = applySingleRowTable(levelIndex, { ...baseFilters, modifiers, aggregate: true });
            renderTableBody(tbody, rows);
        } else {
            const rows = applyDefaultsTable(levelIndex, baseFilters);
            renderTableBody(tbody, rows);
        }
    } catch (e) {
        if (hint) hint.textContent = String(e);
        renderTableBody(tbody, []);
        return;
    }

    if (panelEls.foot) {
        const modeText = mode === 'single' ? 'Фильтр' : 'Наборы';
        const n = panelEls.tbody.querySelectorAll('tr.df-row-main').length || 0;
        panelEls.foot.textContent = n ? `${modeText}: ${n}` : '';
    }
}

function refreshTable() {
    runDamageTableLoad(() => refreshTableSync());
}

function refreshMetaSync(levelIndex) {
    if (!panelEls) {
        return;
    }
    const metaFn = getMetaFn();
    const hint = panelEls.hint;
    if (typeof metaFn !== 'function') {
        if (hint) hint.textContent = 'WASM не загружен';
        return;
    }

    const tr0 = typeof getTimeRangeJSON === 'function' ? getTimeRangeJSON() : '{}';
    let raw = metaFn(levelIndex, '', tr0);
    if (!raw || raw === 'null') {
        fillSelect(panelEls.initiator, [], '— нет данных —');
        fillSelect(panelEls.recipient, [], 'Любая цель');
        fillSelect(panelEls.weapon, [], 'Любое оружие');
        return;
    }
    let meta = {};
    try {
        meta = JSON.parse(raw);
    } catch (_) {
        if (hint) hint.textContent = 'Ошибка метаданных';
        return;
    }

    const players = Array.isArray(meta.players) ? meta.players : [];
    const prevInitiator = panelEls.initiator.value.trim();
    fillSelect(panelEls.initiator, players, '— выберите игрока —');

    if (prevInitiator && players.includes(prevInitiator)) {
        panelEls.initiator.value = prevInitiator;
    } else {
        panelEls.initiator.value = emptyOptionValue;
    }

    const initiator = panelEls.initiator.value.trim();
    if (!initiator) {
        fillSelect(panelEls.recipient, [], panelEls.perspective && panelEls.perspective.value === 'recipient' ? 'Любой источник' : 'Любая цель');
        fillSelect(panelEls.weapon, [], 'Любое оружие');
        populateCustomModifiers([]);
        refreshTableSync();
        return;
    }

    const perspective = panelEls.perspective && panelEls.perspective.value === 'recipient'
        ? 'recipient'
        : 'initiator';
    if (perspective === 'recipient') {
        const sources = players.filter((p) => p !== initiator);
        fillSelect(panelEls.recipient, sources, 'Любой источник');
        raw = metaFn(levelIndex, '', tr0);
        if (!raw || raw === 'null') {
            fillSelect(panelEls.weapon, [], 'Любое оружие');
            populateCustomModifiers([]);
            return;
        }
        meta = JSON.parse(raw);
    } else {
        raw = metaFn(levelIndex, initiator, tr0);
        if (!raw || raw === 'null') {
            fillSelect(panelEls.recipient, [], 'Любая цель');
            fillSelect(panelEls.weapon, [], 'Любое оружие');
            populateCustomModifiers([]);
            return;
        }
        meta = JSON.parse(raw);
        const recipients = Array.isArray(meta.recipients) ? meta.recipients : [];
        fillSelect(panelEls.recipient, recipients, 'Любая цель');
    }
    const weapons = Array.isArray(meta.weapons) ? meta.weapons : [];
    fillSelect(panelEls.weapon, weapons, 'Любое оружие');
    refreshModifiersMeta(levelIndex);
    refreshTableSync();
}

function refreshMeta(levelIndex) {
    if (!panelEls) {
        return;
    }
    const metaFn = getMetaFn();
    const hint = panelEls.hint;
    if (typeof metaFn !== 'function') {
        if (hint) hint.textContent = 'WASM не загружен';
        return;
    }
    runDamageTableLoad(() => refreshMetaSync(levelIndex));
}

export function setupDamageTablePanel(getMatchIndex) {
    const section = document.querySelector('.damage-table-section');
    if (!section) {
        return;
    }

    const initiator = document.getElementById('df_initiator');
    const recipient = document.getElementById('df_recipient');
    const weapon = document.getElementById('df_weapon');
    const tbody = document.getElementById('df_tbody');
    const hint = document.getElementById('df_hint');
    const foot = document.getElementById('df_foot');
    const clearMods = document.getElementById('df_clear_modifiers');
    const mode = document.getElementById('df_table_mode');
    const perspective = document.getElementById('df_perspective');
    const customBlock = document.getElementById('df_custom_modifiers_block');

    if (!initiator || !recipient || !weapon || !tbody || !mode || !perspective || !customBlock) {
        return;
    }

    panelEls = {
        getMatchIndex,
        initiator,
        recipient,
        weapon,
        tbody,
        hint,
        foot,
        mode,
        perspective,
    };

    function setModeVisibility() {
        customBlock.style.display = panelEls.mode.value === 'single' ? 'block' : 'none';
        if (panelEls.mode.value === 'single') {
            const levelIndex = getLevelIndex(panelEls.getMatchIndex);
            refreshModifiersMeta(levelIndex);
        }
        refreshTable();
    }

    initiator.addEventListener('change', () => {
        const levelIndex = getLevelIndex(panelEls.getMatchIndex);
        refreshMeta(levelIndex);
    });
    recipient.addEventListener('change', () => refreshTable());
    weapon.addEventListener('change', () => refreshTable());
    mode.addEventListener('change', () => setModeVisibility());
    perspective.addEventListener('change', () => {
        const levelIndex = getLevelIndex(panelEls.getMatchIndex);
        refreshMeta(levelIndex);
    });

    if (clearMods) {
        clearMods.addEventListener('click', () => {
            const wrap = document.getElementById('df_custom_modifiers_wrap');
            if (wrap) {
                const selects = wrap.querySelectorAll('select.df-mod-select');
                selects.forEach((s) => { s.value = ''; });
            }
            refreshTable();
        });
    }

    // If custom modifiers change, update table.
    const wrap = document.getElementById('df_custom_modifiers_wrap');
    if (wrap) {
        wrap.addEventListener('change', (e) => {
            if (e.target && e.target.matches('select.df-mod-select')) {
                refreshTable();
            }
        });
    }

    tbody.addEventListener('click', (e) => {
        const btn = e.target.closest('.df-expand-btn');
        if (!btn || !tbody.contains(btn)) {
            return;
        }
        e.preventDefault();
        const tr = btn.closest('tr');
        if (!tr || !tr.classList.contains('df-row-main')) {
            return;
        }
        const detail = tr.nextElementSibling;
        if (!detail || !detail.classList.contains('df-row-detail')) {
            return;
        }
        detail.classList.toggle('df-row-detail-hidden');
        const isHidden = detail.classList.contains('df-row-detail-hidden');
        btn.setAttribute('aria-expanded', String(!isHidden));
        btn.textContent = isHidden ? '▶' : '▼';
    });

    // Initial load.
    const levelIndex = getLevelIndex(getMatchIndex);
    refreshMeta(levelIndex);

    return { refresh: () => refreshMeta(getLevelIndex(getMatchIndex)), apply: () => refreshTable() };
}

