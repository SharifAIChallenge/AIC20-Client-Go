package model

type CastUnitSpell struct {
	TypeId         int   `json:"typeId"`
	Id             int   `json:"id"`
	CasterId       int   `json:"casterId"`
	UnitId         int   `json:"unitId"`
	PathId         int   `json:"pathId"`
	Target         *Cell `json:"cell"`
	AffectedUnits  []int `json:"affectedUnits"`
	RemainingTurns int   `json:"remainingTurns"`
	CastThisTurn   bool  `json:"wasCastThisTurn"`
}
