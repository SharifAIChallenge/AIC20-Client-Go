package model

type Map struct {
	Width       int        `json:"rows"`
	Height      int        `json:"cols"`
	Paths       []Path     `json:"paths"`
	Units       []Unit     `json:"units"`
	Kings       []King     `json:"kings"`
	UnitsInCell [][][]Unit `json:"unitsInCell"`
}
