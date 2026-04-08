/**
 * Match timeline: zoom (view window), markers, WASM time range = окно зума.
 */

import { deferAfterPaint, hideChartPreloader, showChartPreloader } from './chart_preloader.js';
import { gaEvent } from './analytics.js';
import { getFocusedPlayer } from './player_focus.js';
import { getTimelineEmptyFocusParsed } from './client_wasm_cache.js';

const MIN_VIEW_SPAN_SEC = 0.4;

let matchDurationSec = 0;
let viewStart = 0;
let viewEnd = 0;
let timelineReady = false;
let timelineLoadGen = 0;
let matchStartUnixMs = 0;

let rootEl = null;
let trackEl = null;
let brushSurfaceEl = null;
let brushPreviewEl = null;
let markersLayerEl = null;
let hintEl = null;
let legendEl = null;
let markerLabelsLayerEl = null;
let timeAxisEl = null;
let resetBtnEl = null;
let tooltipEl = null;

let dragMode = null; // 'left' | 'right' | 'move' | 'brush' | null
let brushAnchorSec = 0;
let lastMarkers = [];
let rangeChangeCb = null;
/** Обновление таблицы combat-логов под выбранный интервал (ставится в setupTimeline). */
let combatLogRefresh = null;
let getMatchIndexRef = null;
const swapCookieName = 'lux_team_swap_by_match';
let teamsSwapped = false;

/** Общий курсор времени (сек от начала матча) для синхронизации с графиками «ситуация в бою». */
let syncCursorSec = null;
const syncCursorListeners = new Set();

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

function displayKind(kind) {
    if (!teamsSwapped) {
        return kind;
    }
    if (kind === 'spawn') return 'enemy_spawn';
    if (kind === 'enemy_spawn') return 'spawn';
    if (kind === 'kill') return 'death';
    if (kind === 'death') return 'kill';
    return kind;
}

function updateLegend() {
    if (!legendEl) {
        return;
    }
    const spawnLabel = teamsSwapped ? 'Спавн (враги)' : 'Спавн (союзники)';
    const enemySpawnLabel = teamsSwapped ? 'Спавн (союзники)' : 'Спавн (враги)';
    const killLabel = teamsSwapped ? 'Убийство (союзник)' : 'Убийство (враг)';
    const deathLabel = teamsSwapped ? 'Смерть союзника' : 'Смерть противника';
    legendEl.innerHTML = ''
        + `<span class="timeline-legend-item"><i class="timeline-dot spawn"></i> ${spawnLabel}</span>`
        + `<span class="timeline-legend-item"><i class="timeline-dot enemy-spawn"></i> ${enemySpawnLabel}</span>`
        + `<span class="timeline-legend-item"><i class="timeline-dot kill"></i> ${killLabel}</span>`
        + `<span class="timeline-legend-item"><i class="timeline-dot death"></i> ${deathLabel}</span>`;
}

function formatSec(sec) {
    if (typeof sec !== 'number' || Number.isNaN(sec)) {
        return '—';
    }
    const m = Math.floor(sec / 60);
    const s = sec - m * 60;
    const mm = String(m).padStart(2, '0');
    const ss = s < 10 ? `0${s.toFixed(2)}` : s.toFixed(2);
    return `${mm}:${ss}`;
}

function formatClockSec(sec) {
    if (typeof sec !== 'number' || Number.isNaN(sec)) {
        return '—';
    }
    if (!matchStartUnixMs || !Number.isFinite(matchStartUnixMs)) {
        return formatSec(sec);
    }
    const d = new Date(matchStartUnixMs + sec * 1000);
    return d.toLocaleTimeString('ru-RU', {
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false,
    });
}

function clamp(v, lo, hi) {
    return Math.min(Math.max(v, lo), hi);
}

function truncate(str, maxLen) {
    const s = String(str || '');
    if (s.length <= maxLen) {
        return s;
    }
    if (maxLen <= 1) {
        return '…';
    }
    return `${s.slice(0, maxLen - 1)}…`;
}

function viewSpanSec() {
    return Math.max(viewEnd - viewStart, 1e-6);
}

/** Шаг между засечками шкалы (сек), «красивый» для текущего масштаба. */
function niceTimeAxisStep(spanSec, approxTicks) {
    if (spanSec <= 0 || approxTicks < 1) {
        return 0.1;
    }
    const raw = spanSec / approxTicks;
    if (!Number.isFinite(raw) || raw <= 0) {
        return 0.05;
    }
    const exp = Math.floor(Math.log10(raw));
    const base = Math.pow(10, exp);
    const f = raw / base;
    let nice;
    if (f <= 1) {
        nice = 1;
    } else if (f <= 2) {
        nice = 2;
    } else if (f <= 5) {
        nice = 5;
    } else {
        nice = 10;
    }
    return Math.max(nice * base, 1e-6);
}

