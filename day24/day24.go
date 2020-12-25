package day24

import (
	"strings"
)

type HexPoint struct {
	X, Y int
}

type HexGrid map[HexPoint]bool

func (hp HexPoint) W() HexPoint {
	return HexPoint{hp.X - 1, hp.Y}
}

func (hp HexPoint) E() HexPoint {
	return HexPoint{hp.X + 1, hp.Y}
}

func (hp HexPoint) NW() HexPoint {
	return HexPoint{hp.X, hp.Y - 1}
}

func (hp HexPoint) SE() HexPoint {
	return HexPoint{hp.X, hp.Y + 1}
}

func (hp HexPoint) NE() HexPoint {
	return HexPoint{hp.X + 1, hp.Y - 1}
}

func (hp HexPoint) SW() HexPoint {
	return HexPoint{hp.X - 1, hp.Y + 1}
}

func (hp HexPoint) Neighbors() []HexPoint {
	return []HexPoint{hp.W(), hp.NW(), hp.NE(), hp.E(), hp.SE(), hp.SW()}
}

func FollowPath(path string) HexPoint {
	pos := HexPoint{0, 0}

	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'w':
			pos = pos.W()
		case 'e':
			pos = pos.E()
		case 'n':
			i++
			switch path[i] {
			case 'w':
				pos = pos.NW()
			case 'e':
				pos = pos.NE()
			}
		case 's':
			i++
			switch path[i] {
			case 'w':
				pos = pos.SW()
			case 'e':
				pos = pos.SE()
			}
		}
	}

	return pos
}

func (hg HexGrid) FlipPath(path string) {
	target := FollowPath(path)
	hg[target] = !hg[target]
}

func PrepareTiles(paths []string) HexGrid {
	grid := make(HexGrid)
	for _, line := range paths {
		grid.FlipPath(line)
	}
	return grid
}

func (hg HexGrid) CountBlackTiles() int {
	count := 0
	for _, black := range hg {
		if black {
			count++
		}
	}
	return count
}

func (hg HexGrid) Step() HexGrid {
	countBlackNeighbors := make(map[HexPoint]int)
	for point, black := range hg {
		if black {
			for _, neighbor := range point.Neighbors() {
				countBlackNeighbors[neighbor]++
			}
		}
	}
	nextGrid := make(HexGrid)
	for point, blackNeighbors := range countBlackNeighbors {
		currentlyBlack := hg[point]
		if (currentlyBlack && (blackNeighbors == 1 || blackNeighbors == 2)) || !currentlyBlack && blackNeighbors == 2 {
			nextGrid[point] = true
		}
	}
	return nextGrid
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	return PrepareTiles(lines).CountBlackTiles()
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	grid := PrepareTiles(lines)
	for i := 0; i < 100; i++ {
		grid = grid.Step()
	}
	return grid.CountBlackTiles()
}

func ListBlackTiles(hg HexGrid) []HexPoint {
	list := make([]HexPoint, len(hg))
	count := 0
	for point, black := range hg {
		if black {
			list[count] = point
			count++
		}
	}
	return list[:count]
}
