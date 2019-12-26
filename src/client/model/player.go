package model

type Player struct {
	playerId      int
	ap            int
	upgradeTokens int
	deck          []int //TODO make it a unit array
	hand          []int
	spells        []Spell
	receivedSpell int
	isAlive       bool //TODO remove? duplicate also in king
	king          King
}
