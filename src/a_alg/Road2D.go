package a_alg

type Road2D struct {
	Content []*Cell2D
	grid    Grid2D
}

func NewRoad2D(grid Grid2D) *Road2D {
	curCell := grid.Finish
	path := []*Cell2D{curCell}
	for !curCell.IsStart() {
		curCell = curCell.info.PrevCell.(*Cell2D)
		path = append(path, curCell)
	}
	return &Road2D{Content:path, grid:grid}
}

func (this Road2D)String() string {
	return cellsGrid2DToStringWithConverter(this.grid.Content, func(cell *Cell2D) (string) {
		if !cell.IsStart() && !cell.IsFinish() && this.isInRoad(cell) {
			return cellDirectionsConverter(cell)
		}
		return standardCell2DConverter(cell)
	})
}

func (this Road2D) isInRoad(cell *Cell2D) bool {
	for _, cellFromRoad := range this.Content {
		if cellFromRoad == cell {
			return true
		}
	}
	return false
}


