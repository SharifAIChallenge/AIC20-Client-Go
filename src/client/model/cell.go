package model

type Cell struct {
	row, col int
	units    []Unit
}

func (cell Cell) Equals(cell2 Cell) bool { //TODO is this okay?
	return cell.row == cell2.row && cell.col == cell2.col
}
