package model

type BaseUnit struct {
	typeId     int
	maxHp      int
	baseAttack int
	baseRange  int
	target     string
	isFlying   bool
	isMultiple bool
}
