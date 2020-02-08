package model

type BaseUnit struct {
	TypeId     int    `json:"typeId"`
	MaxHp      int    `json:"maxHp"`
	BaseAttack int    `json:"baseAttack"`
	BaseRange  int    `json:"baseRange"`
	TargetType string `json:"target"` //TODO json name? //TODO Enum
	IsFlying   bool   `json:"isFlying"`
	IsMultiple bool   `json:"isMultiple"`
	Ap         int    `json:"ap"`
}
