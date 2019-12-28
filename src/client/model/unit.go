package model

type Unit struct {
	unitId            int
	playerId          int
	pathId            int
	baseUnit          BaseUnit
	cell              Cell //TODO only pointer?( or empty cell)
	hp                int
	damageLevel       int
	rangeLevel        int
	attack            int
	rng               int `json:"range"`
	activePoisons     int
	isHasted          bool
	isClone           bool
	wasDamageUpgraded bool
	wasRangeUpgraded  bool
	wasPlayedThisTurn bool
}
