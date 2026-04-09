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
import {
    getLevelsMetaParsedCached,
    getPlayerTeamsFromTimelineCached,
    invalidateClientWasmCaches,
    primeTimelineEmptyFocus,
} from './client_wasm_cache.js'

import './wasm_exec.js'

const demoGameLogURL = new URL('../assets/2026.04.03 22.23.53.889/game.log', import.meta.url);
const demoCombatLogURL = new URL('../assets/2026.04.03 22.23.53.889/combat.log', import.meta.url);

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
const loadDemoLogsBtn = document.getElementById('load_demo_logs_btn');
const matchSelect = document.getElementById('match_select');
const playerFocusSelect = document.getElementById('graph_player_focus_select');
const warriorNickSelect = document.getElementById('warrior_nick_select');
const simpleUiModeCb = document.getElementById('simple_ui_mode_cb');
const swapCookieName = 'lux_team_swap_by_match';
const warriorCookieName = 'lux_warrior_nick';
const simpleUiCookieName = 'lux_simple_ui_mode';
const timelineResetTopBtn = document.getElementById('timeline_reset_top_btn');

let currentMetric = 'damage';
const intFmt = new Intl.NumberFormat('ru-RU');

function getSelectedMatchIndex() {
    if (!matchSelect || matchSelect.disabled) {
        return 0;
    }
    return Number(matchSelect.value) || 0;
}

function refreshCharts() {
    ApplyParsedCharts(getSelectedMatchIndex(), currentMetric);
}

function formatSummaryDuration(sec) {
    const n = Number(sec) || 0;
    const m = Math.floor(n / 60);
    const s = Math.floor(n % 60);
    return `${String(m).padStart(2, '0')}:${String(s).padStart(2, '0')}`;
}

