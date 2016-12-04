package a_alg

import (
	"math"
	"fmt"
)

type GAlgorithmImpl struct {
	verticalMovementCost int
	diagonalMovementCost int
}

func NewGAlgorithm(movementCost int) GAlgorithmImpl {
	return GAlgorithmImpl{
		diagonalMovementCost:int(math.Sqrt2 * float64(movementCost)),
		verticalMovementCost:movementCost,
	}
}

func (this GAlgorithmImpl) Calculate(from Cell, to Cell) (int) {
	orAll := 0
	xorAll := 0
	for i, pos := range from.GetPosition().GetAsSlice() {
		distance := abs(pos - to.GetPosition().GetAsSlice()[i])
		orAll |= distance
		xorAll ^= distance
	}
	if orAll == 1 {
		if xorAll == 1 {
			return this.verticalMovementCost + from.GetInfo().GValue
		} else {
			return this.diagonalMovementCost + from.GetInfo().GValue
		}
	}
	panic(fmt.Sprintf("Number of ribs between %v and %v is invalid", from, to))
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
