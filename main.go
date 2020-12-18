package main

import (
	"HSteffensen/AoC2020/aocCommon"
	"HSteffensen/AoC2020/day1"
	"HSteffensen/AoC2020/day10"
	"HSteffensen/AoC2020/day11"
	"HSteffensen/AoC2020/day12"
	"HSteffensen/AoC2020/day13"
	"HSteffensen/AoC2020/day14"
	"HSteffensen/AoC2020/day2"
	"HSteffensen/AoC2020/day3"
	"HSteffensen/AoC2020/day4"
	"HSteffensen/AoC2020/day5"
	"HSteffensen/AoC2020/day6"
	"HSteffensen/AoC2020/day7"
	"HSteffensen/AoC2020/day8"
	"HSteffensen/AoC2020/day9"
	"fmt"
	"strings"
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

func runDay3() {
	data := day3.ReadTerrainGrid(day3.Input)
	result1 := data.TreesOnAngle(3, 1)
	fmt.Println("Day 3, Part 1 answer:", result1)
	result2 := data.TreesOnAngle(1, 1) * data.TreesOnAngle(3, 1) * data.TreesOnAngle(5, 1) * data.TreesOnAngle(7, 1) * data.TreesOnAngle(1, 2)
	fmt.Println("Day 3, Part 2 answer:", result2)
}

func runDay4() {
	data := day4.ReadPassports(day4.Input)
	result1 := day4.CountValidPassports(data)
	fmt.Println("Day 4, Part 1 answer:", result1)
	result2 := day4.CountValidPassportsValues(data)
	fmt.Println("Day 4, Part 2 answer:", result2)
}

func runDay5() {
	data := strings.Split(day5.Input, "\n")
	result1 := day5.MaxSeatId(data)
	fmt.Println("Day 5, Part 1 answer:", result1)
	result2 := day5.MySeatId(data)
	fmt.Println("Day 5, Part 2 answer:", result2)
}

func runDay6() {
	data := day6.Groups(day6.Input)
	result1 := day6.Part1(data)
	fmt.Println("Day 6, Part 1 answer:", result1)
	result2 := day6.Part2(data)
	fmt.Println("Day 6, Part 2 answer:", result2)
}

func runDay7() {
	data := day7.ReadInput(day7.Input)
	result1 := len(data.BagsWhichCanContain("shiny gold"))
	fmt.Println("Day 7, Part 1 answer:", result1)
	result2 := data.BagsWithin("shiny gold")
	fmt.Println("Day 7, Part 2 answer:", result2)
}

func runDay8() {
	data := day8.ReadInput(day8.Input)
	result1 := data.FindLoopAcc()
	fmt.Println("Day 8, Part 1 answer:", result1)
	result2 := data.FixLoopAcc()
	fmt.Println("Day 8, Part 2 answer:", result2)
}

func runDay9() {
	data := aocCommon.ReadInputToInts(day9.Input)
	result1 := data[day9.FirstInvalidIndex(data, 25)]
	fmt.Println("Day 9, Part 1 answer:", result1)
	result2 := day9.FindEncryptionWeakness(data, 25)
	fmt.Println("Day 9, Part 2 answer:", result2)
}

func runDay10() {
	data := aocCommon.ReadInputToInts(day10.Input)
	result1 := day10.Part1(data)
	fmt.Println("Day 10, Part 1 answer:", result1)
	result2 := day10.CountLegalArrangements(data)
	fmt.Println("Day 10, Part 2 answer:", result2)
}

func runDay11() {
	data := day11.NewSeatGrid(day11.ReadInput(day11.Input))
	result1 := data.Part1()
	fmt.Println("Day 11, Part 1 answer:", result1)
	result2 := data.Part2()
	fmt.Println("Day 11, Part 2 answer:", result2)
}

func runDay12() {
	result1 := day12.Part1(day12.Input)
	fmt.Println("Day 12, Part 1 answer:", result1)
	result2 := day12.Part2(day12.Input)
	fmt.Println("Day 12, Part 2 answer:", result2)
}

func runDay13() {
	result1 := day13.Part1(day13.Input)
	fmt.Println("Day 13, Part 1 answer:", result1)
	result2 := day13.Part2(day13.Input)
	fmt.Println("Day 13, Part 2 answer:", result2)
}

func runDay14() {
	result1 := day14.Part1(day14.Input)
	fmt.Println("Day 14, Part 1 answer:", result1)
	result2 := day14.Part2(day14.Input)
	fmt.Println("Day 14, Part 2 answer:", result2)
}

func main() {
	runDay1()
	runDay2()
	runDay3()
	runDay4()
	runDay5()
	runDay6()
	runDay7()
	runDay8()
	runDay9()
	runDay10()
	runDay11()
	runDay12()
	runDay13()
	runDay14()
}