function updateMatchSummary() {
    const card = document.getElementById('match_summary_card');
    const grid = document.getElementById('match_summary_grid');
    if (!card) {
        return;
    }
    const fn = globalThis.getMatchSummaryJSON;
    if (typeof fn !== 'function') {
        card.hidden = true;
        return;
    }
    const idx = getSelectedMatchIndex();
    let data = null;
    try {
        const raw = fn(idx);
        data = JSON.parse(raw || '{}');
    } catch (_) {
        card.hidden = true;
        return;
    }
    if (!data || typeof data !== 'object') {
        card.hidden = true;
        return;
    }
    const map = String(data.map_name || '').trim();
    const mode = String(data.game_mode || '').trim();
    const meta = [mode, map, formatSummaryDuration(data.duration_sec)].filter(Boolean).join(' · ');
    const setText = (id, v) => {
        const el = document.getElementById(id);
        if (el) {
            el.textContent = v;
        }
    };
    const setTeamPanelTitles = (left, right) => {
        const titles = document.querySelectorAll('[data-team-panel] .graph-box-title');
        if (!titles || titles.length < 2) {
            return;
        }
        titles[0].textContent = String(left || 'Team 1');
        titles[1].textContent = String(right || 'Team 2');
    };
    const shortRole = (s) => {
        const v = String(s || '').toLowerCase();
        if (v.includes('союз')) return 'Союзники';
        if (v.includes('враг')) return 'Враги';
        return '';
    };
    const shortSpawn = (s) => {
        const v = String(s || '').toLowerCase();
        if (v.includes('левый спавн')) return 'левый спавн';
        if (v.includes('правый спавн')) return 'правый спавн';
        return '';
    };
    const panelTitle = (team, fallback) => {
        if (!team || typeof team !== 'object') {
            return fallback;
        }
        const name = String(team.name || fallback || '').trim();
        const role = shortRole(team.role);
        const spawn = shortSpawn(team.role);
        const parts = [name, role, spawn].filter(Boolean);
        return parts.join(' · ') || fallback;
    };
    setText('match_summary_meta', meta);
    const teams = Array.isArray(data.teams) ? data.teams : [];
    if (grid) {
        const safe = (s) => String(s || '')
            .replace(/&/g, '&amp;')
            .replace(/</g, '&lt;')
            .replace(/>/g, '&gt;')
            .replace(/"/g, '&quot;');
        const rows = teams.length ? teams : [
            {
                name: String(data.team_left_name || 'Team 1'),
                role: 'Союзники',
                damage: Number(data.damage_left) || 0,
                heal: Number(data.heal_left) || 0,
                kills: Number(data.kills_left) || 0,
            },
            {
                name: String(data.team_right_name || 'Team 2'),
                role: 'Враги',
                damage: Number(data.damage_right) || 0,
                heal: Number(data.heal_right) || 0,
                kills: Number(data.kills_right) || 0,
            },
        ];
        grid.innerHTML = rows.map((t) => `
            <div class="match-summary-team">
                <div class="match-summary-team-name">${safe(t.name || 'Team')}</div>
                <div>Урон: <b>${intFmt.format(Math.round(Number(t.damage) || 0))}</b></div>
                <div>Лечение: <b>${intFmt.format(Math.round(Number(t.heal) || 0))}</b></div>
                <div>Киллы: <b>${String(Math.round(Number(t.kills) || 0))}</b></div>
            </div>
        `).join('');
        if (teams.length) {
            const ally = teams.find((t) => String(t.role || '').toLowerCase().includes('союз')) || teams[0];
            const enemy = teams.find((t) => t !== ally) || teams[1] || null;
            setTeamPanelTitles(panelTitle(ally, 'Team 1'), panelTitle(enemy, 'Team 2'));
        } else {
            setTeamPanelTitles(String(data.team_left_name || 'Team 1'), String(data.team_right_name || 'Team 2'));
        }
    }
    card.hidden = false;
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

function isSimpleUiMode() {
    return readCookieValue(simpleUiCookieName) === '1';
}

function applySimpleUiMode(enabled) {
    const on = Boolean(enabled);
    if (typeof document !== 'undefined' && document.body) {
        document.body.classList.toggle('simple-ui-mode', on);
    }
    if (simpleUiModeCb) {
        simpleUiModeCb.checked = on;
    }
}

function setSimpleUiMode(enabled) {
    writeCookieValue(simpleUiCookieName, enabled ? '1' : '0', 60 * 60 * 24 * 365 * 5);
    applySimpleUiMode(enabled);
}

function playerTeamsFromTimeline(matchIndex) {
    return getPlayerTeamsFromTimelineCached(matchIndex);
}

function warriorSwapFromRoster(matchIndex, nick) {
    const fn = globalThis.getTeamPanelsRosterJSON;
    if (typeof fn !== 'function') {
        return null;
    }
    try {
        const raw = fn(matchIndex);
        const parsed = JSON.parse(raw || '{}');
        const canon = String(nick || '').trim().toLowerCase();
        const inNames = (arr) => (Array.isArray(arr) ? arr : []).some(
            (n) => String(n || '').trim().toLowerCase() === canon,
        );
        if (inNames(parsed.left)) {
            return false;
        }
        if (inNames(parsed.right)) {
            return true;
        }
    } catch (_) {
        return null;
    }
    return null;
}

function warriorSwapFromCharts(matchIndex, nick) {
    const fromRoster = warriorSwapFromRoster(matchIndex, nick);
    if (fromRoster !== null) {
        return fromRoster;
    }
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
    // Сначала roster панелей графиков (легкий JSON), иначе тот же источник, что и графики (damage).
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
    playerFocusSelect.innerHTML = '<option value="">Игрок не выбран</option>';
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
    const meta = getLevelsMetaParsedCached();
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
    primeTimelineEmptyFocus(idx);
    // Графики без ручного swap/warrior-переопределения: союзники определяются по localClientTeamID.
    refreshCharts();
    loadPlayerFocusOptions();
    updateMatchSummary();
    updateWatcherBanner();
    const simpleUi = isSimpleUiMode();
    deferAfterPaint(() => {
        if (!simpleUi && damagePanel && typeof damagePanel.refresh === 'function') {
            damagePanel.refresh();
        }
        deferAfterPaint(() => {
            if (!simpleUi && battleInsightCtl && typeof battleInsightCtl.refresh === 'function') {
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

applySimpleUiMode(isSimpleUiMode());
if (simpleUiModeCb) {
    simpleUiModeCb.addEventListener('change', () => {
        setSimpleUiMode(!!simpleUiModeCb.checked);
        gaEvent('lux_simple_ui_toggle', { enabled: simpleUiModeCb.checked ? '1' : '0' });
        refreshAll();
    });
}

if (warriorNickSelect) {
    warriorNickSelect.closest('.warrior-picker')?.setAttribute('hidden', 'hidden');
}

function renderMatchOptions() {
    if (!matchSelect) {
        return;
    }
    const meta = getLevelsMetaParsedCached();
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

if (document.getElementById('graphs_root')) {
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
    invalidateClientWasmCaches();
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

function prepareForNewLogs() {
    clearTeamSwapCookieState();
    clearFocusedPlayers();
    if (playerFocusSelect) {
        playerFocusSelect.innerHTML = '<option value="">Игрок не выбран</option>';
    }
}

async function loadDemoLogs() {
    try {
        const [gameRes, combatRes] = await Promise.all([fetch(demoGameLogURL), fetch(demoCombatLogURL)]);
        if (!gameRes.ok || !combatRes.ok) {
            gaEvent('lux_parse_error', {
                reason: 'demo_logs_not_found',
                game_status: String(gameRes.status || 0),
                combat_status: String(combatRes.status || 0),
            });
            return;
        }
        const [rawGame, rawCombat] = await Promise.all([gameRes.text(), combatRes.text()]);
        if (!rawGame || !rawCombat) {
            gaEvent('lux_parse_error', { reason: 'demo_logs_empty' });
            return;
        }
        prepareForNewLogs();
        runParseAndRenderLogs({ rawGame, rawCombat });
        gaEvent('lux_demo_logs_loaded', { source: '2026.04.03 22.23.53.889' });
    } catch (err) {
        gaEvent('lux_parse_error', {
            reason: 'demo_logs_fetch_failed',
            message: String(err && err.message ? err.message : err).slice(0, 120),
        });
    }
}


pickLogs.addEventListener('click', function () {
    // Позволяет выбрать тот же самый файл/папку и снова получить событие change.
    this.value = '';
});

if (loadDemoLogsBtn) {
    loadDemoLogsBtn.addEventListener('click', () => {
        loadDemoLogs();
    });
}
pickLogs.addEventListener('change', function () {
    prepareForNewLogs();
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
    // На некоторых браузерах повторный выбор той же папки не триггерит change без явного сброса.
    this.value = '';
});