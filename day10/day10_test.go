package day10_test

import (
	"HSteffensen/AoC2020/aocCommon"
	"HSteffensen/AoC2020/day10"
	"testing"
)

func TestCountJoltageJumps(t *testing.T) {
	data1 := aocCommon.ReadInputToInts(day10.ExampleInput1)
	if result1, _, result3 := day10.CountJoltageJumps(data1); result1 != 7 || result3 != 5 {
		t.Errorf("Incorrect joltage jumps: expected (%v, %v), got (%v, %v)", 7, 5, result1, result3)
	}
	data2 := aocCommon.ReadInputToInts(day10.ExampleInput2)
	if result1, _, result3 := day10.CountJoltageJumps(data2); result1 != 22 || result3 != 10 {
		t.Errorf("Incorrect joltage jumps: expected (%v, %v), got (%v, %v)", 22, 10, result1, result3)
	}
}

func TestCountLegalArrangements(t *testing.T) {
	data1 := aocCommon.ReadInputToInts(day10.ExampleInput1)
	if result, expected := day10.CountLegalArrangements(data1), 8; result != expected {
		t.Errorf("Incorrect arrangements count: expected %v, got %v", expected, result)
	}
	data2 := aocCommon.ReadInputToInts(day10.ExampleInput2)
	if result, expected := day10.CountLegalArrangements(data2), 19208; result != expected {
		t.Errorf("Incorrect arrangements count: expected %v, got %v", expected, result)
	}
	data3 := []int{}
	expecteds := []int{1, 1, 2, 4, 7}
	if result, expected := day10.CountLegalArrangements(data3), expecteds[0]; result != expected {
		t.Errorf("Incorrect arrangements count: expected %v, got %v", expected, result)
	}
	for i := 1; i < 5; i++ {
		data3 = append(data3, i)
		if result, expected := day10.CountLegalArrangements(data3), expecteds[i]; result != expected {
			t.Errorf("Incorrect arrangements count: expected %v, got %v", expected, result)
		}
	}
}
