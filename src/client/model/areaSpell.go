package model

type AreaSpell struct {
	typeId     int
	turnEffect int
	rng        int `json:"range"`
	power      int
	isDamaging bool
}

func (areaSpell AreaSpell) GetTypeId() int {
	return areaSpell.typeId
}

func (areaSpell AreaSpell) IsAreaSpell() bool {
	return true
}
