package day19

import "strings"

type RuleName string
type RuleLiteral string
type RuleValue interface { // RuleName or RuleLiteral
	IsLiteral() bool // mostly to avoid 'RuleValue' being just 'interface{}'
}
type RuleChoice []RuleValue
type Rule []RuleChoice
type Grammar map[RuleName]Rule

func (rn RuleName) IsLiteral() bool {
	return false
}

func (rn RuleLiteral) IsLiteral() bool {
	return true
}

func (g Grammar) helperMatches(input string, currentRule RuleName, depth int) (bool, string) {
	if (string(currentRule) == "8" || string(currentRule) == "11") && depth > 10000 {
		return false, input
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
				result, bInput := g.helperMatches(branchInput, val, depth+1)
				branchInput = bInput
				branchResult = branchResult && result
			}
		}
		if branchResult {
			return true, branchInput
		}
	}
	return false, input
}

func (g Grammar) Matches(input string) bool {
	result, remaining := g.helperMatches(input, RuleName("0"), 0)

	return result && remaining == ""
}

func CountMatches(gram Grammar, words []string) int {
	total := 0

	for _, word := range words {
		if gram.Matches(word) {
			total += 1
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
