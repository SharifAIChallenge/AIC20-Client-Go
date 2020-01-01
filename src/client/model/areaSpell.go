package model

type AreaSpell struct {
	typ      string `json:"type"`
	typeId   int
	priority int
	duration int
	rng      int `json:"range"`
	power    int
	target   string
}

func (areaSpell AreaSpell) GetTypeId() int {
	return areaSpell.typeId
}

func (areaSpell AreaSpell) IsAreaSpell() bool {
	return true
}

func (areaSpell AreaSpell) GetType() string {
	return areaSpell.typ
}

func (areaSpell AreaSpell) GetPriority() int {
	return areaSpell.priority
}
