package model

type Spell struct {
	Type       string `json:"type"`
	TypeId     int    `json:"typeId"`
	Duration   int    `json:"duration"`
	Priority   int    `json:"priority"`
	Range      int    `json:"range"`
	Power      int    `json:"power"`
	IsDamaging bool   `json:"isDamaging"`
	Target     string `json:"target"`
}

func (spell Spell) IsAreaSpell() bool {
	return !spell.IsUnitSpell()
}

func (spell Spell) IsUnitSpell() bool {
	return spell.Type == "TELE"
}
