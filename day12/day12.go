package day12

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Ship interface {
	Location() Point
	MoveCardinal(string, int)
	MoveForward(int)
	Rotate(int)
}

type Point struct {
	X int
	Y int
}

type CardinalShip struct {
	location Point
	Facing   int
}

type WaypointShip struct {
	location Point
	Waypoint Point
}

func (p Point) Add(other Point) Point {
	p.X += other.X
	p.Y += other.Y
	return p
}

func (p Point) Subtract(other Point) Point {
	p.X -= other.X
	p.Y -= other.Y
	return p
}

func (p Point) Multiply(magnitude int) Point {
	p.X *= magnitude
	p.Y *= magnitude
	return p
}

func (p Point) Magnitude() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) Rotate(times int) Point {
	times %= 4
	if times < 0 {
		times += 4
	}
	for i := 0; i < times; i++ {
		p.X, p.Y = -p.Y, p.X
	}
	return p
}

func (s *CardinalShip) Location() Point {
	return s.location
}

func (s *WaypointShip) Location() Point {
	return s.location
}

func (s *CardinalShip) MoveDirection(direction, distance int) {
	s.location = s.location.Add(Point{1, 0}.Rotate(direction).Multiply(distance))
}

func (s *CardinalShip) MoveCardinal(direction string, distance int) {
	directions := map[string]int{"E": 0, "N": 1, "W": 2, "S": 3}
	if directionInt, ok := directions[direction]; ok {
		s.MoveDirection(directionInt, distance)
	} else {
		panic(fmt.Sprintf("Bad cardinal direction: %#v", direction))
	}
}

func (s *WaypointShip) MoveCardinal(direction string, distance int) {
	directions := map[string]int{"E": 0, "N": 1, "W": 2, "S": 3}
	if directionInt, ok := directions[direction]; ok {
		s.Waypoint = s.Waypoint.Add(Point{1, 0}.Rotate(directionInt).Multiply(distance))
	} else {
		panic(fmt.Sprintf("Bad cardinal direction: %#v", direction))
	}
}

func (s *CardinalShip) MoveForward(distance int) {
	s.MoveDirection(s.Facing, distance)
}

func (s *WaypointShip) MoveForward(distance int) {
	s.location = s.location.Add(s.Waypoint.Multiply(distance))
}

func (s *CardinalShip) Rotate(times int) {
	s.Facing += times
}

func (s *WaypointShip) Rotate(times int) {
	s.Waypoint = s.Waypoint.Rotate(times)
}

func TravelStep(s Ship, step string) {
	direction := step[:1]
	distance, err := strconv.Atoi(step[1:])
	if err != nil {
		panic(fmt.Sprintf("Bad Ship step distance: %#v", step[1:]))
	}
	switch direction {
	case "R":
		distance = -distance
		fallthrough
	case "L":
		leftTurns := distance / 90
		s.Rotate(leftTurns)
	case "F":
		s.MoveForward(distance)
	default:
		s.MoveCardinal(direction, distance)
	}
}

func TravelPath(s Ship, steps []string) {
	for _, step := range steps {
		TravelStep(s, step)
	}
}

func Part1(input string) int {
	ship1 := CardinalShip{}
	steps := strings.Fields(input)
	TravelPath(&ship1, steps)
	return ship1.Location().Magnitude()
}

func Part2(input string) int {
	ship1 := WaypointShip{Waypoint: Point{10, 1}}
	steps := strings.Fields(input)
	TravelPath(&ship1, steps)
	return ship1.Location().Magnitude()
}
