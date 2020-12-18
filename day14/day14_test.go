package day14_test

import (
	"HSteffensen/AoC2020/aocCommon"
	"HSteffensen/AoC2020/day14"
	"fmt"
	"reflect"
	"testing"
)

func TestReadMask(t *testing.T) {
	mask := day14.ReadMask("X")
	if eZeroes, eOnes := 0, 0; mask.Zeroes != eZeroes || mask.Ones != eOnes {
		t.Errorf("Incorrect mask: expected (%v, %v), got (%v, %v)", eZeroes, eOnes, mask.Zeroes, mask.Ones)
	}
	mask = day14.ReadMask("1")
	if eZeroes, eOnes := 0, 1; mask.Zeroes != eZeroes || mask.Ones != eOnes {
		t.Errorf("Incorrect mask: expected (%v, %v), got (%v, %v)", eZeroes, eOnes, mask.Zeroes, mask.Ones)
	}
	mask = day14.ReadMask("0")
	if eZeroes, eOnes := 1, 0; mask.Zeroes != eZeroes || mask.Ones != eOnes {
		t.Errorf("Incorrect mask: expected (%v, %v), got (%v, %v)", eZeroes, eOnes, mask.Zeroes, mask.Ones)
	}
	mask = day14.ReadMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	if eZeroes, eOnes := 2, 64; mask.Zeroes != eZeroes || mask.Ones != eOnes {
		t.Errorf("Incorrect mask: expected (%v, %v), got (%v, %v)", eZeroes, eOnes, mask.Zeroes, mask.Ones)
	}
}

func TestApplyMask(t *testing.T) {
	mask, value := day14.Mask{0, 0}, 24
	if expected, result := 24, mask.ApplyTo(value); expected != result {
		t.Errorf("Incorrect mask apply: expected %v, got %v", expected, result)
	}
	mask, value = day14.Mask{1, 0}, 25
	if expected, result := 24, mask.ApplyTo(value); expected != result {
		t.Errorf("Incorrect mask apply: expected %v, got %v", expected, result)
	}
	mask, value = day14.Mask{0, 1}, 24
	if expected, result := 25, mask.ApplyTo(value); expected != result {
		t.Errorf("Incorrect mask apply: expected %v, got %v", expected, result)
	}
}

func TestPart1(t *testing.T) {
	if expected, result := 165, day14.Part1(day14.ExampleInput1); expected != result {
		t.Errorf("Incorrect day14 part1: expected %v, got %v", expected, result)
	}
}

func AssertSameMaskSlice(left, right []day14.Mask) bool {
	leftStrings := make([]string, len(left))
	rightStrings := make([]string, len(right))

	for i, mask := range left {
		leftStrings[i] = fmt.Sprint(mask)
	}
	for i, mask := range right {
		rightStrings[i] = fmt.Sprint(mask)
	}
	return aocCommon.AssertSameStringSlice(leftStrings, rightStrings)
}

func TestMultiMask(t *testing.T) {
	expected := []day14.Mask{{0b0, 0b0}}
	result := day14.MultiMask("0")
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect day14 MultiMask: expected %v, got %v", expected, result)
	}
	expected = []day14.Mask{{0b0, 0b1}}
	result = day14.MultiMask("1")
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect day14 MultiMask: expected %v, got %v", expected, result)
	}
	expected = []day14.Mask{{0b1, 0b0}, {0b0, 0b1}}
	result = day14.MultiMask("X")
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect day14 MultiMask: expected %v, got %v", expected, result)
	}
	expected = []day14.Mask{{0b10, 0b01}, {0b00, 0b11}}
	result = day14.MultiMask("X1")
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect day14 MultiMask: expected %v, got %v", expected, result)
	}
	expected = []day14.Mask{{0b10, 0b01}, {0b00, 0b11}, {0b11, 0b00}, {0b01, 0b10}}
	result = day14.MultiMask("XX")
	if !AssertSameMaskSlice(expected, result) {
		t.Errorf("Incorrect day14 MultiMask: expected %v, got %v", expected, result)
	}
	expected = []day14.Mask{{0b100001, 0b010010}, {0b100000, 0b010011}, {0b000001, 0b110010}, {0b000000, 0b110011}}
	result = day14.MultiMask("X1001X")
	if !AssertSameMaskSlice(expected, result) {
		t.Errorf("Incorrect day14 MultiMask: expected %v, got %v", expected, result)
	}
	expected = []day14.Mask{{0b100001, 0b010010}, {0b100000, 0b010011}, {0b000001, 0b110010}, {0b000000, 0b110011}}
	result = day14.MultiMask("000000000000000000000000000000X1001X")
	if !AssertSameMaskSlice(expected, result) {
		t.Errorf("Incorrect day14 MultiMask: expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	if expected, result := 208, day14.Part2(day14.ExampleInput2); expected != result {
		t.Errorf("Incorrect day14 part1: expected %v, got %v", expected, result)
	}
}
