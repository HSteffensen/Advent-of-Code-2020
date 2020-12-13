package day12_test

import (
	"HSteffensen/AoC2020/day12"
	"testing"
)

func TestPart1(t *testing.T) {
	if expected, result := 25, day12.Part1(day12.ExampleInput); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	if expected, result := 286, day12.Part2(day12.ExampleInput); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}
