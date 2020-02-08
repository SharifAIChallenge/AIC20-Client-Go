package model

type Path struct {
	Id    int    `json:"id"`
	Cells []Cell `json:"cells"`
}
