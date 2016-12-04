package a_alg

import "fmt"

type cells2D [][]*Cell2D

type Grid2D struct {
	Start   *Cell2D
	Finish  *Cell2D
	Content cells2D

	YLen    int
	XLen    int
}

func (this Grid2D) GetStart() Cell {
	return this.Start
}

func (this Grid2D) GetFinish() Cell {
	return this.Finish
}

func (this Grid2D) String() string {
	return fmt.Sprintf("YLen = %d, XLen = %d \n%s", this.YLen, this.XLen, cellsGrid2DToString(this.Content))
}

func (this Grid2D) ToInfoString() string {
	return "\n" + cellsGrid2DToStringWithDirections(this.Content)
}

func NewGrid2D(stringRepresentation string) Grid2D {
	cells := stringToCellsGrid2D(stringRepresentation)
	return NewGrid2DFromCells(cells)
}

func NewGrid2DFromCells(cells cells2D) Grid2D {
	instance := Grid2D{Content:cells}
	instance.YLen = findAndValidateYLen(cells)
	instance.XLen = findAndValidateXLen(cells)
	instance.Start, instance.Finish = findStartAndFinish(cells)
	return instance
}

func findStartAndFinish(cells cells2D) (start, finish *Cell2D) {
	startFound, finishFound := false, false
	for _, line := range cells {
		for _, cell := range line {
			switch cell.Type {
			case StartCell:
				if startFound {
					panic("More then one start cell")
				}
				startFound = true
				start = cell
			case FinishCell:
				if finishFound {
					panic("More then one finish cell")
				}
				finishFound = true
				finish = cell
			}
		}
	}
	if !finishFound || !startFound {
		panic("Start or finish isn't present")
	}
	return
}

func findAndValidateYLen(cells cells2D) int {
	length := len(cells)
	if length < 1 {
		panic("Grid y len < 1")
	}
	return length
}

func findAndValidateXLen(cells cells2D) int {
	length := len(cells[0])
	for _, element := range cells {
		if len(element) != length {
			panic("Grid isn't rectangle")
		}
	}
	return length
}

func (this Grid2D) getNeighbors(cell Cell) (res []Cell) {
	res = make([]Cell, 0, 8)
	cell2D := cell.(*Cell2D)
	res = this.addToNeighbors(res, cell2D.GetX() - 1, cell2D.GetY() + 1)
	res = this.addToNeighbors(res, cell2D.GetX(), cell2D.GetY() + 1)
	res = this.addToNeighbors(res, cell2D.GetX() + 1, cell2D.GetY() + 1)
	res = this.addToNeighbors(res, cell2D.GetX() + 1, cell2D.GetY())

	res = this.addToNeighbors(res, cell2D.GetX() + 1, cell2D.GetY() - 1)
	res = this.addToNeighbors(res, cell2D.GetX(), cell2D.GetY() - 1)
	res = this.addToNeighbors(res, cell2D.GetX() - 1, cell2D.GetY() - 1)
	res = this.addToNeighbors(res, cell2D.GetX() - 1, cell2D.GetY())
	return
}

func (this Grid2D) addToNeighbors(res []Cell, x, y int) []Cell {
	if cell := this.getCell(x, y); cell != nil {
		return append(res, cell)
	}
	return res
}

func (this Grid2D) getCell(x, y int) *Cell2D {
	if x < 0 || x >= this.XLen || y < 0 || y >= this.YLen {
		return nil
	}
	return this.Content[y][x]
}
