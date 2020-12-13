package day13

import (
	"strconv"
	"strings"
)

var ExampleInput = `939
7,13,x,x,59,x,31,19`

var Input = `1000508
29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,467,x,x,x,x,x,x,x,23,x,x,x,x,13,x,x,x,17,x,19,x,x,x,x,x,x,x,x,x,x,x,443,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41`

func ReadInput(input string) (int, []string) {
	split := strings.Fields(input)
	time, err := strconv.Atoi(split[0])
	if err != nil {
		panic("Bad input string for day13")
	}
	busses := strings.Split(split[1], ",")
	return time, busses
}
