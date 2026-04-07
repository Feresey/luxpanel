// Импортируйте наш пользовательский CSS
import '../scss/styles.scss'

// Импортируйте весь JS Bootstrap
import * as bootstrap from 'bootstrap'

import { CreateCharts, ApplyParsedCharts, applyTeamSwapState, clearTeamSwapCookieState, setupGraphViewToolbar } from './charts.js'
import { setupDamageTablePanel } from './damage_table.js'
import { setupTimeline } from './timeline.js'
import { setupBattleInsightCharts } from './battle_insight_charts.js'
import { deferAfterPaint } from './chart_preloader.js'
import { byteSizeBucket, gaEvent } from './analytics.js'
import { clearFocusedPlayers, getFocusedPlayer, setFocusedPlayer } from './player_focus.js'

import './wasm_exec.js'

if (WebAssembly) {
    const wasmT0 = performance.now();
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("gojs.wasm"), go.importObject)
        .then((result) => {
            go.run(result.instance);
            gaEvent('lux_wasm_ready', {
                wasm_load_ms: Math.round(performance.now() - wasmT0),
            });
            if (timelineCtl && typeof timelineCtl.refresh === 'function') {
                timelineCtl.refresh();
            }
            refreshAll();
        })
        .catch((err) => {
            gaEvent('lux_wasm_error', {
                message: String(err && err.message ? err.message : err).slice(0, 120),
            });
        });
} else {
    console.log("WebAssembly is not supported in your browser");
    gaEvent('lux_wasm_unsupported', {});
}

const pickLogs = document.getElementById('pick_logs');
const matchSelect = document.getElementById('match_select');
const playerFocusSelect = document.getElementById('graph_player_focus_select');
const warriorNickSelect = document.getElementById('warrior_nick_select');
const swapCookieName = 'lux_team_swap_by_match';
const warriorCookieName = 'lux_warrior_nick';
const timelineResetTopBtn = document.getElementById('timeline_reset_top_btn');

let currentMetric = 'damage';

function getSelectedMatchIndex() {
    if (!matchSelect || matchSelect.disabled) {
        return 0;
    }
    return Number(matchSelect.value) || 0;
}

function refreshCharts() {
    ApplyParsedCharts(getSelectedMatchIndex(), currentMetric);
}

function readSwapStateForMatch(matchIndex) {
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
        return map[String(Number(matchIndex) || 0)] === 1;
    } catch (_) {
        return false;
    }
}

function hasExplicitSwapState(matchIndex) {
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
        const key = String(Number(matchIndex) || 0);
        return Object.prototype.hasOwnProperty.call(map, key);
    } catch (_) {
        return false;
    }
}

function selectedTeamIDs(matchIndex, allyTeamID, enemyTeamID) {
    const isSwap = readSwapStateForMatch(matchIndex);
    if (isSwap) {
        return {
            ally: enemyTeamID || allyTeamID || 0,
            enemy: allyTeamID || enemyTeamID || 0,
        };
    }
    return {
        ally: allyTeamID || 0,
        enemy: enemyTeamID || 0,
    };
}

function readCookieValue(name) {
    if (typeof document === 'undefined') {
        return '';
    }
    const row = document.cookie.split('; ').find((x) => x.startsWith(`${name}=`));
    if (!row) {
        return '';
    }
    try {
        return decodeURIComponent(row.slice(name.length + 1));
    } catch (_) {
        return '';
    }
}

function writeCookieValue(name, value, maxAgeSec) {
    if (typeof document === 'undefined') {
        return;
    }
    const v = encodeURIComponent(String(value || ''));
    document.cookie = `${name}=${v}; path=/; max-age=${maxAgeSec}; samesite=lax`;
}

function getWarriorNick() {
    return readCookieValue(warriorCookieName).trim();
}

function setWarriorNick(v) {
    writeCookieValue(warriorCookieName, String(v || '').trim(), 60 * 60 * 24 * 365 * 5);
}

