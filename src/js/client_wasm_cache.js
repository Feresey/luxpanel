/**
 * Клиентский кеш между вызовами WASM в одном «волне» refresh и общий снимок таймлайна без фокуса.
 * Инвалидируется после нового parseFiles.
 */

let timelineEmptyByIdx = new Map(); // matchIndex -> parsed JSON object

let levelsMetaRaw = null;
let levelsMetaValid = false;

export function invalidateClientWasmCaches() {
    timelineEmptyByIdx.clear();
    levelsMetaRaw = null;
    levelsMetaValid = false;
}

/**
 * Один вызов getTimelineJSON(idx, '') на refresh; дальше — из кеша.
 */
export function primeTimelineEmptyFocus(matchIndex) {
    const idx = Number(matchIndex) || 0;
    const fn = globalThis.getTimelineJSON;
    if (typeof fn !== 'function') {
        timelineEmptyByIdx.set(idx, null);
        return;
    }
    const raw = fn(idx, '');
    let data = null;
    try {
        data = JSON.parse(raw || '{}');
    } catch (_) {
        data = null;
    }
    timelineEmptyByIdx.set(idx, data);
}

/**
 * Данные таймлайна с пустым фокусом (маркеры, ally/enemy id).
 */
export function getTimelineEmptyFocusParsed(matchIndex) {
    const idx = Number(matchIndex) || 0;
    if (timelineEmptyByIdx.has(idx)) {
        return timelineEmptyByIdx.get(idx);
    }
    const fn = globalThis.getTimelineJSON;
    if (typeof fn !== 'function') {
        return null;
    }
    try {
        const raw = fn(idx, '');
        const data = JSON.parse(raw || '{}');
        timelineEmptyByIdx.set(idx, data);
        return data;
    } catch (_) {
        return null;
    }
}

function buildByPlayerFromMarkers(markers) {
    const byPlayer = new Map();
    if (!Array.isArray(markers)) {
        return byPlayer;
    }
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
    return byPlayer;
}

/**
 * { allyTeamID, enemyTeamID, byPlayer } как раньше playerTeamsFromTimeline.
 */
export function getPlayerTeamsFromTimelineCached(matchIndex) {
    const data = getTimelineEmptyFocusParsed(matchIndex);
    if (!data || typeof data !== 'object') {
        return { allyTeamID: 0, enemyTeamID: 0, byPlayer: new Map() };
    }
    const markers = Array.isArray(data.markers) ? data.markers : [];
    return {
        allyTeamID: Number(data.ally_team_id) || 0,
        enemyTeamID: Number(data.enemy_team_id) || 0,
        byPlayer: buildByPlayerFromMarkers(markers),
    };
}

/** Один parse getLevelsMetaJSON на волну после parseFiles. */
export function getLevelsMetaRawCached() {
    const fn = globalThis.getLevelsMetaJSON;
    if (typeof fn !== 'function') {
        return '[]';
    }
    if (!levelsMetaValid || levelsMetaRaw == null) {
        levelsMetaRaw = fn();
        levelsMetaValid = true;
    }
    return levelsMetaRaw;
}

export function getLevelsMetaParsedCached() {
    try {
        const raw = getLevelsMetaRawCached();
        const meta = JSON.parse(raw || '[]');
        return Array.isArray(meta) ? meta : [];
    } catch (_) {
        return [];
    }
}
