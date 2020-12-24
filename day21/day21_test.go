package day21_test

import (
	"HSteffensen/AoC2020/day21"
	"testing"
)

func TestPart1(t *testing.T) {
	if expected, result := 5, day21.Part1(day21.ExampleInput1); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	if expected, result := "mxmxvkd,sqjhc,fvjkl", day21.Part2(day21.ExampleInput1); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}
