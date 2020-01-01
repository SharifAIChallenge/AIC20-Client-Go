package model

type UnitSpell struct {
	Type       string `json:"type"`
	TypeId     int    `json:"typeId"`
	Priority   int    `json:"priority"`
	TurnEffect int    `json:"turnEffect"`
}

func (unitSpell UnitSpell) GetTypeId() int {
	return unitSpell.TypeId
}

func (unitSpell UnitSpell) IsAreaSpell() bool {
	return false
}

func (unitSpell UnitSpell) GetType() string {
	return unitSpell.Type
}

func (unitSpell UnitSpell) GetPriority() int {
	return unitSpell.Priority
}
