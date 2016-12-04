package a_alg

type Coordinates interface {
	GetAsSlice() []int
}

type Cell interface {
	LessThen(cell interface{}) bool
	IsStart() bool
	IsFinish() bool
	IsBarrier() bool
	IsEmpty() bool
	GetPosition() Coordinates
	GetInfo() CellInfo
	SetInfo(CellInfo)
	Close()
}
