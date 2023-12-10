package parser

type GameLogLineType int

const (
	GameLogLineTypeConnect GameLogLineType = iota
	GameLogLineTypeAddPlayer
	GameLogLineTypeLevelStarting
	GameLogLineTypeLevelStarted
	GameLogLineTypePlayerLeave
	GameLogLineTypeAVGNetStat
	GameLogLineTypeGameFinished
)
