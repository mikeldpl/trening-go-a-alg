package a_alg

type Coordinates2D struct {
	x int
	y int
}

func (this Coordinates2D) GetAsSlice() []int {
	return []int{this.y, this.x}
}

type Cell2D struct {
	BaseCell
	Position Coordinates2D
}

func (this *Cell2D) GetX() int {
	return this.Position.x
}

func (this *Cell2D) GetY() int {
	return this.Position.y
}

func (this *Cell2D) LessThen(cell interface{}) bool {
	return this.BaseCell.LessThen(&cell.(*Cell2D).BaseCell)
}
func (this *Cell2D) GetPosition() Coordinates {
	return this.Position
}