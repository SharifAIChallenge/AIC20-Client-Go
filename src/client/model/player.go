package model

type Player struct {
	PlayerId      int     `json:"playerId"`
	Ap            int     `json:"ap"`
	UpgradeTokens int     `json:"upgradeTokens"`
	Deck          []int   `json:"deck"`
	Hand          []int   `json:"hand"`
	Spells        []Spell `json:"spells"`
	ReceivedSpell int     `json:"receivedSpell"`
	IsAlive       bool    `json:"isAlive"` //TODO remove? duplicate also in king
	King          King    `json:"king"`
}