function renderTimeAxis() {
    if (!timeAxisEl || !trackEl) {
        return;
    }
    if (!matchDurationSec || matchDurationSec <= 0 || viewEnd <= viewStart) {
        timeAxisEl.innerHTML = '';
        timeAxisEl.style.display = 'none';
        return;
    }
    timeAxisEl.style.display = 'block';
    const widthPx = trackEl.getBoundingClientRect().width;
    const approxTicks = clamp(Math.floor(widthPx / 72), 4, 12);
    const span = viewSpanSec();
    const step = niceTimeAxisStep(span, approxTicks);
    let t0 = Math.ceil(viewStart / step) * step;
    if (t0 < viewStart - 1e-9) {
        t0 += step;
    }
    timeAxisEl.innerHTML = '';
    let n = 0;
    for (let t = t0; t <= viewEnd + 1e-9 && n < 48; t += step, n++) {
        const pct = ((t - viewStart) / span) * 100;
        if (pct < -0.5 || pct > 100.5) {
            continue;
        }
        const tick = document.createElement('div');
        tick.className = 'timeline-time-tick';
        tick.style.left = `${clamp(pct, 0, 100)}%`;
        const line = document.createElement('span');
        line.className = 'timeline-time-tick-line';
        line.setAttribute('aria-hidden', 'true');
        const label = document.createElement('span');
        label.className = 'timeline-time-tick-label';
        label.textContent = formatClockSec(t);
        tick.appendChild(line);
        tick.appendChild(label);
        timeAxisEl.appendChild(tick);
    }
    if (timeAxisEl.childElementCount === 0 && span > 1e-9) {
        const ends = [
            { t: viewStart, pct: 0 },
            { t: viewEnd, pct: 100 },
        ];
        for (let i = 0; i < ends.length; i++) {
            const { t, pct } = ends[i];
            const tick = document.createElement('div');
            tick.className = 'timeline-time-tick';
            tick.style.left = `${pct}%`;
            const line = document.createElement('span');
            line.className = 'timeline-time-tick-line';
            line.setAttribute('aria-hidden', 'true');
            const label = document.createElement('span');
            label.className = 'timeline-time-tick-label';
            label.textContent = formatClockSec(t);
            tick.appendChild(line);
            tick.appendChild(label);
            timeAxisEl.appendChild(tick);
        }
    }
}

/** Вторая строка тултипа: оружие / ассисты (данные из combat). */
function killDeathTooltipExtra(m) {
    if (!m || typeof m !== 'object') {
        return '';
    }
    const w = (m.weapon || '').trim();
    const assists = Array.isArray(m.assists)
        ? m.assists.map((x) => String(x || '').trim()).filter(Boolean)
        : [];
    const parts = [];
    if (w) {
        parts.push(`Оружие: ${w}`);
    }
    if (assists.length) {
        parts.push(`Помогали: ${assists.join(', ')}`);
    }
    return parts.join(' · ');
}

