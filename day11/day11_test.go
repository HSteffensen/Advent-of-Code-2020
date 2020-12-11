package day11_test

import (
	"HSteffensen/AoC2020/day11"
	"fmt"
	"testing"
)

func TestSeatGridPos(t *testing.T) {
	sg := day11.NewSeatGrid([][]day11.Seat{{0, 0, 0}, {0, 1, 0}, {2, 0, 0}})
	t.Logf("\n%v", sg)
	if result, expected := sg.At(1, 1), day11.Empty; result != expected {
		t.Errorf("Incorrect value: expected %v, got %v", expected, result)
	}
	if result, expected := sg.At(0, 2), day11.Occupied; result != expected {
		t.Errorf("Incorrect value: expected %v, got %v", expected, result)
	}
	if result, expected := sg.At(-1, 3), day11.Floor; result != expected {
		t.Errorf("Incorrect value: expected %v, got %v", expected, result)
	}
}

func TestSeatGridAdjacent(t *testing.T) {
	sg1 := day11.NewSeatGrid([][]day11.Seat{{0, 0, 0}, {0, 1, 0}, {2, 0, 0}})
	if result, expected := sg1.Adjacent(1, 1), 1; result != expected {
		t.Errorf("Incorrect adjacent count: expected %v, got %v", expected, result)
	}
	sg2 := day11.NewSeatGrid([][]day11.Seat{{0, 0, 0}, {0, 2, 0}, {2, 0, 0}})
	if result, expected := sg2.Adjacent(1, 1), 1; result != expected {
		t.Errorf("Incorrect adjacent count: expected %v, got %v", expected, result)
	}
	sg3 := day11.NewSeatGrid([][]day11.Seat{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}})
	if result, expected := sg3.Adjacent(1, 1), 8; result != expected {
		t.Errorf("Incorrect adjacent count: expected %v, got %v", expected, result)
	}
}

func TestSeatGridCountOccupied(t *testing.T) {
	sg1 := day11.NewSeatGrid([][]day11.Seat{{0, 0, 0}, {0, 1, 0}, {2, 0, 0}})
	if result, expected := sg1.CountOccupied(), 1; result != expected {
		t.Errorf("Incorrect adjacent count: expected %v, got %v", expected, result)
	}
	sg2 := day11.NewSeatGrid([][]day11.Seat{{0, 0, 0}, {0, 2, 0}, {2, 0, 0}})
	if result, expected := sg2.CountOccupied(), 2; result != expected {
		t.Errorf("Incorrect adjacent count: expected %v, got %v", expected, result)
	}
	sg3 := day11.NewSeatGrid([][]day11.Seat{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}})
	if result, expected := sg3.CountOccupied(), 9; result != expected {
		t.Errorf("Incorrect adjacent count: expected %v, got %v", expected, result)
	}
}

func TestSeatGridStepForward(t *testing.T) {
	sg1 := day11.NewSeatGrid(day11.ReadInput(`L.L
	...
	#.#`))
	expected1 := `#.#
...
#.#`
	if result := fmt.Sprint(sg1.StepForwardPart1()); result != expected1 {
		t.Errorf("Incorrect step forward:\n- expected \n%v\n- got \n%v", expected1, result)
	}
}

func TestPart1(t *testing.T) {
	sg1 := day11.NewSeatGrid(day11.ReadInput(day11.ExampleInput))
	if result, expected := sg1.Part1(), 37; expected != result {
		t.Errorf("Incorrect part 1: expected %v got %v", expected, result)
	}
	if result, expected := fmt.Sprint(sg1), day11.ExampleInput; expected != result {
		t.Error("Incorrect part 1: modified the SeatGrid")
	}
}

func TestFirstVisibleFrom(t *testing.T) {
	sg1 := day11.NewSeatGrid(day11.ReadInput(`L.L
	...
	#.#`))
	eX, eY := 2, 0
	if x, y := sg1.FirstVisibleFrom(0, 0, 1, 0); x != eX || y != eY {
		t.Errorf("Incorrect visible from: expected (%v, %v), got (%v, %v)", eX, eY, x, y)
	}
	eX, eY = 2, 2
	if x, y := sg1.FirstVisibleFrom(0, 0, 1, 1); x != eX || y != eY {
		t.Errorf("Incorrect visible from: expected (%v, %v), got (%v, %v)", eX, eY, x, y)
	}
	eX, eY = 0, 2
	if x, y := sg1.FirstVisibleFrom(0, 0, 0, 1); x != eX || y != eY {
		t.Errorf("Incorrect visible from: expected (%v, %v), got (%v, %v)", eX, eY, x, y)
	}
	eX, eY = -1, -1
	if x, y := sg1.FirstVisibleFrom(0, 0, 0, -1); x != eX || y != eY {
		t.Errorf("Incorrect visible from: expected (%v, %v), got (%v, %v)", eX, eY, x, y)
	}
	eX, eY = -1, -1
	if x, y := sg1.FirstVisibleFrom(1, 1, 0, 1); x != eX || y != eY {
		t.Errorf("Incorrect visible from: expected (%v, %v), got (%v, %v)", eX, eY, x, y)
	}
}

func TestAdjacentVisibleFrom(t *testing.T) {
	sg1 := day11.NewSeatGrid(day11.ReadInput(`L.L
	...
	#.#`))
	if result, expected := sg1.AdjacentVisibleFrom(0, 0), 2; expected != result {
		t.Errorf("Incorrect AdjacentVisibleFrom: expected %v, got %v", expected, result)
	}
	if result, expected := sg1.AdjacentVisibleFrom(0, 2), 1; expected != result {
		t.Errorf("Incorrect AdjacentVisibleFrom: expected %v, got %v", expected, result)
	}
	if result, expected := sg1.AdjacentVisibleFrom(1, 1), 2; expected != result {
		t.Errorf("Incorrect AdjacentVisibleFrom: expected %v, got %v", expected, result)
	}
	if result, expected := sg1.AdjacentVisibleFrom(1, 0), 0; expected != result {
		t.Errorf("Incorrect AdjacentVisibleFrom: expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	sg1 := day11.NewSeatGrid(day11.ReadInput(day11.ExampleInput))
	if result, expected := sg1.Part2(), 26; expected != result {
		t.Errorf("Incorrect part 1: expected %v got %v", expected, result)
	}
	if result, expected := fmt.Sprint(sg1), day11.ExampleInput; expected != result {
		t.Error("Incorrect part 1: modified the SeatGrid")
	}
}
