package model

type Spell interface {
	GetType() string
	GetTypeId() int
	GetPriority() int
	IsAreaSpell() bool
}
