package day19

import (
	"strings"
)

type RuleName string
type RuleLiteral string
type RuleValue interface { // RuleName or RuleLiteral
	IsLiteral() bool // mostly to avoid 'RuleValue' being just 'interface{}'
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

type matchMemoKey struct {
	name  RuleName
	input string
}

type matchMemoValue struct {
	result    bool
	remaining string
}

func (g Grammar) helperMatches(input string, currentRule RuleName, memo map[matchMemoKey]matchMemoValue, depth int) (bool, string) {
	if (string(currentRule) == "8" || string(currentRule) == "11") && depth > 100 {
		return false, input
	}
	memoKey := matchMemoKey{currentRule, input}
	if memoVal, ok := memo[memoKey]; ok {
		return memoVal.result, memoVal.remaining
	}
	startRule := g[currentRule]

CHOICES:
	for _, choice := range startRule {
		branchResult := true
		branchInput := input
		for _, item := range choice {
			switch val := item.(type) {
			case RuleLiteral:
				valString := string(val)
				if strings.HasPrefix(branchInput, valString) {
					branchInput = branchInput[len(valString):]
				} else {
					continue CHOICES
				}
			case RuleName:
				result, bInput := g.helperMatches(branchInput, val, memo, depth+1)
				branchInput = bInput
				branchResult = branchResult && result
			}
		}
		if branchResult {
			memoVal := matchMemoValue{true, branchInput}
			memo[memoKey] = memoVal
			return true, branchInput
		}
	}

	memoVal := matchMemoValue{false, input}
	memo[memoKey] = memoVal
	return false, input
}

func (g Grammar) Matches(input string) bool {
	result, remaining := g.helperMatches(input, RuleName("0"), map[matchMemoKey]matchMemoValue{}, 0)

	return result && remaining == ""
}

func (g Grammar) NextSteps(input string, currentSteps RuleList) (string, []RuleList) {
	result := make([]RuleList, 0, 2)
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
	queue := make([]queueItem, 0, 1000000)
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
