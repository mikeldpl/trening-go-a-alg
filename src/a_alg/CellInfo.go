package a_alg

type CellInfo struct {
	GValue   int
	HValue   int
	PrevCell Cell
	Closed   bool
}

func (this CellInfo) GetFValue() int {
	return this.HValue + this.GValue;
}

func (this CellInfo) LessThen(info CellInfo) bool {
	return this.GetFValue() < info.GetFValue()
}

func (this *CellInfo) Close() {
	this.Closed = true
}

func (this CellInfo) isOpened() bool {
	return this.PrevCell != nil && !this.Closed
}
