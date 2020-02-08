package model

type CastUnitSpell struct {
	TypeId        int `json:"typeId"`
	Spell         *Spell
	Id            int   `json:"id"`
	CasterId      int   `json:"casterId"`
	Cell          *Cell `json:"cell"`
	AffectedUnits []int `json:"affectedUnits"`
	UnitId        int   `json:"unitId"`          //TODO Unit obj
	PathId        int   `json:"pathId"`          //TODO Path obj
	CastThisTurn  bool  `json:"wasCastThisTurn"` //TODO remove
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

func (c CastUnitSpell) GetAffectedUnits() []int {
	return c.AffectedUnits
}

func (c CastUnitSpell) WasCastThisTurn() bool {
	return c.WasCastThisTurn()
}
