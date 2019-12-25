package model

type BaseUnit struct {
	typ, maxHp, attack, rng, target int //TODO handle enums?
	isFlying, isMultiple            bool
}
