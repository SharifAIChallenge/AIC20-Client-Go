package model

type UnitSpell struct {
	spell  Spell
	typeId int
}

func (unitSpell UnitSpell) GetTypeId() int {
	return unitSpell.typeId
}

func (unitSpell UnitSpell) IsAreaSpell() bool {
	return false
}



