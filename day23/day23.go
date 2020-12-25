package day23

import (
	"fmt"
	"strconv"
	"strings"
)

type CupsGame []int

func ReadCups(input string) []int {
	items := strings.Split(input, "")
	list := make([]int, len(items))
	for i, item := range items {
		cup, err := strconv.Atoi(item)
		if err != nil {
			panic(fmt.Sprintf("Bad day23 input: expected integer, got %v", item))
		}
		list[i] = cup
	}
	return list
}

func NewCupsGame(numbers []int) CupsGame {
	game := make([]int, len(numbers)+1)
	for i, item := range numbers {
		game[item] = numbers[(i+1)%len(numbers)]
	}
	game[0] = numbers[0]
	return game
}

func ReadInput(input string) CupsGame {
	return NewCupsGame(ReadCups(input))
}

func InList(list []int, item int) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func (cups CupsGame) Next(target int) int {
	return cups[target]
}

func (cups *CupsGame) SetNext(target, value int) {
	(*cups)[target] = value
}

func (cups CupsGame) Current() int {
	return cups[0]
}

// func (cups CupsGame) ValueBelow(target int) int {
// 	maxValue := len(cups.List)
// 	predecessor := func(x int) int {
// 		return ((x + maxValue - 2) % maxValue) + 1
// 	}
// 	result := predecessor(target)
// 	for cups.IndexOf(result) == -1 {
// 		result = predecessor(result)
// 	}
// 	return result
// }

func (cups CupsGame) Predecessor(target int) int {
	maxValue := len(cups) - 1
	return ((target + maxValue - 2) % maxValue) + 1
}

func (cups *CupsGame) Step() {
	current := cups.Current()
	removed := [3]int{}
	removed[0] = cups.Next(current)
	for i := 1; i <= 2; i++ {
		removed[i] = cups.Next(removed[i-1])
	}

	target := cups.Predecessor(current)
	for InList(removed[:], target) {
		target = cups.Predecessor(target)
	}
	afterTarget := cups.Next(target)
	cups.SetNext(current, cups.Next(removed[2]))
	cups.SetNext(target, removed[0])
	cups.SetNext(removed[2], afterTarget)

	(*cups)[0] = cups.Next(current)
}

func (cups *CupsGame) RunGame(rounds int) {
	for i := 0; i < rounds; i++ {
		cups.Step()
	}
}

func (cups CupsGame) Serialize(start int) []int {
	result := make([]int, len(cups)-1)
	current := start
	for i := range result {
		result[i] = current
		current = cups.Next(current)
	}
	return result
}

func (cups CupsGame) Score1() int {
	result := 0

	for _, item := range cups.Serialize(1)[1:] {
		result *= 10
		result += item
	}

	return result
}

func (cups CupsGame) Score2() int {
	next1 := cups.Next(1)
	next2 := cups.Next(next1)
	return next1 * next2
}

func Part1(input string, rounds int) int {
	game := ReadInput(input)
	game.RunGame(rounds)
	return game.Score1()
}

func Part2(input string) int {
	cupsList := ReadCups(input)
	gameSize := 1000000
	gameExtension := make([]int, gameSize)
	for i := 0; i < gameSize; i++ {
		gameExtension[i] = i + 1
	}
	cupsList = append(cupsList, gameExtension[len(cupsList):]...)

	cupsGame := NewCupsGame(cupsList)
	cupsGame.RunGame(10000000)
	return cupsGame.Score2()
}
