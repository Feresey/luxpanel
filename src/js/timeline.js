/**
 * Match timeline: view zoom, brush range, markers, WASM time range export.
 */

const MIN_GAP_SEC = 0.05;
const MIN_VIEW_SPAN_SEC = 0.4;

let matchDurationSec = 0;
let viewStart = 0;
let viewEnd = 0;
let rangeFrom = 0;
let rangeTo = 0;
let timelineReady = false;

let rootEl = null;
let trackEl = null;
let brushSurfaceEl = null;
let brushPreviewEl = null;
let markersLayerEl = null;
let selectionEl = null;
let selectionDragEl = null;
let handleLeftEl = null;
let handleRightEl = null;
let hintEl = null;
let legendEl = null;
let markerLabelsLayerEl = null;
let resetBtnEl = null;
let tooltipEl = null;

let dragMode = null; // 'left' | 'right' | 'move' | 'brush' | null
let brushAnchorSec = 0;
let lastMarkers = [];
let rangeChangeCb = null;
/** Обновление таблицы combat-логов под выбранный интервал (ставится в setupTimeline). */
let combatLogRefresh = null;

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
    if (m.kind === 'spawn') {
        const n = (m.player || '').trim();
        const sh = (m.player_ship || '').trim();
        if (n && sh) {
            return `Спавн (союзники): ${n} · ${sh}`;
        }
        if (n) {
            return `Спавн (союзники): ${n}`;
        }
        return 'Спавн (союзники)';
    }
    if (m.kind === 'enemy_spawn') {
        const n = (m.player || '').trim();
        const sh = (m.player_ship || '').trim();
        if (n && sh) {
            return `Спавн (враги): ${n} · ${sh}`;
        }
        if (n) {
            return `Спавн (враги): ${n}`;
        }
        return 'Спавн (враги)';
    }
    if (m.kind === 'kill') {
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
            line1 = `Убийство: ${kPart} → ${vPart}`;
        } else {
            line1 = `Убийство: ${kPart || vPart}`;
        }
        const extra = killDeathTooltipExtra(m);
        return extra ? `${line1}\n${extra}` : line1;
    }
    if (m.kind === 'death') {
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
            line1 = `Смерть союзника: ${kPart} → ${vPart}`;
        } else if (vPart) {
            line1 = `Смерть союзника: ${vPart}`;
        } else {
            line1 = 'Смерть союзника';
        }
        const extra = killDeathTooltipExtra(m);
        return extra ? `${line1}\n${extra}` : line1;
    }
    return m.label || '';
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
    let left = clientX + pad;
    let top = clientY + pad;
    if (left + tw > window.innerWidth - 8) {
        left = clientX - tw - pad;
    }
    if (top + th > window.innerHeight - 8) {
        top = clientY - th - pad;
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
        time_from_sec: rangeFrom,
        time_to_sec: rangeTo,
    });
}

function setHint() {
    if (!hintEl) {
        return;
    }
    if (!timelineReady || matchDurationSec <= 0) {
        hintEl.textContent = '';
        return;
    }
    const sameView = Math.abs(viewStart - 0) < 1e-6 && Math.abs(viewEnd - matchDurationSec) < 1e-6;
    const viewPart = sameView
        ? ''
        : ` · окно ${formatSec(viewStart)}–${formatSec(viewEnd)}`;
    hintEl.textContent = `Фильтр: ${formatSec(rangeFrom)} — ${formatSec(rangeTo)}${viewPart} · матч ${formatSec(matchDurationSec)}`;
}

