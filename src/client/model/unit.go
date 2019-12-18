package model

type Unit struct {
	baseUnit             BaseUnit
	cell                 Cell
	unitId, hp, playerId int
	path                 Path //TODO pathid or pointer to path?
	isHasted             bool
}
