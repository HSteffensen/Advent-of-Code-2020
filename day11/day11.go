package day11

import (
	"fmt"
	"strings"
)

type Seat int

type SeatGrid struct {
	grid   [][]Seat
	Width  int
	Height int
}

const (
	Floor Seat = iota
	Empty
	Occupied
)

func (s Seat) String() string {
	switch s {
	case Floor:
		return "."
	case Empty:
		return "L"
	case Occupied:
		return "#"
	default:
		panic(fmt.Sprintf("Bad Seat value: %#v", s))
	}
}

func (sg SeatGrid) String() string {
	list := make([]string, sg.Height)
	for y := 0; y < sg.Height; y++ {
		listX := make([]string, sg.Width)
		for x := 0; x < sg.Width; x++ {
			listX[x] = fmt.Sprint(sg.At(x, y))
		}
		list[y] = strings.Join(listX, "")
	}
	return strings.Join(list, "\n")
}

func NewEmptySeatGrid(width, height int) SeatGrid {
	sg := new(SeatGrid)
	sg.grid = make([][]Seat, height)
	for i := range sg.grid {
		sg.grid[i] = make([]Seat, width)
	}
	sg.Width = width
	sg.Height = height
	return *sg
}

func NewSeatGrid(seats [][]Seat) SeatGrid {
	sg := NewEmptySeatGrid(len(seats[0]), len(seats))
	for y := 0; y < sg.Height; y++ {
		gridY := seats[y]
		if len(gridY) > sg.Width {
			gridY = gridY[:sg.Width]
		}
		copy(sg.grid[y], gridY)
	}
	return sg
}

func (sg SeatGrid) At(x, y int) Seat {
	if sg.PosWithinBounds(x, y) {
		return sg.grid[y][x]
	}
	return Floor
}

func (sg SeatGrid) Set(x, y int, val Seat) {
	sg.grid[y][x] = val
}

func (sg SeatGrid) PosWithinBounds(x, y int) bool {
	return sg.XWithinBounds(x) && sg.YWithinBounds(y)
}

func (sg SeatGrid) XWithinBounds(x int) bool {
	return x >= 0 && x < sg.Width
}

func (sg SeatGrid) YWithinBounds(y int) bool {
	return y >= 0 && y < sg.Height
}

func (sg SeatGrid) Adjacent(x, y int) int {
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if sg.At(x+i, y+j) == Occupied && (i != 0 || j != 0) {
				count++
			}
		}
	}
	return count
}

func (sg SeatGrid) FirstVisibleFrom(x, y, dx, dy int) (int, int) {
	if dx != 0 || dy != 0 {
		x += dx
		y += dy
		for sg.XWithinBounds(x) && sg.YWithinBounds(y) {
			if sg.At(x, y) != Floor {
				return x, y
			}
			x += dx
			y += dy
		}
	} else {
		return x, y
	}
	return -1, -1
}

func (sg SeatGrid) AdjacentVisibleFrom(x, y int) int {
	count := 0
	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			if (dx != 0 || dy != 0) && sg.At(sg.FirstVisibleFrom(x, y, dx, dy)) == Occupied {
				count++
			}
		}
	}
	return count
}

func (sg SeatGrid) CountOccupied() int {
	count := 0
	for y := 0; y < sg.Height; y++ {
		for x := 0; x < sg.Width; x++ {
			if sg.At(x, y) == Occupied {
				count++
			}
		}
	}
	return count
}

func (s Seat) Next(neighborCount, threshold int) Seat {
	switch {
	case s == Empty && neighborCount == 0:
		return Occupied
	case s == Occupied && neighborCount >= threshold:
		return Empty
	}
	return s
}

func (sg SeatGrid) StepForwardPart1() SeatGrid {
	nextSG := NewEmptySeatGrid(sg.Width, sg.Height)
	for y := 0; y < sg.Height; y++ {
		for x := 0; x < sg.Width; x++ {
			nextSG.Set(x, y, sg.At(x, y).Next(sg.Adjacent(x, y), 4))
		}
	}
	return nextSG
}

func (sg SeatGrid) Part1() int {
	prevCount, currCount := -1, sg.CountOccupied()
	for prevCount != currCount {
		prevCount = currCount
		sg = sg.StepForwardPart1()
		currCount = sg.CountOccupied()
	}
	return currCount
}

func (sg SeatGrid) StepForwardPart2() SeatGrid {
	nextSG := NewEmptySeatGrid(sg.Width, sg.Height)
	for y := 0; y < sg.Height; y++ {
		for x := 0; x < sg.Width; x++ {
			nextSG.Set(x, y, sg.At(x, y).Next(sg.AdjacentVisibleFrom(x, y), 5))
		}
	}
	return nextSG
}

func (sg SeatGrid) Part2() int {
	prevCount, currCount := -1, sg.CountOccupied()
	for prevCount != currCount {
		prevCount = currCount
		sg = sg.StepForwardPart2()
		currCount = sg.CountOccupied()
	}
	return currCount
}
