package model

type CastUnitSpell struct {
	typeId          int
	id              int
	casterId        int
	unitId          int
	pathId          int
	target          Cell `json:"cell"`
	affectedUnits   []int
	remainingTurns  int //TODO remvoe this for unit spells?
	wasCastThisTurn bool
}

func (c CastUnitSpell) GetId() int {
	return c.id
}

func (c CastUnitSpell) GetTypeId() int {
	return c.typeId
}

func (c CastUnitSpell) GetCasterId() int {
	return c.casterId
}

func (c CastUnitSpell) GetCell() Cell {
	return c.GetCell()
}

func (c CastUnitSpell) GetRemainingTurns() int {
	return c.remainingTurns
}

func (c CastUnitSpell) GetAffectedUnits() []int {
	return c.affectedUnits
}

func (c CastUnitSpell) WasCastThisTurn() bool {
	return c.wasCastThisTurn
}
