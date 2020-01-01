package model

type Path struct {
	Cells  []Cell `json:"cells"`
	PathId int    `json:"id"`
}
