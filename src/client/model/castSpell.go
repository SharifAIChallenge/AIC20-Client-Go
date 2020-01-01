package model

type CastSpell interface { //TODO which methods to keep?
	GetId() int
	GetTypeId() int
	GetCasterId() int
	GetCell() Cell
	GetRemainingTurns() int
	GetAffectedUnits() []int
	WasCastThisTurn() bool
}
