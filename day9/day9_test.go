package day9_test

import (
	"HSteffensen/AoC2020/aocCommon"
	"HSteffensen/AoC2020/day9"
	"testing"
)

func TestValidNextNumber(t *testing.T) {
	data1 := []int{1, 2}
	if !day9.ValidNextNumber(data1, 3) {
		t.Error("Incorrect result test 1")
	}
	if day9.ValidNextNumber(data1, 4) {
		t.Error("Incorrect result test 2")
	}
	data2 := aocCommon.ReadInputToInts(day9.ExampleInput)[:7]
	if !day9.ValidNextNumber(data2[:6], data2[6]) {
		t.Error("Incorrect result test 3")
	}
}

func TestFirstInvalidIndex(t *testing.T) {
	data1 := []int{1, 2, 4}
	if actual, expected := data1[day9.FirstInvalidIndex(data1, 2)], 4; actual != expected {
		t.Errorf("Incorrect result: expected %v, got %v", expected, actual)
	}
	data2 := aocCommon.ReadInputToInts(day9.ExampleInput)
	if actual, expected := data2[day9.FirstInvalidIndex(data2, 5)], 127; actual != expected {
		t.Errorf("Incorrect result: expected %v, got %v", expected, actual)
	}
}

func TestFindEncryptionWeakness(t *testing.T) {
	data1 := []int{1, 2, 3, 5, 5}
	if actual, expected := day9.FindEncryptionWeakness(data1, 2), 5; actual != expected {
		t.Errorf("Incorrect result: expected %v, got %v", expected, actual)
	}
	data2 := aocCommon.ReadInputToInts(day9.ExampleInput)
	if actual, expected := day9.FindEncryptionWeakness(data2, 5), 62; actual != expected {
		t.Errorf("Incorrect result: expected %v, got %v", expected, actual)
	}
}
