package day23_test

import (
	"HSteffensen/AoC2020/day23"
	"reflect"
	"testing"
)

func TestReadInput(t *testing.T) {
	if expected, result := (day23.CupsGame{1, 2, 3, 4, 1}), day23.ReadInput("1234"); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame read input list: expected %v, got %v", expected, result)
	}
	if expected, result := (day23.CupsGame{3, 2, 5, 8, 6, 4, 7, 3, 9, 1}), day23.ReadInput("389125467"); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect CupsGame read input list: expected %v, got %v", expected, result)
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

func TestSerialize(t *testing.T) {
	if expected, result := ([]int{1, 2, 3, 4}), day23.ReadInput("1234").Serialize(1); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect serialize: expected %#v, got %#v", expected, result)
	}
	if expected, result := ([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}), day23.ReadInput("389125467").Serialize(3); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect serialize: expected %#v, got %#v", expected, result)
	}
	if expected, result := ([]int{1, 2, 5, 4, 6, 7, 3, 8, 9}), day23.ReadInput("389125467").Serialize(1); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect serialize: expected %#v, got %#v", expected, result)
	}
}

func TestPredecessor(t *testing.T) {
	game := day23.CupsGame{1, 2, 3, 4, 1}
	if expected, result := 3, game.Predecessor(4); expected != result {
		t.Errorf("Incorrect value below: expected %v, got %v", expected, result)
	}
	if expected, result := 4, game.Predecessor(1); expected != result {
		t.Errorf("Incorrect value below: expected %v, got %v", expected, result)
	}
	game = day23.CupsGame{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	if expected, result := 4, game.Predecessor(5); expected != result {
		t.Errorf("Incorrect value below: expected %v, got %v", expected, result)
	}
	if expected, result := 9, game.Predecessor(1); expected != result {
		t.Errorf("Incorrect value below: expected %v, got %v", expected, result)
	}
	game = day23.CupsGame{3, 2, 5, 8, 6, 4, 7, 3, 9, 1}
	if expected, result := 4, game.Predecessor(5); expected != result {
		t.Errorf("Incorrect value below: expected %v, got %v", expected, result)
	}
	if expected, result := 9, game.Predecessor(1); expected != result {
		t.Errorf("Incorrect value below: expected %v, got %v", expected, result)
	}
}
