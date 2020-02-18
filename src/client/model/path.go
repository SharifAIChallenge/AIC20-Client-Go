package model

type Path struct {
	Id    int     `json:"id"`
	Cells []*Cell `json:"cells"`
}

func (path Path) Equal(path2 Path) bool {
	return path.Id == path2.Id
}
