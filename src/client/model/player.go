package model

type Player struct {
	playerId, ap, upgradeTokens int
	deck                        []Unit
	hand                        []Unit
	spells                      []Spell
	isAlive                     bool
	king                        King
}
