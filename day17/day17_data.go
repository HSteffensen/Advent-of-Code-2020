package day17

import "strings"

var Input = `#.#.##.#
#.####.#
...##...
#####.##
#....###
##..##..
#..####.
#...#.#.`

var ExampleInput = `.#.
..#
###`

func ReadInput(input string) [][]bool {
	lines := strings.Split(input, "\n")
	result := make([][]bool, len(lines))
	for i, line := range lines {
		result[i] = make([]bool, len(line))
		for j, char := range line {
			result[i][j] = char == '#'
		}
	}
	return result
}
