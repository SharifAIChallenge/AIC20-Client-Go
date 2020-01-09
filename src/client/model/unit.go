package model

type Unit struct {
	UnitId            int       `json:"unitId"`
	PlayerId          int       `json:"playerId"`
	PathId            int       `json:"pathId"`
	BaseUnit          *BaseUnit `json:"baseUnit"`
	Cell              *Cell     `json:"cell"` //TODO only pointer?( or empty cell)
	Hp                int       `json:"hp"`
	DamageLevel       int       `json:"damageLevel"`
	RangeLevel        int       `json:"rangeLevel"`
	Attack            int       `json:"attack"`
	Range             int       `json:"range"`
	IsHasted          bool      `json:"isHasted"`
	IsDuplicate       bool      `json:"isDuplicate"`
	WasDamageUpgraded bool      `json:"wasDamageUpgraded"`
	WasRangeUpgraded  bool      `json:"wasRangeUpgraded"`
	WasPlayedThisTurn bool      `json:"wasPlayedThisTurn"`
	Target            int       `json:"target"`
	TargetCell        *Cell     `json:"targetCell"`
	AffectedSpells    []int     `json:"affectedSpells"`
}
