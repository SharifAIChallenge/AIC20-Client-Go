package model

type CastSpell interface {
	GetId() int
	GetSpell() *Spell
	GetCasterId() int
	GetCell() *Cell
	GetAffectedUnits() []*Unit
}
