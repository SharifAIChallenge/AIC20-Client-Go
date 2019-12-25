package model

type Player struct {
	playerId      int
	ap            int
	upgradeTokens int
	deck          []Unit
	hand          []Unit
	spells        []Spell
	acquiredSpell int
	isAlive       bool //TODO remove? duplicate also in king
	king          King
}
