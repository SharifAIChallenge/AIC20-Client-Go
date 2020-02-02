package model

type BaseUnit struct {
	TypeId     int    `json:"typeId"`
	MaxHp      int    `json:"maxHp"`
	BaseAttack int    `json:"baseAttack"`
	BaseRange  int    `json:"baseRange"`
	Ap         int    `json:"ap"`
	TargetType string `json:"target"` //TODO json name?
	IsFlying   bool   `json:"isFlying"`
	IsMultiple bool   `json:"isMultiple"`
}
