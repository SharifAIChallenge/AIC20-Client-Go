package model

type CastAreaSpell struct {
	TypeId         int   `json:"typeId"`
	Id             int   `json:"id"`
	CasterId       int   `json:"casterId"`
	Cell           *Cell `json:"cell"`
	AffectedUnits  []int `json:"affectedUnits"`
	RemainingTurns int   `json:"remainingTurns"`
	CastThisTurn   bool  `json:"wasCastThisTurn"`
}
