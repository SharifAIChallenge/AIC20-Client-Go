package model

type UnitSpell struct {
	typeId     int
	turnEffect int
}

func (unitSpell UnitSpell) GetTypeId() int {
	return unitSpell.typeId
}

func (unitSpell UnitSpell) IsAreaSpell() bool {
	return false
}
