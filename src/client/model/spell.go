package model

type Spell struct {
	Type       string `json:"type"`
	TypeId     int    `json:"typeId"`
	Priority   int    `json:"priority"`
	Duration   int    `json:"duration"`
	Range      int    `json:"range"`
	Power      int    `json:"Power"`
	IsDamaging bool   //TODO json
	Target     string `json:"target"`
}