function formatMarkerTooltip(m) {
    if (!m || typeof m !== 'object') {
        return '';
    }
    const tsec = typeof m.time_sec === 'number' ? m.time_sec : null;
    const tline = tsec === null ? '' : `Время: ${formatClockSec(tsec)}\n`;
    if (m.kind === 'spawn') {
        const dKind = displayKind(m.kind);
        const n = (m.player || '').trim();
        const sh = (m.player_ship || '').trim();
        const side = dKind === 'spawn' ? 'союзники' : 'враги';
        if (n && sh) {
            return `${tline}Спавн (${side}): ${n} · ${sh}`;
        }
        if (n) {
            return `${tline}Спавн (${side}): ${n}`;
        }
        return `${tline}Спавн (${side})`;
    }
    if (m.kind === 'enemy_spawn') {
        const dKind = displayKind(m.kind);
        const n = (m.player || '').trim();
        const sh = (m.player_ship || '').trim();
        const side = dKind === 'spawn' ? 'союзники' : 'враги';
        if (n && sh) {
            return `${tline}Спавн (${side}): ${n} · ${sh}`;
        }
        if (n) {
            return `${tline}Спавн (${side}): ${n}`;
        }
        return `${tline}Спавн (${side})`;
    }
    if (m.kind === 'kill') {
        const dKind = displayKind(m.kind);
        const ks = (m.killer_ship || '').trim();
        const vs = (m.victim_ship || '').trim();
        const kPart = m.killer
            ? (ks ? `${m.killer} (${ks})` : m.killer)
            : '';
        const vPart = m.victim
            ? (vs ? `${m.victim} (${vs})` : m.victim)
            : '';
        let line1 = '';
        if (kPart && vPart) {
            line1 = dKind === 'kill'
                ? `Убийство: ${kPart} → ${vPart}`
                : `Смерть союзника: ${kPart} → ${vPart}`;
        } else {
            line1 = dKind === 'kill'
                ? `Убийство: ${kPart || vPart}`
                : `Смерть союзника: ${kPart || vPart}`;
        }
        const extra = killDeathTooltipExtra(m);
        const body = extra ? `${line1}\n${extra}` : line1;
        return `${tline}${body}`;
    }
    if (m.kind === 'death') {
        const dKind = displayKind(m.kind);
        const ks = (m.killer_ship || '').trim();
        const vs = (m.victim_ship || '').trim();
        const kPart = m.killer
            ? (ks ? `${m.killer} (${ks})` : m.killer)
            : '';
        const vName = (m.victim || m.player || '').trim();
        const vPart = vName
            ? (vs ? `${vName} (${vs})` : vName)
            : '';
        let line1 = '';
        if (kPart && vPart) {
            line1 = dKind === 'death'
                ? `Смерть союзника: ${kPart} → ${vPart}`
                : `Убийство: ${kPart} → ${vPart}`;
        } else if (vPart) {
            line1 = dKind === 'death' ? `Смерть союзника: ${vPart}` : `Убийство: ${vPart}`;
        } else {
            line1 = dKind === 'death' ? 'Смерть союзника' : 'Убийство';
        }
        const extra = killDeathTooltipExtra(m);
        const body = extra ? `${line1}\n${extra}` : line1;
        return `${tline}${body}`;
    }
    return `${tline}${m.label || ''}`;
}

function shortLabelUnderMarker(m) {
    if (!m || typeof m !== 'object') {
        return '';
    }
    if (m.kind === 'spawn' || m.kind === 'enemy_spawn') {
        return truncate((m.player || '').trim(), 16);
    }
    if (m.kind === 'kill') {
        const a = truncate((m.killer || '').trim(), 9);
        const b = truncate((m.victim || '').trim(), 9);
        if (a && b) {
            return `${a}→${b}`;
        }
        return a || b || '';
    }
    if (m.kind === 'death') {
        const a = truncate((m.killer || '').trim(), 9);
        const b = truncate((m.victim || m.player || '').trim(), 9);
        if (a && b) {
            return `${a}→${b}`;
        }
        return truncate((m.victim || m.player || '').trim(), 16);
    }
    return '';
}

/** First line when collapsing: one primary nickname */
function firstDisplayNickname(m) {
    if (!m) {
        return '';
    }
    if (m.kind === 'spawn' || m.kind === 'enemy_spawn') {
        return truncate((m.player || '').trim(), 24);
    }
    if (m.kind === 'kill') {
        return truncate((m.killer || '').trim(), 24);
    }
    if (m.kind === 'death') {
        const kv = `${(m.killer || '').trim()}→${(m.victim || m.player || '').trim()}`;
        if (kv.length > 2) {
            return truncate(kv, 24);
        }
        return truncate((m.victim || m.player || '').trim(), 24);
    }
    return '';
}

function absSecFromClientX(ev, rect) {
    const span = viewSpanSec();
    const x = clamp(ev.clientX - rect.left, 0, rect.width);
    return viewStart + (x / rect.width) * span;
}

function ensureTooltipEl() {
    if (tooltipEl) {
        return tooltipEl;
    }
    tooltipEl = document.createElement('div');
    tooltipEl.className = 'timeline-tooltip';
    tooltipEl.setAttribute('role', 'tooltip');
    document.body.appendChild(tooltipEl);
    return tooltipEl;
}

function showTooltip(clientX, clientY, text) {
    if (!text) {
        return;
    }
    const el = ensureTooltipEl();
    el.textContent = text;
    el.style.display = 'block';
    const pad = 8;
    const tw = el.offsetWidth;
    const th = el.offsetHeight;
    let left = clientX - tw / 2;
    let top = clientY + pad;
    if (trackEl) {
        const rect = trackEl.getBoundingClientRect();
        // Тултип всегда под графиком, чтобы не перекрывать сами линии/маркеры.
        top = rect.bottom + 10;
        left = clamp(clientX - tw / 2, 8, window.innerWidth - tw - 8);
    } else {
        if (left + tw > window.innerWidth - 8) {
            left = clientX - tw - pad;
        }
        if (top + th > window.innerHeight - 8) {
            top = clientY - th - pad;
        }
    }
    el.style.left = `${Math.max(8, left)}px`;
    el.style.top = `${Math.max(8, top)}px`;
}

