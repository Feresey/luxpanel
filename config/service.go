package config

const (
	ShowDamage = "damage"
	ShowHeal   = "heal"
	ShowKill   = "kill"
)

type ServiceConfig struct {
	InputDir  string
	OutputDir string
	TextOut   string
	Show      string
	Damage    string
	Heal      string
	Kill      string
}
