package day24_test

import (
	"HSteffensen/AoC2020/day24"
	"testing"
)

func TestFollowPath(t *testing.T) {
	if expected, result := (day24.HexPoint{0, 0}), day24.FollowPath("nwwswee"); expected != result {
		t.Errorf("Incorrect follow path: expected %v, got %v", expected, result)
	}
	if expected, result := (day24.HexPoint{0, 1}), day24.FollowPath("esew"); expected != result {
		t.Errorf("Incorrect follow path: expected %v, got %v", expected, result)
	}
	if expected, result := (day24.HexPoint{0, 4}), day24.FollowPath("sesesese"); expected != result {
		t.Errorf("Incorrect follow path: expected %v, got %v", expected, result)
	}
	if expected, result := day24.FollowPath("ee"), day24.FollowPath("sesenene"); expected != result {
		t.Errorf("Incorrect follow path: expected %v, got %v", expected, result)
	}
}

func TestPart1(t *testing.T) {
	if expected, result := 10, day24.Part1(day24.ExampleInput1); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}

func TestStep(t *testing.T) {
	grid := make(day24.HexGrid)
	grid[day24.HexPoint{0, 0}] = true
	grid[day24.HexPoint{0, 1}] = true
	if expected, result := 4, grid.Step().CountBlackTiles(); expected != result {
		t.Errorf("Incorrect step: expected %v, got %v", expected, result)
	}

	grid = make(day24.HexGrid)
	grid[day24.HexPoint{0, 0}] = true
	if expected, result := 0, grid.Step().CountBlackTiles(); expected != result {
		t.Errorf("Incorrect step: expected %v, got %v", expected, result)
	}

	grid = make(day24.HexGrid)
	grid[day24.HexPoint{0, 0}] = true
	grid[day24.HexPoint{0, 1}] = true
	grid[day24.HexPoint{1, 0}] = true
	if expected, result := 3, grid.Step().CountBlackTiles(); expected != result {
		t.Errorf("Incorrect step: expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	if expected, result := 2208, day24.Part2(day24.ExampleInput1); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}
