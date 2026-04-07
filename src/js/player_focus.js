const focusedByMatch = new Map();

function keyForMatch(matchIndex) {
    return String(Number(matchIndex) || 0);
}

export function getFocusedPlayer(matchIndex) {
    return focusedByMatch.get(keyForMatch(matchIndex)) || '';
}

export function setFocusedPlayer(matchIndex, playerName) {
    const k = keyForMatch(matchIndex);
    const v = String(playerName || '').trim();
    if (v) {
        focusedByMatch.set(k, v);
    } else {
        focusedByMatch.delete(k);
    }
    if (typeof window !== 'undefined') {
        window.dispatchEvent(new CustomEvent('lux-player-focus-changed', {
            detail: { matchIndex: Number(matchIndex) || 0, player: v },
        }));
    }
}

export function clearFocusedPlayers() {
    focusedByMatch.clear();
    if (typeof window !== 'undefined') {
        window.dispatchEvent(new CustomEvent('lux-player-focus-changed', {
            detail: { cleared: true },
        }));
    }
}
