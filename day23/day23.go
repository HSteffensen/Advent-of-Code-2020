package day23

import (
	"fmt"
	"strconv"
	"strings"
)

type CupsGame []int

func ReadInput(input string) CupsGame {
	items := strings.Split(input, "")
	game := make(CupsGame, len(items))
	for i, item := range items {
		cup, err := strconv.Atoi(item)
		if err != nil {
			panic(fmt.Sprintf("Bad day23 input: expected integer, got %v", item))
		}
		game[i] = cup
	}
	return game
}

func (cups CupsGame) IndexOf(target int) int {
	index := -1
	for i, cup := range cups {
		if cup == target {
			index = i
			break
		}
	}
	return index
}

func (cups *CupsGame) CycleToIndex(index int) {
	*cups = append((*cups)[index:], (*cups)[:index]...)
}

func (cups *CupsGame) CycleToValue(target int) {
	index := cups.IndexOf(target)
	if index == -1 {
		panic(fmt.Sprintf("Bad CupsGame: tried to cycle to %v but it wasn't there in %v", target, *cups))
	}
	cups.CycleToIndex(index)
}

func (cups *CupsGame) InsertCupsAfterFirst(newCups []int) {
	*cups = append(*cups, newCups...)
	copy((*cups)[len(newCups)+1:], (*cups)[1:])
	copy((*cups)[1:len(newCups)+1], newCups)
	// *cups = append((*cups)[:1], append(newCups, (*cups)[1:]...)...) // this way requires more shuffling and allocation
}

func (cups *CupsGame) RemoveCupsAfterFirst() []int {
	result := make([]int, 3)
	copy(result, (*cups)[1:4])
	copy((*cups)[1:], (*cups)[4:])
	*cups = (*cups)[:len(*cups)-3]
	return result
}

func (cups CupsGame) ValueBelow(target int) int {
	maxValue := -1
	for _, cup := range cups {
		if cup > maxValue {
			maxValue = cup
		}
	}
	predecessor := func(x int) int {
		return ((x + maxValue - 2) % maxValue) + 1
	}
	result := predecessor(target)
	for cups.IndexOf(result) == -1 {
		result = predecessor(result)
	}
	return result
}

func (cups *CupsGame) Step() {
	current := (*cups)[0]
	removed := cups.RemoveCupsAfterFirst()
	nextCup := cups.ValueBelow((*cups)[0])
	cups.CycleToValue(nextCup)
	cups.InsertCupsAfterFirst(removed)
	cups.CycleToValue(current)
	cups.CycleToIndex(1)
}

func (cups *CupsGame) RunGame(rounds int) {
	for i := 0; i < rounds; i++ {
		cups.Step()
		if (i+1)%100 == 0 {
			fmt.Printf("Done round %v\n", i+1)
		}
	}
}

func Part1(input string, rounds int) int {
	game := ReadInput(input)
	game.RunGame(rounds)
	game.CycleToValue(1)
	result := 0
	for _, cup := range game[1:] {
		result *= 10
		result += cup
	}
	return result
}

func Part2(input string) int {
	rounds := 10000000
	game := ReadInput(input)
	gameSize := 1000000
	gameExtension := make([]int, gameSize)
	for i := 0; i < gameSize; i++ {
		gameExtension[i] = i + 1
	}
	game = append(game, gameExtension[len(game):]...)
	fmt.Println(game[:20])
	game.RunGame(rounds)
	game.CycleToValue(1)
	return game[1] * game[2]
}
