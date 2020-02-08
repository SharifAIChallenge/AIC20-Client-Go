package model

type Spell struct {
	Type       string `json:"type"` //TODO ENUM
	TypeId     int    `json:"typeId"`
	Duration   int    `json:"duration"`
	Priority   int    `json:"priority"`
	Range      int    `json:"range"`
	Power      int    `json:"power"`
	IsDamaging bool   `json:"isDamaging"` //TODO json
	Target     string `json:"target"`     //TODO ENUM
}

func (spell Spell) IsAreaSpell() bool {
	return !spell.IsUnitSpell()
}

func (spell Spell) IsUnitSpell() bool {
	return spell.Type == "TELE"
}