function playerTeamsFromTimeline(matchIndex) {
    const fn = globalThis.getTimelineJSON;
    if (typeof fn !== 'function') {
        return { allyTeamID: 0, enemyTeamID: 0, byPlayer: new Map() };
    }
    try {
        const raw = fn(matchIndex, '');
        const data = JSON.parse(raw || '{}');
        const markers = Array.isArray(data && data.markers) ? data.markers : [];
        const byPlayer = new Map();
        markers.forEach((m) => {
            const t = Number(m && m.team_id) || 0;
            const kt = Number(m && m.killer_team_id) || 0;
            const vt = Number(m && m.victim_team_id) || 0;
            const p = String((m && m.player) || '').trim();
            const k = String((m && m.killer) || '').trim();
            const v = String((m && m.victim) || '').trim();
            if (p && t) byPlayer.set(p, t);
            if (k && kt) byPlayer.set(k, kt);
            if (v && vt) byPlayer.set(v, vt);
        });
        return {
            allyTeamID: Number(data && data.ally_team_id) || 0,
            enemyTeamID: Number(data && data.enemy_team_id) || 0,
            byPlayer,
        };
    } catch (_) {
        return { allyTeamID: 0, enemyTeamID: 0, byPlayer: new Map() };
    }
}

function warriorSwapFromCharts(matchIndex, nick) {
    const fn = globalThis.getChartsJSON;
    if (typeof fn !== 'function') {
        return null;
    }
    try {
        const raw = fn(matchIndex, 'damage', '{}');
        const parsed = JSON.parse(raw || '[]');
        if (!Array.isArray(parsed) || parsed.length < 2) {
            return null;
        }
        const canon = String(nick || '').trim().toLowerCase();
        const inTeam = (arr) => (Array.isArray(arr) ? arr : [])
            .slice(1)
            .some((c) => String(c && c.name ? c.name : '').trim().toLowerCase() === canon);
        if (inTeam(parsed[0])) {
            return false;
        }
        if (inTeam(parsed[1])) {
            return true;
        }
    } catch (_) {
        return null;
    }
    return null;
}

function syncSwapButtonState(swapped) {
    return applyTeamSwapState(getSelectedMatchIndex(), Boolean(swapped));
}

function applyWarriorAllianceForMatch(matchIndex) {
    // Ник задает только дефолт. Если пользователь уже переключал команду в этом матче,
    // больше не переопределяем его выбор.
    if (hasExplicitSwapState(matchIndex)) {
        return false;
    }
    const nick = getWarriorNick();
    if (!nick) {
        return false;
    }
    // Для верхних графиков опираемся на тот же источник (getChartsJSON), чтобы
    // дефолтный swap гарантированно отражался в левой/правой панели.
    const byCharts = warriorSwapFromCharts(matchIndex, nick);
    if (byCharts !== null) {
        return syncSwapButtonState(byCharts);
    }
    const teamMeta = playerTeamsFromTimeline(matchIndex);
    const tid = teamMeta.byPlayer.get(nick) || 0;
    if (!tid) {
        return false;
    }
    const allyID = teamMeta.allyTeamID || 0;
    const enemyID = teamMeta.enemyTeamID || 0;
    if (!allyID || !enemyID || allyID === enemyID) {
        return false;
    }
    if (tid === allyID) {
        return syncSwapButtonState(false);
    } else if (tid === enemyID) {
        return syncSwapButtonState(true);
    }
    return false;
}

function loadWarriorOptions() {
    if (!warriorNickSelect) {
        return;
    }
    const idx = getSelectedMatchIndex();
    const teamMeta = playerTeamsFromTimeline(idx);
    const selected = selectedTeamIDs(idx, teamMeta.allyTeamID, teamMeta.enemyTeamID);
    const allPlayers = Array.from(teamMeta.byPlayer.keys()).sort((a, b) => String(a).localeCompare(String(b)));
    const allies = [];
    const enemies = [];
    const other = [];
    allPlayers.forEach((p) => {
        const tid = teamMeta.byPlayer.get(p) || 0;
        if (selected.ally && tid === selected.ally) {
            allies.push(p);
        } else if (selected.enemy && tid === selected.enemy) {
            enemies.push(p);
        } else {
            other.push(p);
        }
    });
    warriorNickSelect.innerHTML = '<option value="">Не выбрано</option>';
    const appendGroup = (label, arr) => {
        if (!arr.length) {
            return;
        }
        const g = document.createElement('optgroup');
        g.label = label;
        arr.forEach((p) => {
            const o = document.createElement('option');
            o.value = p;
            o.textContent = p;
            g.appendChild(o);
        });
        warriorNickSelect.appendChild(g);
    };
    appendGroup('Союзники', allies);
    appendGroup('Противники', enemies);
    appendGroup('Прочие', other);
    const saved = getWarriorNick();
    warriorNickSelect.value = allPlayers.includes(saved) ? saved : '';
}

