package day19_test

import (
	"HSteffensen/AoC2020/day19"
	"reflect"
	"testing"
)

func TestReadRule(t *testing.T) {
	name, rule := day19.ReadRule("0: 4 1 5")
	if expected, result := "0", string(name); expected != result {
		t.Errorf("Incorrect rule name: expected %#v, got %#v", expected, result)
	}
	if expected, result := (day19.Rule{{day19.RuleName("4"), day19.RuleName("1"), day19.RuleName("5")}}), rule; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect rule: expected %#v, got %#v", expected, result)
	}

	name, rule = day19.ReadRule("1: 2 3 | 3 2")
	if expected, result := "1", string(name); expected != result {
		t.Errorf("Incorrect rule name: expected %#v, got %#v", expected, result)
	}
	if expected, result := (day19.Rule{{day19.RuleName("2"), day19.RuleName("3")}, {day19.RuleName("3"), day19.RuleName("2")}}), rule; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect rule: expected %#v, got %#v", expected, result)
	}

	name, rule = day19.ReadRule(`4: "a"`)
	if expected, result := "4", string(name); expected != result {
		t.Errorf("Incorrect rule name: expected %#v, got %#v", expected, result)
	}
	if expected, result := (day19.Rule{{day19.RuleLiteral("a")}}), rule; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect rule: expected %#v, got %#v", expected, result)
	}

	name, rule = day19.ReadRule(`4: "a" | 3 2`)
	if expected, result := "4", string(name); expected != result {
		t.Errorf("Incorrect rule name: expected %#v, got %#v", expected, result)
	}
	if expected, result := (day19.Rule{{day19.RuleLiteral("a")}, {day19.RuleName("3"), day19.RuleName("2")}}), rule; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect rule: expected %#v, got %#v", expected, result)
	}
}

func TestGrammarMatches(t *testing.T) {
	if !day19.ReadRules(`0: "a"`).Matches("a") {
		t.Error("Incorrect match: string should match grammar.")
	}

	if !day19.ReadRules(`0: 1
1: "a"`).Matches("a") {
		t.Error("Incorrect match: string should match grammar.")
	}

	if !day19.ReadRules(`0: 1
1: 2
2: "a"`).Matches("a") {
		t.Error("Incorrect match: string should match grammar.")
	}

	if !day19.ReadRules(`0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"`).Matches("aab") {
		t.Error("Incorrect match: string should match grammar.")
	}

	if !day19.ReadRules(`0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"`).Matches("aba") {
		t.Error("Incorrect match: string should match grammar.")
	}

	if !day19.ReadRules(day19.ExampleInput1).Matches("ababbb") {
		t.Error("Incorrect match: string should match grammar.")
	}
	if !day19.ReadRules(day19.ExampleInput1).Matches("abbbab") {
		t.Error("Incorrect match: string should match grammar.")
	}
	if day19.ReadRules(day19.ExampleInput1).Matches("bababa") {
		t.Error("Incorrect match: string should NOT match grammar.")
	}
	if day19.ReadRules(day19.ExampleInput1).Matches("aaabbb") {
		t.Error("Incorrect match: string should NOT match grammar.")
	}
	if day19.ReadRules(day19.ExampleInput1).Matches("aaaabbb") {
		t.Error("Incorrect match: string should NOT match grammar.")
	}

}

func TestPart1(t *testing.T) {
	if expected, result := 2, day19.Part1(day19.ExampleInput1); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	if expected, result := 3, day19.Part1(day19.ExampleInput2); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
	if expected, result := 12, day19.Part2(day19.ExampleInput2); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}

func TestGrammarNextSteps(t *testing.T) {
	gram1 := day19.ReadRules(`0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"`)

	nextString, nextList := gram1.NextSteps("aab", gram1[day19.RuleName("1")][0])
	if expected, result := "ab", nextString; expected != result {
		t.Errorf("Incorrect next string: expected %#v, got %#v", expected, result)
	}
	if expected, result := []day19.RuleList{{}}, nextList; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect next list: expected %v, got %v", expected, result)
	}

	nextString, nextList = gram1.NextSteps("aab", gram1[day19.RuleName("0")][0])
	if expected, result := "aab", nextString; expected != result {
		t.Errorf("Incorrect next string: expected %#v, got %#v", expected, result)
	}
	if expected, result := []day19.RuleList{{day19.RuleLiteral("a"), day19.RuleName("2")}}, nextList; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect next list: expected %v, got %v", expected, result)
	}

	nextString, nextList = gram1.NextSteps("ab", day19.RuleList{day19.RuleName("2")})
	if expected, result := "ab", nextString; expected != result {
		t.Errorf("Incorrect next string: expected %#v, got %#v", expected, result)
	}
	if expected, result := []day19.RuleList{{day19.RuleName("1"), day19.RuleName("3")}, {day19.RuleName("3"), day19.RuleName("1")}}, nextList; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect next list: expected %v, got %v", expected, result)
	}

}

func TestMatches2(t *testing.T) {

	// 	if !day19.ReadRules(`0: "a"`).Matches2("a") {
	// 		t.Error("Incorrect match1: string should match grammar.")
	// 	}

	// 	if !day19.ReadRules(`0: 1
	// 1: "a"`).Matches2("a") {
	// 		t.Error("Incorrect match2: string should match grammar.")
	// 	}

	// 	if !day19.ReadRules(`0: 1
	// 1: 2
	// 2: "a"`).Matches2("a") {
	// 		t.Error("Incorrect match3: string should match grammar.")
	// 	}

	// 	if !day19.ReadRules(`0: 1 2
	// 1: "a"
	// 2: 1 3 | 3 1
	// 3: "b"`).Matches2("aab") {
	// 		t.Error("Incorrect match4: string should match grammar.")
	// 	}

	// 	if !day19.ReadRules(`0: 1 2
	// 1: "a"
	// 2: 1 3 | 3 1
	// 3: "b"`).Matches2("aba") {
	// 		t.Error("Incorrect match5: string should match grammar.")
	// 	}

	// 	if !day19.ReadRules(day19.ExampleInput1).Matches2("ababbb") {
	// 		t.Error("Incorrect match6: string should match grammar.")
	// 	}
	// 	if !day19.ReadRules(day19.ExampleInput1).Matches2("abbbab") {
	// 		t.Error("Incorrect match7: string should match grammar.")
	// 	}
	// 	if day19.ReadRules(day19.ExampleInput1).Matches2("bababa") {
	// 		t.Error("Incorrect match8: string should NOT match grammar.")
	// 	}
	// 	if day19.ReadRules(day19.ExampleInput1).Matches2("aaabbb") {
	// 		t.Error("Incorrect match9: string should NOT match grammar.")
	// 	}
	// 	if day19.ReadRules(day19.ExampleInput1).Matches2("aaaabbb") {
	// 		t.Error("Incorrect match10: string should NOT match grammar.")
	// 	}

	if day19.ReadRules(day19.ExampleInput2R).Matches2("aaaabbaaaabbaaa") {
		t.Error("Incorrect match10: string should NOT match grammar.")
	}

}
