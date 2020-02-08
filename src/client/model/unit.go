package model

type Unit struct {
	BaseUnit       *BaseUnit `json:"baseUnit"`
	Cell           *Cell     `json:"cell"` //TODO only pointer?( or empty cell)
	UnitId         int       `json:"unitId"`
	Hp             int       `json:"hp"`
	Path           *Path     `json:"path"`
	Target         *Unit
	TargetCell     *Cell `json:"targetCell"`
	TargetIfKing   *King
	PlayerId       int          `json:"playerId"`
	DamageLevel    int          `json:"damageLevel"`
	RangeLevel     int          `json:"rangeLevel"`
	Range          int          `json:"range"`
	Attack         int          `json:"attack"`
	IsDuplicate    bool         `json:"isDuplicate"`
	IsHasted       bool         `json:"isHasted"`
	AffectedSpells []*CastSpell `json:"affectedCastSpells"` //TODO CastSpellsOnUnit
}
