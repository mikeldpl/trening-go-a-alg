package a_alg

type CellType int

const (
	StartCell CellType = iota
	FinishCell
	BarrierCell
	EmptyCell
)

type BaseCell struct {
	Type CellType
	info CellInfo
}

func (this *BaseCell)IsStart() bool {
	return this.Type == StartCell
}

func (this *BaseCell)IsFinish() bool {
	return this.Type == FinishCell
}

func (this *BaseCell)IsBarrier() bool {
	return this.Type == BarrierCell
}

func (this *BaseCell)IsEmpty() bool {
	return this.Type == EmptyCell
}

func (this *BaseCell) LessThen(cell interface{}) bool {
	return this.info.LessThen(cell.(*BaseCell).info)
}

func (this *BaseCell) GetInfo() CellInfo {
	return this.info
}

func (this *BaseCell) SetInfo(info CellInfo) {
	this.info = info
}

func (this *BaseCell) Close() {
	this.info.Closed = true
}