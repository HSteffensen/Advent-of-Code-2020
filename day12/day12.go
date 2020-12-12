package day12

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Ship struct {
	X      int
	Y      int
	Facing int
}

func (s *Ship) Move(direction, distance int) {
	dx, dy := 1, 0
	for i := 0; i < direction; i++ {
		dx, dy = -dy, dx
	}
	s.X += dx * distance
	s.Y += dy * distance
}

func (s *Ship) TravelStep(step string) {
	direction := step[:1]
	distance, err := strconv.Atoi(step[1:])
	if err != nil {
		panic(fmt.Sprintf("Bad Ship step distance: %#v", step[1:]))
	}
	switch direction {
	case "N", "S", "E", "W":
		directions := map[string]int{"E": 0, "N": 1, "W": 2, "S": 3}
		s.Move(directions[direction], distance)
	case "R":
		distance = -distance
		fallthrough
	case "L":
		leftTurns := distance / 90
		s.Facing += leftTurns
		if s.Facing < 0 {
			s.Facing += 4
		}
	case "F":
		s.Move(s.Facing, distance)
	}
}

func (s *Ship) TravelPath(steps []string) {
	for _, step := range steps {
		s.TravelStep(step)
	}
}

func Part1(input string) int {
	ship1 := Ship{}
	steps := strings.Fields(input)
	ship1.TravelPath(steps)
	return int(math.Abs(float64(ship1.X)) + math.Abs(float64(ship1.Y)))
}
