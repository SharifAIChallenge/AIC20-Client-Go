package model

type Unit struct {
	UnitId         int         `json:"unitId"`
	PlayerId       int         `json:"playerId"`
	Path           *Path       `json:"path"`
	BaseUnit       *BaseUnit   `json:"baseUnit"`
	Cell           *Cell       `json:"cell"` //TODO only pointer?( or empty cell)
	Hp             int         `json:"hp"`
	DamageLevel    int         `json:"damageLevel"`
	RangeLevel     int         `json:"rangeLevel"`
	Attack         int         `json:"attack"`
	Range          int         `json:"range"`
	IsHasted       bool        `json:"isHasted"`
	IsDuplicate    bool        `json:"isDuplicate"`
	Target         *Unit       `json:"targetUnit"`
	TargetIfKing   *King       `json:"targetIfKing"`
	TargetCell     *Cell       `json:"targetCell"`
	AffectedSpells []*CastSpell `json:"affectedCastSpells"` //TODO CastSpellsOnUnit
}


//TODO GetCastSpellsOnUnit?
