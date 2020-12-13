package day13_test

import (
	"HSteffensen/AoC2020/day13"
	"testing"
)

func TestTimeUntilBus(t *testing.T) {
	if expected, result := 5, day13.TimeUntilBus(939, 59); expected != result {
		t.Errorf("Incorrect time result: expected %v, got %v", expected, result)
	}
	if expected, result := 6, day13.TimeUntilBus(939, 7); expected != result {
		t.Errorf("Incorrect time result: expected %v, got %v", expected, result)
	}
	if expected, result := 10, day13.TimeUntilBus(939, 13); expected != result {
		t.Errorf("Incorrect time result: expected %v, got %v", expected, result)
	}
}

func TestPart1(t *testing.T) {
	if expected, result := 295, day13.Part1(day13.ExampleInput); expected != result {
		t.Errorf("Incorrect Part1 result: expected %v, got %v", expected, result)
	}
}

func TestFindDeparture(t *testing.T) {
	if expected, result := 77, day13.FindDeparture(0, 7, 13, 1); expected != result {
		t.Errorf("Incorrect Part1 result: expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	testInput := `1
7,13`
	if expected, result := 77, day13.Part2(testInput); expected != result {
		t.Errorf("Incorrect Part1 result: expected %v, got %v", expected, result)
	}
	testInput = `1
17,x,13,19`
	if expected, result := 3417, day13.Part2(testInput); expected != result {
		t.Errorf("Incorrect Part1 result: expected %v, got %v", expected, result)
	}
	testInput = `1
67,7,59,61`
	if expected, result := 754018, day13.Part2(testInput); expected != result {
		t.Errorf("Incorrect Part1 result: expected %v, got %v", expected, result)
	}
	testInput = `1
67,x,7,59,61`
	if expected, result := 779210, day13.Part2(testInput); expected != result {
		t.Errorf("Incorrect Part1 result: expected %v, got %v", expected, result)
	}
	testInput = `1
67,7,x,59,61`
	if expected, result := 1261476, day13.Part2(testInput); expected != result {
		t.Errorf("Incorrect Part1 result: expected %v, got %v", expected, result)
	}
	testInput = `1
1789,37,47,1889`
	if expected, result := 1202161486, day13.Part2(testInput); expected != result {
		t.Errorf("Incorrect Part1 result: expected %v, got %v", expected, result)
	}
}
