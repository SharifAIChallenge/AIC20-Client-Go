package model

type CastUnitSpell struct {
	TypeId         int   `json:"typeId"`
	Id             int   `json:"id"`
	CasterId       int   `json:"casterId"`
	UnitId         int   `json:"unitId"` //TODO Unit obj
	PathId         int   `json:"pathId"` //TODO Path obj
	Target         *Cell `json:"cell"`
	AffectedUnits  []int `json:"affectedUnits"`
	RemainingTurns int   `json:"remainingTurns"`
	CastThisTurn   bool  `json:"wasCastThisTurn"` //TODO remove
}
