package model

type Map struct {
	RowNum      int         `json:"rows"`
	ColNum      int         `json:"cols"`
	Paths       []*Path     `json:"paths"`
	Units       []*Unit     `json:"units"`
	Kings       []*King     `json:"kings"`
	UnitsInCell [][][]*Unit `json:"unitsInCell"`
	Cells       [][]*Cell   `json:"cells"`
}

func (mp Map) getCell(row, col int) *Cell { //TODO Cell or *Cell?
	cell := mp.Cells[row][col]
	return cell
}
