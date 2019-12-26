package model

type CastAreaSpell struct {
	typeId        int
	casterId      int
	cell          Cell
	affectedUnits []int //TODO unit instead of unitId
}
