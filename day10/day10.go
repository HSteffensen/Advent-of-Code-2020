package day10

import (
	"sort"
)

func CountJoltageJumps(adapters []int) (int, int, int) {
	sort.Ints(adapters)
	counter := make(map[int]int)
	counter[adapters[0]]++
	for i := 0; i < len(adapters)-1; i++ {
		counter[adapters[i+1]-adapters[i]]++
	}
	counter[3]++
	return counter[1], counter[2], counter[3]
}

func Part1(numbers []int) int {
	a, _, b := CountJoltageJumps(numbers)
	return a * b
}

func CountLegalArrangements(adapters []int) int {
	sort.Ints(adapters)
	adapters = append([]int{0}, adapters...)
	counts := make(map[int]int)
	counts[0] = 1

	for i := 1; i < len(adapters); i++ {
		val := adapters[i]
		for j := 1; j <= 3 && val-j >= 0; j++ {
			counts[val] += counts[val-j]
		}
	}
	return counts[adapters[len(adapters)-1]]
}
