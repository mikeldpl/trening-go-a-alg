package a_alg

import (
	"errors"
	"log"
)

type Searcher struct {
	grid     Grid
	HAlg     HAlgorithm
	GAlg     GAlgorithm
	openList LinkedPriorityQueue
}

func SearchStandardAStar2D(picture string) (error, *Road2D) {
	grid := NewGrid2D(picture)
	searcher := Searcher{
		grid: grid,
		GAlg: NewGAlgorithm(10),
		HAlg: NewHAlgorithm(grid.Finish, 6),
	}
	if searcher.FindRoad() {
		return nil, NewRoad2D(grid);
	}
	return errors.New("Can't find road from start to finish"), nil
}

func (this *Searcher) FindRoad() bool {
	if this.grid.GetStart().GetPosition() == this.grid.GetFinish().GetPosition() {
		return false
	}
	ok := true
	for nextCell := this.grid.GetStart(); ok && nextCell != nil; nextCell, ok = this.openList.Pop().(Cell) {
		nextCell.Close()
		for _, neighbor := range this.grid.getNeighbors(nextCell) {
			if neighbor.GetInfo().Closed {
				continue
			}
			if neighbor.IsEmpty() {
				info := this.calculateInfo(neighbor, nextCell)
				this.ProcessEmptyCell(neighbor, info)
			} else if neighbor.IsFinish() {
				neighbor.SetInfo(CellInfo{PrevCell: nextCell})
				return true
			}
		}
		log.Printf("%v\n", nextCell)
		log.Println(this.grid.ToInfoString())
	}
	return false
}

func (this*Searcher) ProcessEmptyCell(neighbor Cell, info CellInfo) {
	if neighbor.GetInfo().isOpened() {
		if info.LessThen(neighbor.GetInfo()) {
			neighbor.SetInfo(info)
		}
	} else {
		neighbor.SetInfo(info)
		this.openList.Add(neighbor)
	}
}

func (this *Searcher) calculateInfo(calculated, prevCell Cell) CellInfo {
	return CellInfo{
		PrevCell: prevCell,
		GValue: this.GAlg.Calculate(prevCell, calculated),
		HValue: this.HAlg.Calculate(calculated),
	}
}