function hideTooltip() {
    if (tooltipEl) {
        tooltipEl.style.display = 'none';
    }
}

function bindMarkerHover(el, text) {
    el.addEventListener('pointerenter', (e) => {
        showTooltip(e.clientX, e.clientY, text);
    });
    el.addEventListener('pointermove', (e) => {
        if (tooltipEl && tooltipEl.style.display === 'block') {
            showTooltip(e.clientX, e.clientY, text);
        }
    });
    el.addEventListener('pointerleave', () => {
        hideTooltip();
    });
}

export function getTimeRangeJSON() {
    if (!timelineReady || matchDurationSec <= 0) {
        return '{}';
    }
    return JSON.stringify({
        time_from_sec: viewStart,
        time_to_sec: viewEnd,
    });
}

export function getTimeRangeBounds() {
    if (!timelineReady || matchDurationSec <= 0) {
        return { from: 0, to: 0 };
    }
    return { from: viewStart, to: viewEnd };
}

export function setTimelineCursorSec(sec) {
    if (sec !== null && (typeof sec !== 'number' || Number.isNaN(sec))) {
        return;
    }
    syncCursorSec = sec;
    layoutSyncCursorLine();
    syncCursorListeners.forEach((fn) => {
        try {
            fn(sec);
        } catch (_) {
            /* ignore */
        }
    });
}

export function getTimelineCursorSec() {
    return syncCursorSec;
}

export function subscribeTimelineCursor(fn) {
    if (typeof fn !== 'function') {
        return () => {};
    }
    syncCursorListeners.add(fn);
    return () => syncCursorListeners.delete(fn);
}

let cursorLineEl = null;

function layoutSyncCursorLine() {
    if (!cursorLineEl || !trackEl) {
        return;
    }
    if (syncCursorSec === null || !timelineReady || matchDurationSec <= 0 || viewEnd <= viewStart) {
        cursorLineEl.style.display = 'none';
        return;
    }
    const sec = clamp(syncCursorSec, viewStart, viewEnd);
    const span = viewSpanSec();
    const pct = ((sec - viewStart) / span) * 100;
    cursorLineEl.style.display = 'block';
    cursorLineEl.style.left = `${pct}%`;
}

function setHint() {
    if (!hintEl) {
        return;
    }
    if (!timelineReady || matchDurationSec <= 0) {
        hintEl.textContent = '';
        return;
    }
    hintEl.textContent = `Окно: ${formatClockSec(viewStart)} — ${formatClockSec(viewEnd)} · длительность ${formatSec(matchDurationSec)}`;
}

function layoutBrushPreview(fromSec, toSec) {
    if (!brushPreviewEl || !trackEl) {
        return;
    }
    const span = viewSpanSec();
    const a = clamp(Math.min(fromSec, toSec), viewStart, viewEnd);
    const b = clamp(Math.max(fromSec, toSec), viewStart, viewEnd);
    const p0 = ((a - viewStart) / span) * 100;
    const p1 = ((b - viewStart) / span) * 100;
    const left = clamp(Math.min(p0, p1), 0, 100);
    const w = Math.abs(p1 - p0);
    brushPreviewEl.style.display = 'block';
    brushPreviewEl.style.left = `${left}%`;
    brushPreviewEl.style.width = `${w}%`;
}

function hideBrushPreview() {
    if (brushPreviewEl) {
        brushPreviewEl.style.display = 'none';
    }
}

function visibleMarkersInView() {
    const out = [];
    for (let i = 0; i < lastMarkers.length; i++) {
        const m = lastMarkers[i];
        const t = typeof m.time_sec === 'number' ? m.time_sec : 0;
        if (t >= viewStart - 1e-9 && t <= viewEnd + 1e-9) {
            out.push(m);
        }
    }
    return out;
}

/** Рисуем «убийство» раньше, «смерть» последней — иначе жёлтая полоска полностью перекрывает красную/фиолетовую при том же time_sec. */
const markerPaintOrder = {
    spawn: 0,
    enemy_spawn: 1,
    kill: 2,
    death: 3,
};

