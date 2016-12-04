package a_alg

type hAlgorithmImpl struct {
	finishCell Cell
	costOfStep int
}

func NewHAlgorithm(finish Cell, costOfStep int) HAlgorithm {
	return hAlgorithmImpl{finishCell: finish, costOfStep: costOfStep}
}

func (this hAlgorithmImpl) Calculate(from Cell) (res int) {
	for i, pos := range from.GetPosition().GetAsSlice() {
		res += abs(pos - this.finishCell.GetPosition().GetAsSlice()[i])
	}
	res *= this.costOfStep
	return
}