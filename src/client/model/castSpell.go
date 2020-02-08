package model

type CastSpell interface { //TODO which methods to keep?
	GetId() int
	GetSpell() *Spell
	GetCasterId() int
	GetCell() *Cell
	GetAffectedUnits() []int
	WasCastThisTurn() bool
}