function loadPlayerFocusOptions() {
    if (!playerFocusSelect) {
        return;
    }
    const idx = getSelectedMatchIndex();
    let players = [];
    const fn = globalThis.getDamageFilterMetaJSON;
    if (typeof fn === 'function') {
        try {
            const raw = fn(idx, '', '{}');
            const meta = JSON.parse(raw || '{}');
            if (meta && Array.isArray(meta.players)) {
                players = meta.players.slice().sort((a, b) => String(a).localeCompare(String(b)));
            }
        } catch (_) {
            players = [];
        }
    }
    const prev = getFocusedPlayer(idx);
    const warrior = getWarriorNick();
    const defaultPlayer = prev || (players.includes(warrior) ? warrior : '');
    const teamMeta = playerTeamsFromTimeline(idx);
    const selected = selectedTeamIDs(idx, teamMeta.allyTeamID, teamMeta.enemyTeamID);
    const allies = [];
    const enemies = [];
    const other = [];
    players.forEach((p) => {
        const tid = teamMeta.byPlayer.get(p) || 0;
        if (selected.ally && tid === selected.ally) {
            allies.push(p);
        } else if (selected.enemy && tid === selected.enemy) {
            enemies.push(p);
        } else {
            other.push(p);
        }
    });
    playerFocusSelect.innerHTML = '<option value="">Все игроки</option>';
    const appendGroup = (label, arr) => {
        if (!arr.length) {
            return;
        }
        const g = document.createElement('optgroup');
        g.label = label;
        arr.forEach((p) => {
            const o = document.createElement('option');
            o.value = p;
            o.textContent = p;
            g.appendChild(o);
        });
        playerFocusSelect.appendChild(g);
    };
    appendGroup('Союзники', allies);
    appendGroup('Противники', enemies);
    appendGroup('Прочие', other);
    playerFocusSelect.value = players.includes(defaultPlayer) ? defaultPlayer : '';
    if (playerFocusSelect.value !== prev) {
        setFocusedPlayer(idx, playerFocusSelect.value || '');
    }
}

const damagePanel = setupDamageTablePanel(getSelectedMatchIndex);

