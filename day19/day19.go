package day19

import (
	"strings"
)

type RuleName string
type RuleLiteral string
type RuleValue interface { // RuleName or RuleLiteral
	IsLiteral() bool // mostly to avoid 'RuleValue' being just 'interface{}' - idk if this is best practice but seems sane for type checking reasons
}
type RuleList []RuleValue
type Rule []RuleList
type Grammar map[RuleName]Rule

func (rn RuleName) IsLiteral() bool {
	return false
}

func (rn RuleLiteral) IsLiteral() bool {
	return true
}

func (g Grammar) NextSteps(input string, currentSteps RuleList) (string, []RuleList) {
	result := make([]RuleList, 0, 10)
	if len(currentSteps) > 0 {
		switch val := currentSteps[0].(type) {
		case RuleLiteral:
			valString := string(val)
			if strings.HasPrefix(input, valString) {
				input = input[len(valString):]
				result = append(result, currentSteps[1:])
			}
		case RuleName:
			for _, ruleOptions := range g[val] {
				nextResult := make(RuleList, 0, len(currentSteps)+2)
				nextResult = append(nextResult, ruleOptions...)
				nextResult = append(nextResult, currentSteps[1:]...)
				result = append(result, nextResult)
			}
		}
	}
	return input, result
}

func (g Grammar) Matches2(input string) bool {
	type queueItem struct {
		input string
		list  RuleList
	}
	queue := make([]queueItem, 0, 1000)
	var queueify = func(inp string, rl []RuleList) []queueItem {
		ql := make([]queueItem, 0, len(rl))
		for _, r := range rl {
			if len(r) > 0 {
				ql = append(ql, queueItem{inp, r})
			}
		}
		return ql
	}
	queue = append(queue, queueify(input, g[RuleName("0")])...)

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		nextString, nextQueues := g.NextSteps(next.input, next.list)
		if nextString == "" && (len(nextQueues) == 0 || (len(nextQueues) == 1 && len(nextQueues[0]) == 0)) {
			return true
		} else if nextString != "" && len(nextQueues) > 0 {
			queue = append(queue, queueify(nextString, nextQueues)...)
		}
	}

	return false
}

func CountMatches(gram Grammar, words []string) int {
	total := 0

	for _, word := range words {
		if gram.Matches2(word) {
			total++
		}
	}

	return total
}

func Part1(input string) int {
	gram := ReadRules(input)
	words := ReadWords(input)
	return CountMatches(gram, words)
}

func Part2(input string) int {
	gram := ReadRules(input)
	words := ReadWords(input)
	name8, rule8 := ReadRule("8: 42 | 42 8")
	name11, rule11 := ReadRule("11: 42 31 | 42 11 31")
	gram[name8] = rule8
	gram[name11] = rule11
	return CountMatches(gram, words)
}