function sortMarkersForPaint(markers) {
    return markers.slice().sort((a, b) => {
        const ta = typeof a.time_sec === 'number' ? a.time_sec : 0;
        const tb = typeof b.time_sec === 'number' ? b.time_sec : 0;
        if (ta !== tb) {
            return ta - tb;
        }
        const pa = markerPaintOrder[a.kind] ?? 10;
        const pb = markerPaintOrder[b.kind] ?? 10;
        if (pa !== pb) {
            return pa - pb;
        }
        return String(a.kind || '').localeCompare(String(b.kind || ''));
    });
}

function markerKindCounts(markers) {
    const counts = {
        spawn: 0,
        enemy_spawn: 0,
        kill: 0,
        death: 0,
        other: 0,
    };
    (Array.isArray(markers) ? markers : []).forEach((m) => {
        const k = String((m && m.kind) || '');
        if (k in counts) {
            counts[k] += 1;
        } else {
            counts.other += 1;
        }
    });
    return counts;
}

function renderMarkers() {
    if (!markersLayerEl || !matchDurationSec || matchDurationSec <= 0 || viewEnd <= viewStart) {
        if (markersLayerEl) markersLayerEl.innerHTML = '';
        if (markerLabelsLayerEl) markerLabelsLayerEl.innerHTML = '';
        renderTimeAxis();
        return;
    }
    markersLayerEl.innerHTML = '';
    if (markerLabelsLayerEl) {
        markerLabelsLayerEl.innerHTML = '';
    }
    const kindClass = {
        spawn: 'timeline-marker-spawn',
        enemy_spawn: 'timeline-marker-enemy-spawn',
        kill: 'timeline-marker-kill',
        death: 'timeline-marker-death',
    };
    const visible = sortMarkersForPaint(visibleMarkersInView());
    const span = viewSpanSec();

    const withLabel = visible.filter((m) => shortLabelUnderMarker(m));
    const collapseLabels = withLabel.length > 1;

    for (let i = 0; i < visible.length; i++) {
        const m = visible[i];
        const dKind = displayKind(m.kind);
        const t = typeof m.time_sec === 'number' ? m.time_sec : 0;
        const pct = clamp(((t - viewStart) / span) * 100, 0, 100);
        const dot = document.createElement('div');
        dot.className = `timeline-marker ${kindClass[dKind] || 'timeline-marker-other'}`;
        dot.style.left = `${pct}%`;
        const tt = formatMarkerTooltip(m);
        dot.setAttribute('aria-label', tt);
        bindMarkerHover(dot, tt);
        markersLayerEl.appendChild(dot);

        if (markerLabelsLayerEl && !collapseLabels) {
            const slot = document.createElement('div');
            slot.className = 'timeline-marker-label-slot';
            slot.style.left = `${pct}%`;
            bindMarkerHover(slot, tt);
            const text = shortLabelUnderMarker(m);
            if (text) {
                const spanEl = document.createElement('span');
                spanEl.className = 'timeline-marker-label-text';
                spanEl.textContent = text;
                spanEl.setAttribute('aria-label', tt);
                bindMarkerHover(spanEl, tt);
                slot.appendChild(spanEl);
            }
            markerLabelsLayerEl.appendChild(slot);
        }
    }

    if (markerLabelsLayerEl && collapseLabels && withLabel.length > 0) {
        const row = document.createElement('div');
        row.className = 'timeline-marker-label-single';
        const m0 = withLabel[0];
        const tt = formatMarkerTooltip(m0);
        row.textContent = firstDisplayNickname(m0);
        row.setAttribute('aria-label', tt);
        bindMarkerHover(row, tt);
        markerLabelsLayerEl.appendChild(row);
    }
    renderTimeAxis();
}

function applyZoomFromBrush(a, b) {
    let lo = clamp(Math.min(a, b), 0, matchDurationSec);
    let hi = clamp(Math.max(a, b), 0, matchDurationSec);
    if (hi - lo < MIN_VIEW_SPAN_SEC) {
        const mid = (lo + hi) / 2;
        lo = clamp(mid - MIN_VIEW_SPAN_SEC / 2, 0, matchDurationSec);
        hi = clamp(lo + MIN_VIEW_SPAN_SEC, 0, matchDurationSec);
        if (hi - lo < MIN_VIEW_SPAN_SEC) {
            hi = matchDurationSec;
            lo = Math.max(0, hi - MIN_VIEW_SPAN_SEC);
        }
    }
    viewStart = lo;
    viewEnd = hi;
}

/** Зум по двум отметкам времени (сек). Для кисти на дорожке и на графиках ситуации в бою. */
export function commitZoomFromBrush(a, b) {
    if (!timelineReady || matchDurationSec <= 0) {
        return;
    }
    applyZoomFromBrush(a, b);
    hideBrushPreview();
    renderMarkers();
    layoutSyncCursorLine();
    setHint();
    if (typeof rangeChangeCb === 'function') {
        rangeChangeCb();
    }
    if (typeof combatLogRefresh === 'function') {
        combatLogRefresh();
    }
}