function layoutSelection() {
    if (!selectionEl || !selectionDragEl || !handleLeftEl || !handleRightEl || !trackEl) {
        return;
    }
    if (!matchDurationSec || matchDurationSec <= 0 || viewEnd <= viewStart) {
        selectionEl.style.display = 'none';
        selectionDragEl.style.display = 'none';
        handleLeftEl.style.display = 'none';
        handleRightEl.style.display = 'none';
        return;
    }
    const span = viewSpanSec();
    const p0 = ((rangeFrom - viewStart) / span) * 100;
    const p1 = ((rangeTo - viewStart) / span) * 100;
    const leftPct = clamp(p0, 0, 100);
    const rightPct = clamp(p1, 0, 100);
    const widthPct = Math.max(rightPct - leftPct, 0);
    selectionEl.style.display = 'block';
    selectionDragEl.style.display = widthPct > 0 ? 'block' : 'none';
    selectionEl.style.left = `${leftPct}%`;
    selectionEl.style.width = `${widthPct}%`;
    selectionDragEl.style.left = `${leftPct}%`;
    selectionDragEl.style.width = `${widthPct}%`;
    // На всю ширину дорожки слой перетаскивания перехватывает клики и мешает рисовать новое окно кистью.
    selectionDragEl.style.pointerEvents = widthPct >= 99.5 ? 'none' : 'auto';
    handleLeftEl.style.display = 'block';
    handleRightEl.style.display = 'block';
    handleLeftEl.style.left = `${leftPct}%`;
    handleRightEl.style.left = `${rightPct}%`;
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

function renderMarkers() {
    if (!markersLayerEl || !matchDurationSec || matchDurationSec <= 0 || viewEnd <= viewStart) {
        if (markersLayerEl) markersLayerEl.innerHTML = '';
        if (markerLabelsLayerEl) markerLabelsLayerEl.innerHTML = '';
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
        const t = typeof m.time_sec === 'number' ? m.time_sec : 0;
        const pct = clamp(((t - viewStart) / span) * 100, 0, 100);
        const dot = document.createElement('div');
        dot.className = `timeline-marker ${kindClass[m.kind] || 'timeline-marker-other'}`;
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
}

function applyRange(from, to) {
    if (!matchDurationSec || matchDurationSec <= 0) {
        rangeFrom = 0;
        rangeTo = 0;
        return;
    }
    let a = clamp(from, 0, matchDurationSec);
    let b = clamp(to, 0, matchDurationSec);
    if (b - a < MIN_GAP_SEC) {
        if (a + MIN_GAP_SEC <= matchDurationSec) {
            b = a + MIN_GAP_SEC;
        } else {
            a = Math.max(0, matchDurationSec - MIN_GAP_SEC);
            b = matchDurationSec;
        }
    }
    if (a > b) {
        [a, b] = [b, a];
    }
    rangeFrom = a;
    rangeTo = b;
}

function applyViewAndRangeFromBrush(a, b) {
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
    rangeFrom = lo;
    rangeTo = hi;
}

function resetTimelineView() {
    if (!matchDurationSec || matchDurationSec <= 0) {
        return;
    }
    viewStart = 0;
    viewEnd = matchDurationSec;
    rangeFrom = 0;
    rangeTo = matchDurationSec;
    hideBrushPreview();
    renderMarkers();
    layoutSelection();
    setHint();
    if (typeof rangeChangeCb === 'function') {
        rangeChangeCb();
    }
    if (typeof combatLogRefresh === 'function') {
        combatLogRefresh();
    }
}

function onPointerMove(ev) {
    if (!dragMode || !trackEl) {
        return;
    }
    const rect = trackEl.getBoundingClientRect();
    const sec = absSecFromClientX(ev, rect);
    if (dragMode === 'left') {
        applyRange(sec, rangeTo);
    } else if (dragMode === 'right') {
        applyRange(rangeFrom, sec);
    } else if (dragMode === 'move') {
        const w = rangeTo - rangeFrom;
        let nf = sec - (w / 2);
        let nt = nf + w;
        nf = clamp(nf, 0, matchDurationSec - w);
        nt = nf + w;
        applyRange(nf, nt);
    } else if (dragMode === 'brush') {
        layoutBrushPreview(brushAnchorSec, sec);
    }
    layoutSelection();
    setHint();
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
        applyViewAndRangeFromBrush(brushAnchorSec, endSec);
        hideBrushPreview();
        renderMarkers();
    }
    dragMode = null;
    document.removeEventListener('pointermove', onPointerMove);
    document.removeEventListener('pointerup', endDrag);
    document.removeEventListener('pointercancel', endDrag);
    layoutSelection();
    setHint();
    if (was && typeof rangeChangeCb === 'function') {
        rangeChangeCb();
    }
    if (was && typeof combatLogRefresh === 'function') {
        combatLogRefresh();
    }
}

export function setupTimeline(getMatchIndex, onRangeChange) {
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
    legendEl.innerHTML = ''
        + '<span class="timeline-legend-item"><i class="timeline-dot spawn"></i> Спавн (союзники)</span>'
        + '<span class="timeline-legend-item"><i class="timeline-dot enemy-spawn"></i> Спавн (враги)</span>'
        + '<span class="timeline-legend-item"><i class="timeline-dot kill"></i> Убийство (враг)</span>'
        + '<span class="timeline-legend-item"><i class="timeline-dot death"></i> Смерть союзника</span>';

    const trackWrap = document.createElement('div');
    trackWrap.className = 'timeline-track-wrap';

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

    selectionEl = document.createElement('div');
    selectionEl.className = 'timeline-selection';
    selectionEl.setAttribute('aria-hidden', 'true');

    selectionDragEl = document.createElement('div');
    selectionDragEl.className = 'timeline-selection-drag';
    selectionDragEl.setAttribute('aria-hidden', 'true');

    markersLayerEl = document.createElement('div');
    markersLayerEl.className = 'timeline-markers-layer';

    handleLeftEl = document.createElement('button');
    handleLeftEl.type = 'button';
    handleLeftEl.className = 'timeline-handle timeline-handle-left';
    handleLeftEl.setAttribute('aria-label', 'Начало интервала');

    handleRightEl = document.createElement('button');
    handleRightEl.type = 'button';
    handleRightEl.className = 'timeline-handle timeline-handle-right';
    handleRightEl.setAttribute('aria-label', 'Конец интервала');

    trackEl.appendChild(brushSurfaceEl);
    trackEl.appendChild(brushPreviewEl);
    trackEl.appendChild(selectionEl);
    trackEl.appendChild(selectionDragEl);
    trackEl.appendChild(markersLayerEl);
    trackEl.appendChild(handleLeftEl);
    trackEl.appendChild(handleRightEl);

    markerLabelsLayerEl = document.createElement('div');
    markerLabelsLayerEl.className = 'timeline-marker-labels';

    trackWrap.appendChild(trackEl);
    trackWrap.appendChild(markerLabelsLayerEl);

    const actionsRow = document.createElement('div');
    actionsRow.className = 'timeline-actions';
    resetBtnEl = document.createElement('button');
    resetBtnEl.type = 'button';
    resetBtnEl.className = 'timeline-reset-btn';
    resetBtnEl.textContent = 'Сбросить масштаб и фильтр';
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

    handleLeftEl.addEventListener('pointerdown', (e) => {
        e.preventDefault();
        e.stopPropagation();
        hideTooltip();
        dragMode = 'left';
        handleLeftEl.setPointerCapture(e.pointerId);
        document.addEventListener('pointermove', onPointerMove);
        document.addEventListener('pointerup', endDrag);
        document.addEventListener('pointercancel', endDrag);
    });

    handleRightEl.addEventListener('pointerdown', (e) => {
        e.preventDefault();
        e.stopPropagation();
        hideTooltip();
        dragMode = 'right';
        handleRightEl.setPointerCapture(e.pointerId);
        document.addEventListener('pointermove', onPointerMove);
        document.addEventListener('pointerup', endDrag);
        document.addEventListener('pointercancel', endDrag);
    });

    selectionDragEl.addEventListener('pointerdown', (e) => {
        e.preventDefault();
        e.stopPropagation();
        hideTooltip();
        dragMode = 'move';
        selectionDragEl.setPointerCapture(e.pointerId);
        document.addEventListener('pointermove', onPointerMove);
        document.addEventListener('pointerup', endDrag);
        document.addEventListener('pointercancel', endDrag);
    });

    window.addEventListener('resize', () => {
        layoutSelection();
    });

    function loadMatch() {
        hideTooltip();
        const fn = globalThis.getTimelineJSON;
        if (typeof fn !== 'function') {
            timelineReady = false;
            matchDurationSec = 0;
            lastMarkers = [];
            renderMarkers();
            setHint();
            return;
        }
        const idx = typeof getMatchIndex === 'function' ? getMatchIndex() : 0;
        const raw = fn(idx);
        if (!raw || raw === 'null') {
            timelineReady = false;
            matchDurationSec = 0;
            lastMarkers = [];
            renderMarkers();
            setHint();
            return;
        }
        let data;
        try {
            data = JSON.parse(raw);
        } catch (_) {
            timelineReady = false;
            return;
        }
        matchDurationSec = typeof data.end_sec === 'number' ? data.end_sec : 0;
        lastMarkers = Array.isArray(data.markers) ? data.markers : [];
        timelineReady = matchDurationSec > 0;
        viewStart = 0;
        viewEnd = matchDurationSec;
        rangeFrom = 0;
        rangeTo = matchDurationSec;
        hideBrushPreview();
        renderMarkers();
        layoutSelection();
        setHint();
        renderCombatLogLines(idx);
    }

    return { refresh: loadMatch, reset: resetTimelineView };
}
