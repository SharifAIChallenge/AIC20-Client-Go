package model

type Spell interface {
	GetTypeId() int
	IsAreaSpell() bool
}
