package day6

import "strings"

func GroupAnswers(group []string) string {
	charUsed := make(map[rune]bool)
	for _, char := range strings.Join(group, "") {
		charUsed[char] = true
	}
	answers := ""
	for c := 'a'; c <= 'z'; c++ {
		if _, ok := charUsed[c]; ok {
			answers += string(c)
		}
	}
	return answers
}

func GroupAnswersEveryone(group []string) string {
	charUsed := make(map[rune]int)
	for _, char := range strings.Join(group, "") {
		charUsed[char]++
	}
	answers := ""
	for c := 'a'; c <= 'z'; c++ {
		if charUsed[c] == len(group) {
			answers += string(c)
		}
	}
	return answers
}

func Part1(groups [][]string) int {
	sum := 0
	for _, g := range groups {
		sum += len(GroupAnswers(g))
	}
	return sum
}

func Part2(groups [][]string) int {
	sum := 0
	for _, g := range groups {
		sum += len(GroupAnswersEveryone(g))
	}
	return sum
}
