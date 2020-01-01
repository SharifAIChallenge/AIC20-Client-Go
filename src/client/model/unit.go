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
	isHasted          bool
	isDuplicate       bool
	wasDamageUpgraded bool
	wasRangeUpgraded  bool
	wasPlayedThisTurn bool
	target            int
	affectedSpells    []int
}
