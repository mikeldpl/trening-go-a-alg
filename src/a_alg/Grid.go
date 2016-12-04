package a_alg

type Grid interface {
	GetStart() Cell
	GetFinish() Cell
	getNeighbors(Cell) []Cell
	ToInfoString() string
}
