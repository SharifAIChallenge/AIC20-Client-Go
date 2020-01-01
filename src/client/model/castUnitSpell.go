package model

type CastUnitSpell struct {
	TypeId         int   `json:"typeId"`
	Id             int   `json:"id"`
	CasterId       int   `json:"casterId"`
	UnitId         int   `json:"unitId"`
	PathId         int   `json:"pathId"`
	Target         Cell  `json:"cell"`
	AffectedUnits  []int `json:"affectedUnits"`
	RemainingTurns int   `json:"remainingTurns"`
	CastThisTurn   bool  `json:"wasCastThisTurn"`
}

func (c CastUnitSpell) GetId() int {
	return c.Id
}

func (c CastUnitSpell) GetTypeId() int {
	return c.TypeId
}

func (c CastUnitSpell) GetCasterId() int {
	return c.CasterId
}

func (c CastUnitSpell) GetCell() Cell {
	return c.Target
}

func (c CastUnitSpell) GetRemainingTurns() int {
	return c.RemainingTurns
}

func (c CastUnitSpell) GetAffectedUnits() []int {
	return c.AffectedUnits
}

func (c CastUnitSpell) WasCastThisTurn() bool {
	return c.CastThisTurn
}
