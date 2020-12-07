package main

import (
	"HSteffensen/AoC2020/aocCommon"
	"HSteffensen/AoC2020/day1"
	"HSteffensen/AoC2020/day2"
	"fmt"
)

func runDay1() {
	data := aocCommon.ReadInputToInts(day1.Input1)
	result1, result2 := day1.FindPairSummingToN(data, 2020)
	fmt.Println("Day 1, Part 1 answer:", result1*result2)
	result1, result2, result3 := day1.FindTripleSummingToN(data, 2020)
	fmt.Println("Day 1, Part 2 answer:", result1*result2*result3)
}

func runDay2() {
	data := day2.InputToPasswordVal(day2.Input1)
	result1 := day2.CountValidPasswords(data, day2.CheckPasswordPart1)
	fmt.Println("Day 2, Part 1 answer:", result1)
	result2 := day2.CountValidPasswords(data, day2.CheckPasswordPart2)
	fmt.Println("Day 2, Part 2 answer:", result2)
}

func main() {
	runDay1()
	runDay2()
}
