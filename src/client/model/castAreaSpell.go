package model

type CastAreaSpell struct {
	TypeId         int `json:"typeId"`
	Spell          *Spell
	Id             int   `json:"id"`
	CasterId       int   `json:"casterId"`
	Cell           *Cell `json:"cell"`
	AffectedUnits  []int `json:"affectedUnits"`
	RemainingTurns int   `json:"remainingTurns"`
	CastThisTurn   bool  `json:"wasCastThisTurn"` //TODO remove
}

func (c CastAreaSpell) GetId() int {
	return c.Id
}

func (c CastAreaSpell) GetSpell() *Spell {
	return c.Spell
}

func (c CastAreaSpell) GetCasterId() int {
	return c.CasterId
}

func (c CastAreaSpell) GetCell() *Cell {
	return c.Cell
}

func (c CastAreaSpell) GetAffectedUnits() []int {
	return c.AffectedUnits
}

func (c CastAreaSpell) WasCastThisTurn() bool {
	return c.WasCastThisTurn()
}
