package a_alg

import (
	"strings"
)

const (
	startMark = "O"
	finishMark = "X"
	barrierMark = "#"
	emptyMark = "."

	cellSeparator = " "
	cellSeparatorForOutput = "\t"
	lineSeparator = "\n"
)

var (
	markToCellTypeMap = map[string]CellType{
		startMark: StartCell,
		"0": StartCell,
		"o": StartCell,
		finishMark: FinishCell,
		"x": FinishCell,
		barrierMark: BarrierCell,
		emptyMark: EmptyCell,
	}

	cellTypeToMarkMap = map[CellType]string{
		StartCell: startMark,
		FinishCell: finishMark,
		BarrierCell: barrierMark,
		EmptyCell: emptyMark,
	}

	directionToMark = [3][3]string{
		{"↖", "↑", "↗"},
		{"←", "*", "→"},
		{"↙", "↓", "↘"},
	}
)

func stringToCellsGrid2D(stringRepresentation string) cells2D {
	cells := [][]*Cell2D{}
	for i, line := range strings.Split(stringRepresentation, lineSeparator) {
		if line != "" {
			cellLine := []*Cell2D{}
			for j, cellChar := range strings.Split(line, cellSeparator) {
				cellType, ok := markToCellTypeMap[cellChar]
				if !ok {
					continue
				}
				cell := &Cell2D{
					BaseCell: BaseCell{
						Type:cellType,
					},
					Position:Coordinates2D{x:j, y:i},
				}
				cellLine = append(cellLine, cell)
			}
			cells = append(cells, cellLine)
		}
	}
	return cells
}

func cellsGrid2DToStringWithDirections(cells cells2D) string {
	return cellsGrid2DToStringWithConverter(cells, cell2DWithDirectionsConverter)
}

func cell2DWithDirectionsConverter(cell *Cell2D) string {
	if cell.info.PrevCell != nil {
		return cellDirectionsConverter(cell)
	}
	return cellTypeToMarkMap[cell.Type]
}

func cellDirectionsConverter(from *Cell2D) string {
	to := *from.info.PrevCell.(*Cell2D)
	xFrom, yFrom := from.GetX(), from.GetY();
	xTo, yTo := to.GetX(), to.GetY();

	return directionToMark[yTo - yFrom + 1][xTo - xFrom + 1]
}

func cellsGrid2DToString(cells cells2D) string {
	return cellsGrid2DToStringWithConverter(cells, standardCell2DConverter)
}

func standardCell2DConverter(cell *Cell2D) string {
	return cellTypeToMarkMap[cell.Type]
}

func cellsGrid2DToStringWithConverter(cells cells2D, converter func(*Cell2D) (string)) (result string) {
	for _, line := range cells {
		for _, cell := range line {
			result += converter(cell) + cellSeparatorForOutput
		}
		result += lineSeparator
	}
	return
}