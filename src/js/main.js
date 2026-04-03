// Импортируйте наш пользовательский CSS
import '../scss/styles.scss'

// Импортируйте весь JS Bootstrap
import * as bootstrap from 'bootstrap'

import { CreateCharts, ApplyParsedCharts, setupGraphViewToolbar } from './charts.js'
import { setupDamageTablePanel } from './damage_table.js'
import { setupTimeline } from './timeline.js'
import { setupBattleInsightCharts } from './battle_insight_charts.js'

import './wasm_exec.js'

if (WebAssembly) {
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("gojs.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
        if (timelineCtl && typeof timelineCtl.refresh === 'function') {
            timelineCtl.refresh();
        }
        refreshAll();
    });
} else {
    console.log("WebAssembly is not supported in your browser")
}

const pickLogs = document.getElementById('pick_logs');
const matchSelect = document.getElementById('match_select');

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
    refreshCharts();
    updateWatcherBanner();
    if (damagePanel && typeof damagePanel.refresh === 'function') {
        damagePanel.refresh();
    }
    if (battleInsightCtl && typeof battleInsightCtl.refresh === 'function') {
        battleInsightCtl.refresh();
    }
}

const timelineCtl = setupTimeline(getSelectedMatchIndex, refreshAll);
const battleInsightCtl = setupBattleInsightCharts(getSelectedMatchIndex);

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
        if (timelineCtl && typeof timelineCtl.refresh === 'function') {
            timelineCtl.refresh();
        }
        refreshAll();
    });
}

if (document.getElementById('pieChart1') && document.getElementById('pieChart2')) {
    CreateCharts();
    setupGraphViewToolbar(refreshCharts);
}

pickLogs.addEventListener('change', function () {
    var data = {
        rawGame: "",
        rawCombat: "",
    }
    const files = this.files
    console.log("called", files);

    for (let f = 0; f < files.length; f++) {
        const file = files[f];

        if (file.name === 'combat.log') {
            const reader = new FileReader();
            reader.onload = function (e) {
                data.rawCombat = e.target.result;
                console.log('combat.log loaded', data.rawCombat.length);
                if (data.rawGame.length) {
                    parseFiles(data.rawGame, data.rawCombat);
                    renderMatchOptions();
                }
            }.bind(this);
            reader.onerror = function (stuff) { console.log('combat error', stuff); };
            reader.readAsBinaryString(file);
        } else if (file.name === 'game.log') {
            console.log('game.log loaded');
            const reader = new FileReader();
            reader.onload = function (e) {
                data.rawGame = e.target.result;
                console.log('combat.log loaded', data.rawGame.length);
                if (data.rawCombat.length) {
                    parseFiles(data.rawGame, data.rawCombat);
                    renderMatchOptions();
                }
            }.bind(this);
            reader.onerror = function (stuff) { console.log('game error', stuff); };
            reader.readAsBinaryString(file);
        }
    }
})