function resetTimelineView() {
    if (!matchDurationSec || matchDurationSec <= 0) {
        return;
    }
    viewStart = 0;
    viewEnd = matchDurationSec;
    hideBrushPreview();
    renderMarkers();
    layoutSyncCursorLine();
    setHint();
    gaEvent('lux_timeline_zoom_reset', {});
    if (typeof rangeChangeCb === 'function') {
        rangeChangeCb();
    }
    if (typeof combatLogRefresh === 'function') {
        combatLogRefresh();
    }
}

function onPointerMove(ev) {
    if (!dragMode || !trackEl || dragMode !== 'brush') {
        return;
    }
    const rect = trackEl.getBoundingClientRect();
    const sec = absSecFromClientX(ev, rect);
    layoutBrushPreview(brushAnchorSec, sec);
    setTimelineCursorSec(sec);
}

function endDrag(ev) {
    const was = dragMode;
    if (was === 'brush' && ev && brushSurfaceEl) {
        try {
            brushSurfaceEl.releasePointerCapture(ev.pointerId);
        } catch (_) {
            /* ignore */
        }
    }
    if (was === 'brush' && trackEl) {
        let endSec = brushAnchorSec;
        if (ev && ev.clientX != null) {
            const rect = trackEl.getBoundingClientRect();
            endSec = absSecFromClientX(ev, rect);
        }
        applyZoomFromBrush(brushAnchorSec, endSec);
        hideBrushPreview();
        renderMarkers();
        layoutSyncCursorLine();
        setHint();
    }
    dragMode = null;
    document.removeEventListener('pointermove', onPointerMove);
    document.removeEventListener('pointerup', endDrag);
    document.removeEventListener('pointercancel', endDrag);
    if (was && typeof rangeChangeCb === 'function') {
        rangeChangeCb();
    }
    if (was && typeof combatLogRefresh === 'function') {
        combatLogRefresh();
    }
}

