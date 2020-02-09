package model

type CastAreaSpell struct {
	Spell          *Spell
	Id             int   `json:"id"`
	CasterId       int   `json:"casterId"`
	Cell           *Cell `json:"cell"`
	AffectedUnits  []*Unit
	RemainingTurns int   `json:"remainingTurns"`
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

func (c CastAreaSpell) GetAffectedUnits() []*Unit {
	return c.AffectedUnits
}