function escapeHtml(s) {
    return String(s)
        .replace(/&/g, '&amp;')
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;')
        .replace(/"/g, '&quot;');
}

/** Секунды от начала матча → mm:ss.xx как на таймлайне */
function formatMatchSec(sec) {
    if (typeof sec !== 'number' || Number.isNaN(sec)) {
        return '—';
    }
    const m = Math.floor(sec / 60);
    const s = sec - m * 60;
    const mm = String(m).padStart(2, '0');
    const ss = s < 10 ? `0${s.toFixed(2)}` : s.toFixed(2);
    return `${mm}:${ss}`;
}

function updateWatcherBanner() {
    const el = document.getElementById('watcher_banner');
    if (!el) {
        return;
    }
    let meta = [];
    if (typeof getLevelsMetaJSON === 'function') {
        try {
            const raw = getLevelsMetaJSON();
            meta = JSON.parse(raw);
            if (!Array.isArray(meta)) {
                meta = [];
            }
        } catch (_) {
            meta = [];
        }
    }
    const m = meta[getSelectedMatchIndex()];
    if (!m || !m.watcher_active) {
        el.hidden = true;
        el.innerHTML = '';
        return;
    }
    el.hidden = false;
    const c = formatMatchSec(m.watcher_connect_sec);
    const d = formatMatchSec(m.watcher_disconnect_sec);
    const names = Array.isArray(m.watcher_names) && m.watcher_names.length
        ? ` (${m.watcher_names.map(escapeHtml).join(', ')})`
        : '';
    el.innerHTML = `<strong>Большой брат наблюдает за вами!</strong>${names}<br>Подключение наблюдателя: ${c} · Отключение: ${d}`;
}

function refreshAll() {
    const idx = getSelectedMatchIndex();
    // Сначала синхронизируем charts-state для текущего матча (в т.ч. swap-кнопку).
    refreshCharts();
    // Затем применяем дефолт от "Кто ты, воин?" уже к актуальному состоянию матча.
    const autoSwapChanged = applyWarriorAllianceForMatch(idx);
    if (autoSwapChanged) {
        refreshCharts();
    }
    loadWarriorOptions();
    loadPlayerFocusOptions();
    updateWatcherBanner();
    deferAfterPaint(() => {
        if (damagePanel && typeof damagePanel.refresh === 'function') {
            damagePanel.refresh();
        }
        deferAfterPaint(() => {
            if (battleInsightCtl && typeof battleInsightCtl.refresh === 'function') {
                battleInsightCtl.refresh();
            }
        });
    });
}

const timelineCtl = setupTimeline(getSelectedMatchIndex, refreshAll);
const battleInsightCtl = setupBattleInsightCharts(getSelectedMatchIndex);

if (timelineResetTopBtn) {
    timelineResetTopBtn.addEventListener('click', () => {
        if (timelineCtl && typeof timelineCtl.reset === 'function') {
            timelineCtl.reset();
        }
    });
}

function setupMetricButtons() {
    const wrap = document.querySelector('.metric-buttons');
    if (!wrap) {
        return;
    }
    wrap.addEventListener('click', (e) => {
        const btn = e.target.closest('.metric-btn');
        if (!btn || !btn.dataset.metric) {
            return;
        }
        currentMetric = btn.dataset.metric;
        gaEvent('lux_metric_change', { metric: currentMetric });
        wrap.querySelectorAll('.metric-btn').forEach((b) => b.classList.toggle('active', b === btn));
        refreshCharts();
    });
}
setupMetricButtons();

function setupBattleInsightTooltips() {
    document.querySelectorAll('.battle-insight-help[data-bs-toggle="tooltip"]').forEach((el) => {
        if (bootstrap.Tooltip.getInstance(el)) {
            return;
        }
        new bootstrap.Tooltip(el, {
            container: 'body',
            trigger: 'hover focus',
        });
    });
}
setupBattleInsightTooltips();

if (warriorNickSelect) {
    warriorNickSelect.addEventListener('change', () => {
        setWarriorNick(warriorNickSelect.value || '');
        applyWarriorAllianceForMatch(getSelectedMatchIndex());
        loadWarriorOptions();
        loadPlayerFocusOptions();
        refreshAll();
    });
}

function renderMatchOptions() {
    if (!matchSelect) {
        return;
    }
    let meta = [];
    if (typeof getLevelsMetaJSON === 'function') {
        try {
            const raw = getLevelsMetaJSON();
            meta = JSON.parse(raw);
            if (!Array.isArray(meta)) {
                meta = [];
            }
        } catch (e) {
            console.warn('getLevelsMetaJSON', e);
            meta = [];
        }
    }
    let levels = meta.length;
    if (levels === 0 && typeof showLevels === 'function') {
        levels = Number(showLevels()) || 0;
    }
    matchSelect.innerHTML = "";
    if (levels <= 0) {
        const opt = document.createElement('option');
        opt.value = "0";
        opt.textContent = "Match 1";
        matchSelect.appendChild(opt);
        matchSelect.disabled = true;
        if (timelineCtl && typeof timelineCtl.refresh === 'function') {
            timelineCtl.refresh();
        }
        refreshAll();
        return;
    }
    for (let i = 0; i < levels; i++) {
        const opt = document.createElement('option');
        opt.value = String(i);
        const m = meta[i];
        if (m && m.label) {
            opt.textContent = m.label;
        } else {
            opt.textContent = `Match ${i + 1}`;
        }
        if (m) {
            const parts = [];
            if (m.game_mode) {
                parts.push(m.game_mode);
            }
            if (m.map_name) {
                parts.push(m.map_name);
            }
            if (m.session_id) {
                parts.push(`session ${m.session_id}`);
            }
            if (m.start_time) {
                parts.push(m.start_time);
            }
            opt.title = parts.join(' · ');
        }
        matchSelect.appendChild(opt);
    }
    matchSelect.disabled = false;
    matchSelect.value = "0";
    if (timelineCtl && typeof timelineCtl.refresh === 'function') {
        timelineCtl.refresh();
    }
    refreshAll();
}

if (matchSelect) {
    matchSelect.addEventListener('change', () => {
        gaEvent('lux_match_change', {
            match_index: Number(matchSelect.value) || 0,
        });
        if (timelineCtl && typeof timelineCtl.refresh === 'function') {
            timelineCtl.refresh();
        }
        refreshAll();
    });
}

if (playerFocusSelect) {
    playerFocusSelect.addEventListener('change', () => {
        const idx = getSelectedMatchIndex();
        const player = playerFocusSelect.value || '';
        setFocusedPlayer(idx, player);
        gaEvent('lux_player_focus_change', { match_index: idx, has_player: player ? '1' : '0' });
        if (timelineCtl && typeof timelineCtl.refresh === 'function') {
            timelineCtl.refresh();
        }
        if (battleInsightCtl && typeof battleInsightCtl.refresh === 'function') {
            battleInsightCtl.refresh();
        }
    });
}

if (document.getElementById('pieChart1') && document.getElementById('pieChart2')) {
    CreateCharts();
    setupGraphViewToolbar(refreshCharts);
}

function runParseAndRenderLogs(data) {
    const tAll = performance.now();
    if (typeof parseFiles !== 'function') {
        gaEvent('lux_parse_error', { reason: 'no_parseFiles' });
        return;
    }
    const tParse = performance.now();
    try {
        parseFiles(data.rawGame, data.rawCombat);
    } catch (err) {
        gaEvent('lux_parse_error', {
            message: String(err && err.message ? err.message : err).slice(0, 120),
        });
        return;
    }
    const logParseMs = Math.round(performance.now() - tParse);
    const tMatch = performance.now();
    renderMatchOptions();
    const matchListMs = Math.round(performance.now() - tMatch);
    const n = matchSelect && matchSelect.options ? matchSelect.options.length : 0;
    gaEvent('lux_logs_parsed', {
        log_parse_ms: logParseMs,
        match_list_ms: matchListMs,
        parse_total_ms: Math.round(performance.now() - tAll),
        game_size_bucket: byteSizeBucket(data.rawGame && data.rawGame.length),
        combat_size_bucket: byteSizeBucket(data.rawCombat && data.rawCombat.length),
        match_count: n,
    });
}

pickLogs.addEventListener('change', function () {
    clearTeamSwapCookieState();
    clearFocusedPlayers();
    if (playerFocusSelect) {
        playerFocusSelect.innerHTML = '<option value="">Все игроки</option>';
    }
    const data = {
        rawGame: '',
        rawCombat: '',
    };
    const files = this.files;
    console.log('called', files);

    if (files && files.length) {
        gaEvent('lux_log_folder_selected', { entry_count: files.length });
    }

    for (let f = 0; f < files.length; f++) {
        const file = files[f];

        if (file.name === 'combat.log') {
            const reader = new FileReader();
            reader.onload = function (e) {
                data.rawCombat = e.target.result;
                console.log('combat.log loaded', data.rawCombat.length);
                if (data.rawGame.length) {
                    runParseAndRenderLogs(data);
                }
            }.bind(this);
            reader.onerror = function () {
                console.log('combat error');
                gaEvent('lux_file_read_error', { which: 'combat' });
            };
            reader.readAsBinaryString(file);
        } else if (file.name === 'game.log') {
            console.log('game.log loaded');
            const reader = new FileReader();
            reader.onload = function (e) {
                data.rawGame = e.target.result;
                console.log('game.log loaded', data.rawGame.length);
                if (data.rawCombat.length) {
                    runParseAndRenderLogs(data);
                }
            }.bind(this);
            reader.onerror = function () {
                console.log('game error');
                gaEvent('lux_file_read_error', { which: 'game' });
            };
            reader.readAsBinaryString(file);
        }
    }
});