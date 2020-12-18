package day15_test

import (
	"HSteffensen/AoC2020/day15"
	"reflect"
	"testing"
)

func TestGameHistoryReset(t *testing.T) {
	gh := day15.GameHistory{}
	input := []int{0, 3, 6}
	gh.Reset(input)
	if expected, result := []uint32{1, 0, 0, 2, 0, 0, 0, 0}, gh.MostRecentTime[:8]; !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect recently seen: expected %v, got %v", expected, result)
	}
	if expected, result := input[len(input)-1], int(gh.NextNumber); result != expected {
		t.Errorf("Incorrect next number: expected %v, got %v", expected, result)
	}
	if expected, result := 3, int(gh.Time); expected != result {
		t.Errorf("Incorrect time: expected %v, got %v", expected, result)
	}
}

// func TestGameHistoryStep(t *testing.T) {
// 	gh := day15.GameHistory{History: []int{0, 3}, NextNumber: 6, MostRecentTime: map[int]int{0: 0, 3: 1}}
// 	t.Log("before step 1:", gh)
// 	gh.Step()
// 	t.Log("after step 1:", gh)
// 	if expected, result := 0, gh.NextNumber; expected != result {
// 		t.Errorf("Incorrect step 1 history: expected %v, got %v", expected, result)
// 	}
// 	if expected, result := 0, gh.MostRecentTime[0]; expected != result {
// 		t.Errorf("Incorrect step 1 map[0]: expected %v, got %v", expected, result)
// 	}

// 	gh.Step()
// 	t.Log("after step 2:", gh)
// 	if expected, result := 3, gh.NextNumber; expected != result {
// 		t.Errorf("Incorrect step 2 history: expected %v, got %v", expected, result)
// 	}
// 	if expected, result := 3, gh.MostRecentTime[0]; expected != result {
// 		t.Errorf("Incorrect step 2 map[0]: expected %v, got %v", expected, result)
// 	}

// 	gh.Step()
// 	t.Log("after step 3:", gh)
// 	gh.Step()
// 	t.Log("after step 4:", gh)
// 	gh.Step()
// 	t.Log("after step 5:", gh)
// 	gh.Step()
// 	t.Log("after step 6:", gh)
// 	gh.Step()
// 	t.Log("after step 7:", gh)
// 	if expected, result := []int{0, 3, 6, 0, 3, 3, 1, 0, 4, 0}, gh.FullHistory(); !reflect.DeepEqual(expected, result) {
// 		t.Errorf("Incorrect step full history: expected %v, got %v", expected, result)
// 	}
// 	if expected, result := map[int]int{0: 7, 3: 5, 6: 2, 1: 6, 4: 8}, gh.MostRecentTime; !reflect.DeepEqual(expected, result) {
// 		t.Errorf("Incorrect step full history: expected %v, got %v", expected, result)
// 	}
// }

func TestRunUntil(t *testing.T) {
	if expected, result := 436, day15.RunUntil([]int{0, 3, 6}, 2020); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 1, day15.RunUntil([]int{1, 3, 2}, 2020); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 10, day15.RunUntil([]int{2, 1, 3}, 2020); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 27, day15.RunUntil([]int{1, 2, 3}, 2020); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 78, day15.RunUntil([]int{2, 3, 1}, 2020); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 438, day15.RunUntil([]int{3, 2, 1}, 2020); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 1836, day15.RunUntil([]int{3, 1, 2}, 2020); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
}

func TestRunUntil2(t *testing.T) {
	if expected, result := 175594, day15.RunUntil([]int{0, 3, 6}, 30000000); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 2578, day15.RunUntil([]int{1, 3, 2}, 30000000); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 3544142, day15.RunUntil([]int{2, 1, 3}, 30000000); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 261214, day15.RunUntil([]int{1, 2, 3}, 30000000); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 6895259, day15.RunUntil([]int{2, 3, 1}, 30000000); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 18, day15.RunUntil([]int{3, 2, 1}, 30000000); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
	if expected, result := 362, day15.RunUntil([]int{3, 1, 2}, 30000000); expected != result {
		t.Errorf("Incorrect Part1: expected %v, got %v", expected, result)
	}
}

func BenchmarkPart2(b *testing.B) {
	day15.RunUntil([]int{0, 14, 1, 3, 7, 9}, 30000000)
}
