package model

type AreaSpell struct {
	spell      Spell
	rng        int
	power      int
	turnEffect int
	isDamaging bool
	typeId     int
}

func (areaSpell AreaSpell) GetTypeId() int {
	return areaSpell.typeId
}

func (areaSpell AreaSpell) IsAreaSpell() bool {
	return true
}


