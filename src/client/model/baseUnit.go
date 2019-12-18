package model

type BaseUnit struct {
	typ, maxHp, attack, level, rng, target int //TODO handle enums?
	isFlying, isMultiple                   bool
}
