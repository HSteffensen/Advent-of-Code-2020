package day1_test

import (
	"HSteffensen/AoC2020/aocCommon"
	"HSteffensen/AoC2020/day1"
	"fmt"
	"testing"
)

var TestInput1 = `1721
979
366
299
675
1456`

func TestDay1Part1(t *testing.T) {
	testData := aocCommon.ReadInputToInts(TestInput1)
	result1, result2 := day1.FindPairSummingToN(testData, 2020)
	want1, want2 := 1721, 299
	resultMap := make(map[int]bool)
	resultMap[want1] = false
	resultMap[want2] = false
	resultMap[result1] = true
	resultMap[result2] = true
	if resultMap[want1] && resultMap[want2] {
		fmt.Printf("Wanted: %v and %v. Got: %v and %v\n", want1, want2, result1, result2)
	} else {
		t.Errorf("Wanted: %v and %v. Got: %v and %v", want1, want2, result1, result2)
	}
}
func TestDay1Part2(t *testing.T) {
	testData := aocCommon.ReadInputToInts(TestInput1)
	result1, result2, result3 := day1.FindTripleSummingToN(testData, 2020)
	want1, want2, want3 := 979, 366, 675
	resultMap := make(map[int]bool)
	resultMap[want1] = false
	resultMap[want2] = false
	resultMap[want3] = false
	resultMap[result1] = true
	resultMap[result2] = true
	resultMap[result3] = true
	if resultMap[want1] && resultMap[want2] && resultMap[want3] {
		fmt.Printf("Wanted: %v, %v, and %v. Got: %v, %v, and %v\n", want1, want2, want3, result1, result2, result3)
	} else {
		t.Errorf("Wanted: %v, %v, and %v. Got: %v, %v, and %v\n", want1, want2, want3, result1, result2, result3)
	}
}
