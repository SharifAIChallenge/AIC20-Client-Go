package model

type King struct {
	PlayerId int  `json:"playerId"`
	Center   Cell `json:"center"`
	Hp       int  `json:"hp"`
	Attack   int  `json:"attack"`
	Range    int  `json:"range"`
	IsAlive  bool `json:"isAlive"`
	Target   int  `json:"Target"`
}
