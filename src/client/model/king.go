package model

type King struct {
	playerId int
	center   Cell
	hp       int
	attack   int
	rng      int  `json:"range"`
	isAlive  bool //TODO make it a method?
}
