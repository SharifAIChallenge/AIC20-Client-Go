package model

type AreaSpell struct {
	typeId     int
	duration   int
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
