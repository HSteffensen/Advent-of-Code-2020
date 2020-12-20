package aocCommon

import (
	"fmt"
	"strconv"
	"strings"
)

func ReadInputToInts(data string) []int {
	split := strings.Fields(data)
	if len(split) == 1 {
		split = strings.Split(data, ",")
	}
	var toInts []int
	for _, v := range split {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error in input string to integer list conversion:", v, data)
		}
		toInts = append(toInts, vInt)
	}
	return toInts
}
