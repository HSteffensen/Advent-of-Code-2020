package day17_test

import (
	"HSteffensen/AoC2020/day17"
	"testing"
)

func TestCubeGridPosition(t *testing.T) {

}

func TestPart1(t *testing.T) {
	data := day17.ReadInput(day17.ExampleInput)
	if expected, result := 112, day17.Part1(data, 6); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	data := day17.ReadInput(day17.ExampleInput)
	if expected, result := 848, day17.Part2(data, 6); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}
