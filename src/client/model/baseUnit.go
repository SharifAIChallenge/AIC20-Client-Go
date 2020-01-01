package model

type BaseUnit struct {
	TypeId     int    `json:"typeId"`
	MaxHp      int    `json:"maxHp"`
	BaseAttack int    `json:"baseAttack"`
	BaseRange  int    `json:"baseRange"`
	Ap         int    `json:"ap"`
	Target     string `json:"target"`
	IsFlying   bool   `json:"isFlying"`
	IsMultiple bool   `json:"isMultiple"`
}
