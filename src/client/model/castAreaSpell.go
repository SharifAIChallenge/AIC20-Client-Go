package model

type CastAreaSpell struct {
	typeId          int
	id              int
	casterId        int
	cell            Cell
	affectedUnits   []int //TODO unit instead of unitId
	remainingTurns  int
	wasCastThisTurn bool
}

func (c CastAreaSpell) GetId() int {
	return c.id
}

func (c CastAreaSpell) GetTypeId() int {
	return c.typeId
}

func (c CastAreaSpell) GetCasterId() int {
	return c.casterId
}

func (c CastAreaSpell) GetCell() Cell {
	return c.GetCell()
}

func (c CastAreaSpell) GetRemainingTurns() int {
	return c.remainingTurns
}

func (c CastAreaSpell) GetAffectedUnits() []int {
	return c.affectedUnits
}

func (c CastAreaSpell) WasCastThisTurn() bool {
	return c.wasCastThisTurn
}