export function setupTimeline(getMatchIndex, onRangeChange) {
    getMatchIndexRef = getMatchIndex || null;
    rangeChangeCb = onRangeChange || null;
    rootEl = document.getElementById('timeline_root');
    hintEl = document.getElementById('timeline_hint');
    if (!rootEl) {
        combatLogRefresh = null;
        return { refresh: () => {}, reset: () => {} };
    }

    rootEl.innerHTML = '';
    const wrap = document.createElement('div');
    wrap.className = 'timeline-inner';

    legendEl = document.createElement('div');
    legendEl.className = 'timeline-legend';
    updateLegend();

    const trackWrap = document.createElement('div');
    trackWrap.className = 'timeline-track-wrap';

    const plotSurfaceEl = document.createElement('div');
    plotSurfaceEl.className = 'timeline-plot-surface';

    trackEl = document.createElement('div');
    trackEl.className = 'timeline-track';
    trackEl.setAttribute('role', 'slider');
    trackEl.setAttribute('aria-label', 'Интервал времени матча');

    brushSurfaceEl = document.createElement('div');
    brushSurfaceEl.className = 'timeline-brush-surface';
    brushSurfaceEl.setAttribute('aria-hidden', 'true');

    brushPreviewEl = document.createElement('div');
    brushPreviewEl.className = 'timeline-brush-preview';
    brushPreviewEl.style.display = 'none';

    markersLayerEl = document.createElement('div');
    markersLayerEl.className = 'timeline-markers-layer';

    trackEl.appendChild(brushSurfaceEl);
    trackEl.appendChild(brushPreviewEl);
    trackEl.appendChild(markersLayerEl);

    markerLabelsLayerEl = document.createElement('div');
    markerLabelsLayerEl.className = 'timeline-marker-labels';

    timeAxisEl = document.createElement('div');
    timeAxisEl.className = 'timeline-time-axis';
    timeAxisEl.setAttribute('aria-hidden', 'true');

    plotSurfaceEl.appendChild(trackEl);
    plotSurfaceEl.appendChild(timeAxisEl);
    plotSurfaceEl.appendChild(markerLabelsLayerEl);

    cursorLineEl = document.createElement('div');
    cursorLineEl.className = 'timeline-sync-cursor';
    cursorLineEl.setAttribute('aria-hidden', 'true');
    plotSurfaceEl.appendChild(cursorLineEl);

    trackWrap.appendChild(plotSurfaceEl);

    trackWrap.addEventListener('pointermove', (e) => {
        if (!trackEl || !timelineReady) {
            return;
        }
        const rect = trackEl.getBoundingClientRect();
        setTimelineCursorSec(absSecFromClientX(e, rect));
    });
    trackWrap.addEventListener('pointerleave', () => {
        setTimelineCursorSec(null);
    });

    const actionsRow = document.createElement('div');
    actionsRow.className = 'timeline-actions';
    resetBtnEl = document.createElement('button');
    resetBtnEl.type = 'button';
    resetBtnEl.className = 'timeline-reset-btn';
    resetBtnEl.textContent = 'Сбросить масштаб';
    resetBtnEl.addEventListener('click', () => resetTimelineView());
    actionsRow.appendChild(resetBtnEl);

    const combatLogDetails = document.createElement('details');
    combatLogDetails.className = 'timeline-combat-loglines';
    const combatLogSummary = document.createElement('summary');
    combatLogSummary.textContent = 'События combat (распарсенные, только выбранный интервал времени)';
    combatLogDetails.appendChild(combatLogSummary);
    const combatLogScroll = document.createElement('div');
    combatLogScroll.className = 'combat-loglines-scroll';
    const combatLogTable = document.createElement('table');
    combatLogTable.className = 'combat-loglines-table';
    const combatThead = document.createElement('thead');
    const hrow = document.createElement('tr');
    ['#', 'Время', 'Тип', 'Содержимое'].forEach((label) => {
        const th = document.createElement('th');
        th.textContent = label;
        hrow.appendChild(th);
    });
    combatThead.appendChild(hrow);
    combatLogTable.appendChild(combatThead);
    const combatLogTbody = document.createElement('tbody');
    combatLogTable.appendChild(combatLogTbody);
    combatLogScroll.appendChild(combatLogTable);
    combatLogDetails.appendChild(combatLogScroll);

    function renderCombatLogLines(matchIdx) {
        combatLogTbody.innerHTML = '';
        const fn = globalThis.getCombatLogLinesJSON;
        if (typeof fn !== 'function') {
            return;
        }
        const tr = typeof getTimeRangeJSON === 'function' ? getTimeRangeJSON() : '{}';
        const raw = fn(matchIdx, tr);
        let rows = [];
        try {
            rows = JSON.parse(raw);
            if (!Array.isArray(rows)) {
                rows = [];
            }
        } catch (_) {
            rows = [];
        }
        rows.forEach((r, i) => {
            const tr = document.createElement('tr');
            const tdN = document.createElement('td');
            tdN.textContent = String(i + 1);
            const tdT = document.createElement('td');
            tdT.textContent = r && r.time != null ? String(r.time) : '';
            const tdK = document.createElement('td');
            tdK.textContent = r && r.kind != null ? String(r.kind) : '';
            const tdS = document.createElement('td');
            tdS.className = 'combat-loglines-summary';
            tdS.textContent = r && r.summary != null ? String(r.summary) : '';
            tr.appendChild(tdN);
            tr.appendChild(tdT);
            tr.appendChild(tdK);
            tr.appendChild(tdS);
            combatLogTbody.appendChild(tr);
        });
    }

    combatLogRefresh = () => {
        const idx = typeof getMatchIndex === 'function' ? getMatchIndex() : 0;
        renderCombatLogLines(idx);
    };

    wrap.appendChild(legendEl);
    wrap.appendChild(trackWrap);
    wrap.appendChild(actionsRow);
    wrap.appendChild(combatLogDetails);
    rootEl.appendChild(wrap);

    brushSurfaceEl.addEventListener('pointerdown', (e) => {
        if (e.button !== 0) {
            return;
        }
        e.preventDefault();
        hideTooltip();
        const rect = trackEl.getBoundingClientRect();
        brushAnchorSec = absSecFromClientX(e, rect);
        dragMode = 'brush';
        layoutBrushPreview(brushAnchorSec, brushAnchorSec);
        try {
            brushSurfaceEl.setPointerCapture(e.pointerId);
        } catch (_) {
            /* ignore */
        }
        document.addEventListener('pointermove', onPointerMove);
        document.addEventListener('pointerup', endDrag);
        document.addEventListener('pointercancel', endDrag);
    });

    window.addEventListener('resize', () => {
        renderTimeAxis();
        layoutSyncCursorLine();
    });
    window.addEventListener('lux-team-swap-changed', () => {
        const idx = typeof getMatchIndex === 'function' ? getMatchIndex() : 0;
        teamsSwapped = readSwapStateForMatch(idx);
        updateLegend();
        renderMarkers();
    });

    function loadMatch() {
        hideTooltip();
        const myGen = ++timelineLoadGen;
        showChartPreloader(plotSurfaceEl, { label: 'Таймлайн…' });
        deferAfterPaint(() => {
            try {
                if (myGen !== timelineLoadGen) {
                    return;
                }
                const fn = globalThis.getTimelineJSON;
                if (typeof fn !== 'function') {
                    timelineReady = false;
                    matchDurationSec = 0;
                    lastMarkers = [];
                    setTimelineCursorSec(null);
                    renderMarkers();
                    setHint();
                    return;
                }
                const idx = typeof getMatchIndex === 'function' ? getMatchIndex() : 0;
                teamsSwapped = readSwapStateForMatch(idx);
                updateLegend();
                const focusPlayer = String(getFocusedPlayer(idx) || '');
                let data = null;
                let timelineWasmMs = 0;
                let timelineJsonMs = 0;
                if (!focusPlayer) {
                    const tw0 = performance.now();
                    const cached = getTimelineEmptyFocusParsed(idx);
                    timelineWasmMs = performance.now() - tw0;
                    if (cached && typeof cached === 'object') {
                        const tj0 = performance.now();
                        data = cached;
                        timelineJsonMs = performance.now() - tj0;
                    }
                }
                if (!data) {
                    const tw0 = performance.now();
                    const raw = fn(idx, focusPlayer);
                    timelineWasmMs = performance.now() - tw0;
                    console.debug('[timeline] raw payload', {
                        match_index: idx,
                        focused_player: focusPlayer,
                        raw_len: typeof raw === 'string' ? raw.length : 0,
                        wasm_ms: Math.round(timelineWasmMs),
                    });
                    if (!raw || raw === 'null') {
                        timelineReady = false;
                        matchDurationSec = 0;
                        lastMarkers = [];
                        setTimelineCursorSec(null);
                        renderMarkers();
                        setHint();
                        return;
                    }
                    const tj0 = performance.now();
                    try {
                        data = JSON.parse(raw);
                    } catch (_) {
                        timelineReady = false;
                        setTimelineCursorSec(null);
                        return;
                    }
                    timelineJsonMs = performance.now() - tj0;
                } else {
                    console.debug('[timeline] raw payload', {
                        match_index: idx,
                        focused_player: focusPlayer,
                        raw_len: 0,
                        wasm_ms: Math.round(timelineWasmMs),
                        cached: true,
                    });
                }
                matchDurationSec = typeof data.end_sec === 'number' ? data.end_sec : 0;
                matchStartUnixMs = typeof data.start_unix_ms === 'number' ? data.start_unix_ms : 0;
                lastMarkers = Array.isArray(data.markers) ? data.markers : [];
                const counts = markerKindCounts(lastMarkers);
                console.debug('[timeline] parsed markers', {
                    match_index: idx,
                    focused_player: focusPlayer,
                    start_unix_ms: matchStartUnixMs,
                    end_sec: matchDurationSec,
                    marker_total: lastMarkers.length,
                    marker_counts: counts,
                    markers_head: lastMarkers.slice(0, 5).map((m) => ({
                        kind: m && m.kind,
                        time_sec: m && m.time_sec,
                        killer: m && m.killer,
                        victim: m && m.victim,
                        player: m && m.player,
                    })),
                    json_ms: Math.round(timelineJsonMs),
                });
                timelineReady = matchDurationSec > 0;
                viewStart = 0;
                viewEnd = matchDurationSec;
                hideBrushPreview();
                const td0 = performance.now();
                renderMarkers();
                layoutSyncCursorLine();
                setHint();
                renderCombatLogLines(idx);
                const timelineRenderMs = performance.now() - td0;
                console.debug('[timeline] rendered', {
                    match_index: idx,
                    view_from_sec: viewStart,
                    view_to_sec: viewEnd,
                    marker_total: lastMarkers.length,
                    render_ms: Math.round(timelineRenderMs),
                });
                gaEvent('lux_timeline_timing', {
                    timeline_wasm_ms: Math.round(timelineWasmMs),
                    timeline_json_ms: Math.round(timelineJsonMs),
                    timeline_data_prep_ms: Math.round(timelineWasmMs + timelineJsonMs),
                    timeline_render_ms: Math.round(timelineRenderMs),
                    marker_count: lastMarkers.length,
                });
            } finally {
                if (myGen === timelineLoadGen) {
                    hideChartPreloader(plotSurfaceEl);
                }
            }
        });
    }

    return { refresh: loadMatch, reset: resetTimelineView };
}

export function getTimelineStartUnixMs() {
    return matchStartUnixMs || 0;
}
