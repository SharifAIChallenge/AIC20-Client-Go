package model

type Unit struct {
	baseUnit                                                   BaseUnit
	cell                                                       Cell
	unitId, hp, playerId, damageLevel, rangeLevel, attack, rng int
	path                                                       Path //TODO pathid or pointer to path?
	isHasted, wasDamageUpgraded, wasRangeUpgraded              bool
}
