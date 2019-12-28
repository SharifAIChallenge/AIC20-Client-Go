package model

type Map struct {
	width       int `json:"rows"`
	height      int `json:"cols"`
	paths       []Path
	units       []Unit
	kings       []King
	unitsInCell [][][]Unit
}
