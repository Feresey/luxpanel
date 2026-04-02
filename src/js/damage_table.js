/**
 * Damage table panel:
 * - defaults mode: multiple rows (modifier presets)
 * - single mode: one row (user-picked modifier tri-state)
 * - columns: source, targets (multi if recipient empty), modifiers (no nil), summary
 */

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
        const s = r.summary || {};
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${escapeHtml(r.source || '')}</td>
            <td>${renderTargetsCell(r.targets)}</td>
            <td>${buildModifiersCell(r.modifiers || {})}</td>
            <td>${escapeHtml(String(typeof s.hits === 'number' ? s.hits : 0))}</td>
            <td>${escapeHtml(formatNum(s.damage))}</td>
            <td>${escapeHtml(formatNum(s.hull))}</td>
            <td>${escapeHtml(formatNum(s.shield))}</td>
            <td>${escapeHtml(formatNum(s.avg_damage))}</td>
        `;
        tbody.appendChild(tr);
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
    if (!parsed || !parsed.row) {
        return [];
    }
    return [parsed.row];
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

    const makeColumn = (list) => {
        const col = document.createElement('div');
        col.className = 'df-mod-column';
        for (let i = 0; i < list.length; i++) {
            col.appendChild(buildTriStateSelect(list[i]));
        }
        return col;
    };

    columns.appendChild(makeColumn(left));
    columns.appendChild(makeColumn(right));

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

function getLevelIndex(getMatchIndex) {
    return typeof getMatchIndex === 'function' ? getMatchIndex() : 0;
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
    const raw = metaFn(levelIndex, initiator);
    if (!raw || raw === 'null') {
        populateCustomModifiers([]);
        return;
    }
    const meta = JSON.parse(raw);
    const modifiers = Array.isArray(meta.modifiers) ? meta.modifiers : [];
    populateCustomModifiers(modifiers);
}

function refreshTable() {
    if (!panelEls) {
        return;
    }
    const levelIndex = getLevelIndex(panelEls.getMatchIndex);
    const hint = panelEls.hint;
    const tbody = panelEls.tbody;

    const initiator = panelEls.initiator.value.trim();
    const recipient = panelEls.recipient.value.trim();
    const weapon = panelEls.weapon.value.trim();
    const damageType = panelEls.damageType.value || 'total';
    const mode = panelEls.mode.value || 'defaults';

    if (!initiator) {
        if (hint) hint.textContent = 'Выберите игрока (источник урона)';
        renderTableBody(tbody, []);
        return;
    }

    if (hint) hint.textContent = '';

    const baseFilters = {
        initiator,
        recipient: recipient || '',
        weapon: weapon || '',
        damage_type: damageType,
    };

    try {
        if (mode === 'single') {
            const modifiers = collectCustomModifiers();
            const rows = applySingleRowTable(levelIndex, { ...baseFilters, modifiers });
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
        const n = mode === 'single' ? 1 : (panelEls.tbody.rows.length || 0);
        panelEls.foot.textContent = n ? `${modeText}: ${n}` : '';
    }
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

    let raw = metaFn(levelIndex, '');
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
    } else if (players.length) {
        panelEls.initiator.value = players[0];
    }

    const initiator = panelEls.initiator.value.trim();
    if (!initiator) {
        return;
    }

    raw = metaFn(levelIndex, initiator);
    if (!raw || raw === 'null') {
        fillSelect(panelEls.recipient, [], 'Любая цель');
        fillSelect(panelEls.weapon, [], 'Любое оружие');
        populateCustomModifiers([]);
        return;
    }

    meta = JSON.parse(raw);
    const recipients = Array.isArray(meta.recipients) ? meta.recipients : [];
    const weapons = Array.isArray(meta.weapons) ? meta.weapons : [];

    fillSelect(panelEls.recipient, recipients, 'Любая цель');
    fillSelect(panelEls.weapon, weapons, 'Любое оружие');
    refreshModifiersMeta(levelIndex);
    refreshTable();
}

export function setupDamageTablePanel(getMatchIndex) {
    const section = document.querySelector('.damage-table-section');
    if (!section) {
        return;
    }

    const initiator = document.getElementById('df_initiator');
    const recipient = document.getElementById('df_recipient');
    const weapon = document.getElementById('df_weapon');
    const damageType = document.getElementById('df_damage_type');
    const tbody = document.getElementById('df_tbody');
    const hint = document.getElementById('df_hint');
    const foot = document.getElementById('df_foot');
    const applyBtn = document.getElementById('df_apply');
    const clearMods = document.getElementById('df_clear_modifiers');
    const mode = document.getElementById('df_table_mode');
    const customBlock = document.getElementById('df_custom_modifiers_block');

    if (!initiator || !recipient || !weapon || !damageType || !tbody || !mode || !customBlock) {
        return;
    }

    panelEls = {
        getMatchIndex,
        initiator,
        recipient,
        weapon,
        damageType,
        tbody,
        hint,
        foot,
        mode,
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
    damageType.addEventListener('change', () => refreshTable());
    mode.addEventListener('change', () => setModeVisibility());

    if (applyBtn) {
        applyBtn.addEventListener('click', () => refreshTable());
    }

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

    // Initial load.
    const levelIndex = getLevelIndex(getMatchIndex);
    refreshMeta(levelIndex);

    return { refresh: () => refreshMeta(getLevelIndex(getMatchIndex)), apply: () => refreshTable() };
}

