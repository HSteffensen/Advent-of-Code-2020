package day18

import (
	"fmt"
	"strconv"
	"strings"
)

type MathToken interface {
	Evaluate([]int) []int
}
type Value int
type Operator2 func(int, int) int

func plusOperator(a, b int) int {
	return a + b
}

func multiplyOperator(a, b int) int {
	return a * b
}

var PlusOperator = Operator2(plusOperator)
var MultiplyOperator = Operator2(multiplyOperator)

func OperatorFrom(op string) Operator2 {
	switch op {
	case "+":
		return PlusOperator
	case "*":
		return MultiplyOperator
	}
	panic(fmt.Sprintf("Bad Operator: %v", op))
}

func (v Value) Evaluate(stack []int) []int {
	return append(stack, int(v))
}

func (op Operator2) Evaluate(stack []int) []int {
	args := stack[len(stack)-2:]
	stack = stack[:len(stack)-2]
	return append(stack, op(args[1], args[0]))
}

func EvaluateReversePolishNotation(equation []MathToken) int {
	stack := make([]int, 0, len(equation))

	for _, token := range equation {
		stack = token.Evaluate(stack)
	}

	if len(stack) == 1 {
		return stack[0]
	}
	panic("Bad equation: ended with too many numbers in calculation stack.")
}

func StringReversePolishNotation(equation []MathToken) string {
	result := ""

	for _, token := range equation {
		switch v := token.(type) {
		case Value:
			result += fmt.Sprintf(" %v", int(v))
		case Operator2:
			switch v(1, 1) {
			case 2:
				result += " +"
			case 1:
				result += " *"
			default:
				result += " ?ERROR?"
			}
		}
	}

	return strings.Trim(result, " ")
}

func ParseMath(input string, precedences map[string]int) []MathToken {
	input = strings.ReplaceAll(strings.ReplaceAll(input, ")", " )"), "(", "( ")
	tokens := strings.Fields(input)

	equation := make([]MathToken, 0, len(input))
	opStack := make([]string, 0, len(input))
	for _, token := range tokens {
		switch token {
		case "+", "*":
			for end := len(opStack) - 1; len(opStack) > 0 && opStack[end] != "(" && precedences[opStack[end]] >= precedences[token]; end = len(opStack) - 1 {
				equation = append(equation, OperatorFrom(opStack[end]))
				opStack = opStack[:end]
			}
			fallthrough
		case "(":
			opStack = append(opStack, token)
		case ")":
			for end := len(opStack) - 1; len(opStack) > 0 && opStack[end] != "("; end = len(opStack) - 1 {
				equation = append(equation, OperatorFrom(opStack[end]))
				opStack = opStack[:end]
			}
			if end := len(opStack) - 1; opStack[end] == "(" {
				opStack = opStack[:end]
			}
		default:
			val, err := strconv.Atoi(token)
			if err != nil {
				panic(fmt.Sprintf("Bad math token: expected an integer, got %#v.", token))
			}
			equation = append(equation, Value(val))
		}
	}
	for len(opStack) > 0 {
		end := len(opStack) - 1
		equation = append(equation, OperatorFrom(opStack[end]))
		opStack = opStack[:end]
	}
	return equation
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		total += EvaluateReversePolishNotation(ParseMath(line, map[string]int{}))
	}
	return total
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	precidence := map[string]int{"+": 1}
	for _, line := range lines {
		total += EvaluateReversePolishNotation(ParseMath(line, precidence))
	}
	return total
}
