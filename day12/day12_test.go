package day12_test

import (
	"HSteffensen/AoC2020/day12"
	"testing"
)

func TestShipMove(t *testing.T) {
	ship1 := day12.Ship{}
	ship1.Move(0, 1)
	if eX, eY := 1, 0; ship1.X != eX && ship1.Y != eY {
		t.Errorf("Incorrect ship position: expected (%v, %v), got (%v, %v)", eX, eY, ship1.X, ship1.Y)
	}

	ship2 := day12.Ship{}
	ship2.Move(1, 1)
	if eX, eY := 0, 1; ship2.X != eX && ship2.Y != eY {
		t.Errorf("Incorrect ship position: expected (%v, %v), got (%v, %v)", eX, eY, ship2.X, ship2.Y)
	}

	ship3 := day12.Ship{}
	ship3.Move(2, 1)
	if eX, eY := -1, 0; ship3.X != eX && ship3.Y != eY {
		t.Errorf("Incorrect ship position: expected (%v, %v), got (%v, %v)", eX, eY, ship3.X, ship3.Y)
	}

	ship4 := day12.Ship{}
	ship4.Move(3, 1)
	if eX, eY := 0, -1; ship4.X != eX && ship4.Y != eY {
		t.Errorf("Incorrect ship position: expected (%v, %v), got (%v, %v)", eX, eY, ship4.X, ship4.Y)
	}

	ship1.Move(0, 1)
	if eX, eY := 2, 0; ship1.X != eX && ship1.Y != eY {
		t.Errorf("Incorrect ship position: expected (%v, %v), got (%v, %v)", eX, eY, ship1.X, ship1.Y)
	}
	ship1.Move(1, 1)
	if eX, eY := 2, 1; ship1.X != eX && ship1.Y != eY {
		t.Errorf("Incorrect ship position: expected (%v, %v), got (%v, %v)", eX, eY, ship1.X, ship1.Y)
	}
	ship1.Move(2, 1)
	if eX, eY := 1, 1; ship1.X != eX && ship1.Y != eY {
		t.Errorf("Incorrect ship position: expected (%v, %v), got (%v, %v)", eX, eY, ship1.X, ship1.Y)
	}
	ship1.Move(3, 1)
	if eX, eY := 1, 0; ship1.X != eX && ship1.Y != eY {
		t.Errorf("Incorrect ship position: expected (%v, %v), got (%v, %v)", eX, eY, ship1.X, ship1.Y)
	}
	ship1.Move(2, 2)
	if eX, eY := -1, 0; ship1.X != eX && ship1.Y != eY {
		t.Errorf("Incorrect ship position: expected (%v, %v), got (%v, %v)", eX, eY, ship1.X, ship1.Y)
	}
}

func TestPart1(t *testing.T) {
	if expected, result := 25, day12.Part1(day12.ExampleInput); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}
