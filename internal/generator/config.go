package main

type Config struct {
	OutputDir     string
	OutputPackage string
}

var typeByName = map[string]string{
	"LogTime":          "time.Time",
	"AuraID":           "int",
	"InitiatorID":      "int",
	"RecipientID":      "int",
	"WinnerTeamID":     "int",
	"SessionID":        "int",
	"PlayerID":         "int",
	"ConnectionStatus": "int",
	"SessionPlayerID":  "int",
	"TeamID":           "int",
	"GroupID":          "*int",
	"ClientTeamID":     "*int",
	"DamageModifiers":  "DamageModifiersMap",
	"CloseReason":      "ConnectionClosedReason",
	"FriendlyFire":     "bool",
	"DamageTotal":      "float32",
	"DamageHull":       "float32",
	"DamageShield":     "float32",
	"Heal":             "float32",
	"GameTime":         "time.Duration",
	"Reward":           "int",
}

var parserByType = map[string]string{
	"time.Time":              "parseTime(now)",
	"int":                    "strconv.Atoi",
	"*int":                   "parseOptional(strconv.Atoi)",
	"float32":                "parseFloat",
	"bool":                   "parseBool",
	"string":                 "",
	"DamageModifiersMap":     "parseDamageModifiers",
	"ConnectionClosedReason": "parseConnectionClosedReason",
	"time.Duration":          "parseSeconds",
}
