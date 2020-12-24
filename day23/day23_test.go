package day23_test

import (
	"HSteffensen/AoC2020/day23"
	"reflect"
	"testing"
)

func TestReadInput(t *testing.T) {
	if expected, result := (day23.CupsGame{1, 2, 3, 4}), day23.ReadInput("1234"); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame read input: expected %v, got %v", expected, result)
	}
	if expected, result := (day23.CupsGame{3, 8, 9, 1, 2, 5, 4, 6, 7}), day23.ReadInput("389125467"); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame read input: expected %v, got %v", expected, result)
	}
}

func TestPart1(t *testing.T) {
	if expected, result := 92658374, day23.Part1("389125467", 10); expected != result {
		t.Errorf("Incorrect part 1: expected %#v, got %#v", expected, result)
	}
	if expected, result := 67384529, day23.Part1("389125467", 100); expected != result {
		t.Errorf("Incorrect part 1: expected %#v, got %#v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	if expected, result := 149245887792, day23.Part2("389125467"); expected != result {
		t.Errorf("Incorrect part 1: expected %#v, got %#v", expected, result)
	}
}

func TestCupsCycleTo(t *testing.T) {
	cups1 := day23.CupsGame{1, 2, 3, 4}
	cups1.CycleToValue(3)
	if expected, result := (day23.CupsGame{3, 4, 1, 2}), cups1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame cycle to: expected %v, got %v", expected, result)
	}
	cups1 = day23.CupsGame{1, 2, 3, 4}
	cups1.CycleToValue(1)
	if expected, result := (day23.CupsGame{1, 2, 3, 4}), cups1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame cycle to: expected %v, got %v", expected, result)
	}
}

func TestInsertCupsAfterFirst(t *testing.T) {
	cups1 := day23.CupsGame{1, 2, 3, 4}
	cups1.InsertCupsAfterFirst([]int{7})
	if expected, result := (day23.CupsGame{1, 7, 2, 3, 4}), cups1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame cycle to: expected %v, got %v", expected, result)
	}
	cups1 = day23.CupsGame{1, 2, 3, 4}
	cups1.InsertCupsAfterFirst([]int{7, 8, 9})
	if expected, result := (day23.CupsGame{1, 7, 8, 9, 2, 3, 4}), cups1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame cycle to: expected %v, got %v", expected, result)
	}
	cups1 = day23.CupsGame{1, 2, 3, 4}
	cups1.InsertCupsAfterFirst([]int{5, 6, 7, 8, 9})
	if expected, result := (day23.CupsGame{1, 5, 6, 7, 8, 9, 2, 3, 4}), cups1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame cycle to: expected %v, got %v", expected, result)
	}
	cups1 = day23.CupsGame{1}
	cups1.InsertCupsAfterFirst([]int{5, 6, 7, 8, 9})
	if expected, result := (day23.CupsGame{1, 5, 6, 7, 8, 9}), cups1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame cycle to: expected %v, got %v", expected, result)
	}
}

func TestRemoveCupsAfterFirst(t *testing.T) {
	cups1 := day23.CupsGame{1, 2, 3, 4}
	removed1 := cups1.RemoveCupsAfterFirst()
	if expected, result := (day23.CupsGame{1}), cups1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame remove: expected %v, got %v", expected, result)
	}
	if expected, result := ([]int{2, 3, 4}), removed1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame remove: expected %v, got %v", expected, result)
	}
	cups1 = day23.CupsGame{1, 2, 3, 4, 5, 6, 7, 8, 9}
	removed1 = cups1.RemoveCupsAfterFirst()
	if expected, result := (day23.CupsGame{1, 5, 6, 7, 8, 9}), cups1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame remove: expected %v, got %v", expected, result)
	}
	if expected, result := ([]int{2, 3, 4}), removed1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame remove: expected %v, got %v", expected, result)
	}
}

func TestValueBelow(t *testing.T) {
	game := day23.CupsGame{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if expected, result := 4, game.ValueBelow(5); expected != result {
		t.Errorf("Incorrect value below: expected %v, got %v", expected, result)
	}
	if expected, result := 9, game.ValueBelow(1); expected != result {
		t.Errorf("Incorrect value below: expected %v, got %v", expected, result)
	}
	game = day23.CupsGame{1, 5, 6, 7, 8, 9}
	if expected, result := 1, game.ValueBelow(5); expected != result {
		t.Errorf("Incorrect value below: expected %v, got %v", expected, result)
	}
	if expected, result := 9, game.ValueBelow(1); expected != result {
		t.Errorf("Incorrect value below: expected %v, got %v", expected, result)
	}
}
