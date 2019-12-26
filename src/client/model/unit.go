package model

type Unit struct {
	unitId            int
	playerId          int
	pathId            int
	baseUnit          BaseUnit
	cell              Cell
	hp                int
	damageLevel       int
	rangeLevel        int
	attack            int
	rng               int `json:"range"`
	isHasted          bool
	isClone           bool
	wasDamageUpgraded bool
	wasRangeUpgraded  bool
}