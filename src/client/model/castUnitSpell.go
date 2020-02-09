package model

type CastUnitSpell struct {
	Spell         *Spell
	Id            int   `json:"id"`
	CasterId      int   `json:"casterId"`
	Cell          *Cell `json:"cell"`
	AffectedUnits []*Unit
	Unit        *Unit
	Path        *Path
}

func (c CastUnitSpell) GetId() int {
	return c.Id
}

func (c CastUnitSpell) GetSpell() *Spell {
	return c.Spell
}

func (c CastUnitSpell) GetCasterId() int {
	return c.CasterId
}

func (c CastUnitSpell) GetCell() *Cell {
	return c.Cell
}

func (c CastUnitSpell) GetAffectedUnits() []*Unit {
	return c.AffectedUnits
}