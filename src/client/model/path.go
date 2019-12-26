package model

type Path struct {
	cells  []Cell //TODO pointers instead?
	pathId int    `json:"id"`
}
