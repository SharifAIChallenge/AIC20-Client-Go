package model

type UnitSpell struct {
	typ        string `json:"type"`
	typeId     int
	priority   int
	turnEffect int
}

func (unitSpell UnitSpell) GetTypeId() int {
	return unitSpell.typeId
}

func (unitSpell UnitSpell) IsAreaSpell() bool {
	return false
}

func (unitSpell UnitSpell) GetType() string {
	return unitSpell.typ
}

func (unitSpell UnitSpell) GetPriority() int {
	return unitSpell.priority
}
