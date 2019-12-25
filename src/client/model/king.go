package model

type King struct {
	playerId int
	center   Cell
	hp       int
	attack   int
	rng      int  `json:"range"`
	isAlive  bool //TODO make a method?
	//TODO isYou,isYourFriend?
}
