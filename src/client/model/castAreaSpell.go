package model

type CastAreaSpell struct {
	TypeId         int   `json:"typeId"`
	Id             int   `json:"id"`
	CasterId       int   `json:"casterId"`
	Cell           Cell  `json:"cell"`
	AffectedUnits  []int `json:"affectedUnits"`
	RemainingTurns int   `json:"remainingTurns"`
	CastThisTurn   bool  `json:"wasCastThisTurn"`
}

func (c CastAreaSpell) GetId() int {
	return c.Id
}

func (c CastAreaSpell) GetTypeId() int {
	return c.TypeId
}

func (c CastAreaSpell) GetCasterId() int {
	return c.CasterId
}

func (c CastAreaSpell) GetCell() Cell {
	return c.Cell
}

func (c CastAreaSpell) GetRemainingTurns() int {
	return c.RemainingTurns
}

func (c CastAreaSpell) GetAffectedUnits() []int {
	return c.AffectedUnits
}

func (c CastAreaSpell) WasCastThisTurn() bool {
	return c.CastThisTurn
}
