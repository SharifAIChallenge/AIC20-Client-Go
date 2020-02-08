package model

type King struct {
	Center     *Cell `json:"center"`
	Hp         int   `json:"hp"`
	Attack     int   `json:"attack"`
	Range      int   `json:"range"`
	IsAlive    bool  `json:"isAlive"`
	PlayerId   int   `json:"playerId"`
	Target     *Unit `json:"targetUnit"` //TODO wtf is targetCell
	TargetCell *Cell `json:"targetCell"`
}
