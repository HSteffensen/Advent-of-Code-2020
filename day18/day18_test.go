package day18_test

import (
	"HSteffensen/AoC2020/day18"
	"testing"
)

func TestParseMath(t *testing.T) {
	p1 := map[string]int{}
	if expected, result := "1", day18.StringReversePolishNotation(day18.ParseMath("1", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := "1 2 +", day18.StringReversePolishNotation(day18.ParseMath("1 + 2", p1)); expected != result {
		t.Errorf("Incorrect ParseMath: expected %#v, got %#v", expected, result)
	}
	if expected, result := "1 2 + 3 *", day18.StringReversePolishNotation(day18.ParseMath("1 + 2 * 3", p1)); expected != result {
		t.Errorf("Incorrect ParseMath: expected %#v, got %#v", expected, result)
	}
	if expected, result := "1 2 + 3 *", day18.StringReversePolishNotation(day18.ParseMath("(1 + 2) * 3", p1)); expected != result {
		t.Errorf("Incorrect ParseMath: expected %#v, got %#v", expected, result)
	}
	if expected, result := "1 2 3 * +", day18.StringReversePolishNotation(day18.ParseMath("1 + (2 * 3)", p1)); expected != result {
		t.Errorf("Incorrect ParseMath: expected %#v, got %#v", expected, result)
	}
	if expected, result := "1 2 3 * + 4 +", day18.StringReversePolishNotation(day18.ParseMath("1 + (2 * 3) + 4", p1)); expected != result {
		t.Errorf("Incorrect ParseMath: expected %#v, got %#v", expected, result)
	}
	if expected, result := "1 2 3 * 4 + + 5 *", day18.StringReversePolishNotation(day18.ParseMath("1 + ((2 * 3) + 4) * 5", p1)); expected != result {
		t.Errorf("Incorrect ParseMath: expected %#v, got %#v", expected, result)
	}

	if expected, result := "1 2 + 3 * 4 + 5 * 6 +", day18.StringReversePolishNotation(day18.ParseMath("1 + 2 * 3 + 4 * 5 + 6", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := "1 2 3 * + 4 5 6 + * +", day18.StringReversePolishNotation(day18.ParseMath("1 + (2 * 3) + (4 * (5 + 6))", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := "2 3 * 4 5 * +", day18.StringReversePolishNotation(day18.ParseMath("2 * 3 + (4 * 5)", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := "5 8 3 * 9 + 3 + 4 * 3 * +", day18.StringReversePolishNotation(day18.ParseMath("5 + (8 * 3 + 9 + 3 * 4 * 3)", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := "5 9 * 7 3 * 3 * 9 + 3 * 8 6 + 4 * + *", day18.StringReversePolishNotation(day18.ParseMath("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := "2 4 + 9 * 6 9 + 8 * 6 + * 6 + 2 + 4 + 2 *", day18.StringReversePolishNotation(day18.ParseMath("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", p1)); expected != result {
		t.Errorf("Incorrect ParseMath: expected %#v, got %#v", expected, result)
	}
}

func TestEvaluate(t *testing.T) {
	p1 := map[string]int{}
	if expected, result := 1, day18.EvaluateReversePolishNotation(day18.ParseMath("1", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 3, day18.EvaluateReversePolishNotation(day18.ParseMath("1 + 2", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 71, day18.EvaluateReversePolishNotation(day18.ParseMath("1 + 2 * 3 + 4 * 5 + 6", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 51, day18.EvaluateReversePolishNotation(day18.ParseMath("1 + (2 * 3) + (4 * (5 + 6))", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 26, day18.EvaluateReversePolishNotation(day18.ParseMath("2 * 3 + (4 * 5)", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 437, day18.EvaluateReversePolishNotation(day18.ParseMath("5 + (8 * 3 + 9 + 3 * 4 * 3)", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 12240, day18.EvaluateReversePolishNotation(day18.ParseMath("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 13632, day18.EvaluateReversePolishNotation(day18.ParseMath("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", p1)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}

	p2 := map[string]int{"+": 1}
	if expected, result := 1, day18.EvaluateReversePolishNotation(day18.ParseMath("1", p2)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 3, day18.EvaluateReversePolishNotation(day18.ParseMath("1 + 2", p2)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 231, day18.EvaluateReversePolishNotation(day18.ParseMath("1 + 2 * 3 + 4 * 5 + 6", p2)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 51, day18.EvaluateReversePolishNotation(day18.ParseMath("1 + (2 * 3) + (4 * (5 + 6))", p2)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 46, day18.EvaluateReversePolishNotation(day18.ParseMath("2 * 3 + (4 * 5)", p2)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 1445, day18.EvaluateReversePolishNotation(day18.ParseMath("5 + (8 * 3 + 9 + 3 * 4 * 3)", p2)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 669060, day18.EvaluateReversePolishNotation(day18.ParseMath("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", p2)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
	if expected, result := 23340, day18.EvaluateReversePolishNotation(day18.ParseMath("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", p2)); expected != result {
		t.Errorf("Incorrect Evaluate: expected %v, got %v", expected, result)
	}
}

func TestPart1(t *testing.T) {
	if expected, result := 71+51+26+437+12240+13632, day18.Part1(day18.ExampleInput1); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	if expected, result := 231+51+46+1445+669060+23340, day18.Part2(day18.ExampleInput1); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
}
