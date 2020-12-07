package day2_test

import (
	"HSteffensen/AoC2020/day2"
	"testing"
)

var TestInput1 = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestDay2Part1(t *testing.T) {
	data := day2.InputToPasswordVal(TestInput1)
	result := day2.CountValidPasswords(data, day2.CheckPasswordPart1)
	if result != 2 {
		t.Errorf("Incorrect number of valid passwords: got %v, expected %v", result, 2)
	}
}

func TestDay2Part2(t *testing.T) {
	data := day2.InputToPasswordVal(TestInput1)
	result := day2.CountValidPasswords(data, day2.CheckPasswordPart2)
	if result != 1 {
		t.Errorf("Incorrect number of valid passwords: got %v, expected %v", result, 1)
	}
}
