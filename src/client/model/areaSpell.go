package model

type AreaSpell struct {
	Type     string `json:"type"`
	TypeId   int    `json:"typeId"`
	Priority int    `json:"priority"`
	Duration int    `json:"duration"`
	Range    int    `json:"range"`
	Power    int    `json:"Power"`
	Target   string `json:"target"`
}

func (areaSpell AreaSpell) GetTypeId() int {
	return areaSpell.TypeId
}

func (areaSpell AreaSpell) IsAreaSpell() bool {
	return true
}

func (areaSpell AreaSpell) GetType() string {
	return areaSpell.Type
}

func (areaSpell AreaSpell) GetPriority() int {
	return areaSpell.Priority
}
