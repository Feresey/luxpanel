package main

import "github.com/Feresey/luxpanel/internal/splitter"

// resolveAllyEnemyTeamIDs returns stable ally/enemy team IDs for UI/chart semantics.
// ally is derived from local client team (when present), enemy is the opposite side from chartTeamSides fallback.
func resolveAllyEnemyTeamIDs(level *splitter.Level) (allyID, enemyID int, ok bool) {
	allyID, allyOK := allyTeamIDForTimeline(level)
	leftID, rightID, lrOK := leftRightTeamIDs(level)
	if !allyOK || !lrOK || allyID == 0 {
		return 0, 0, false
	}
	enemyID = rightID
	if allyID == rightID {
		enemyID = leftID
	}
	if enemyID == 0 || enemyID == allyID {
		return 0, 0, false
	}
	return allyID, enemyID, true
}

