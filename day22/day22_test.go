package day22_test

import (
	"HSteffensen/AoC2020/day22"
	"reflect"
	"testing"
)

func TestReadInput(t *testing.T) {
	if expected, result := (day22.CombatGame{{1}, {2}}), day22.ReadInput(`Player 1:
1

Player 2:
2`); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect read input: expected %v, got %v", expected, result)
	}
	if expected, result := (day22.CombatGame{{9, 2, 6, 3, 1}, {5, 8, 4, 7, 10}}), day22.ReadInput(day22.ExampleInput1); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect read input: expected %v, got %v", expected, result)
	}
}

func TestDeckPlayTopCard(t *testing.T) {
	deck1 := day22.Deck{1, 2, 3}
	if expected, result := 1, deck1.PlayTopCard(); expected != result {
		t.Errorf("Incorrect play top card: expected %v, got %v", expected, result)
	}
	if expected, result := (day22.Deck{2, 3}), deck1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect play top card deck: expected %v, got %v", expected, result)
	}

}

func TestDeckPutOnBottom(t *testing.T) {
	deck1 := day22.Deck{1, 2, 3}
	deck1.PutOnBottom(4, 5)
	if expected, result := (day22.Deck{1, 2, 3, 5, 4}), deck1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect play top card deck: expected %v, got %v", expected, result)
	}
	deck1.PutOnBottom(7, 6)
	if expected, result := (day22.Deck{1, 2, 3, 5, 4, 7, 6}), deck1; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect play top card deck: expected %v, got %v", expected, result)
	}
}

func TestCombatGameWinner(t *testing.T) {
	if expected, result := 0, (day22.CombatGame{{1}, {}}).Winner(); expected != result {
		t.Errorf("Incorrect game winner: expected %v, got %v", expected, result)
	}
	if expected, result := 1, (day22.CombatGame{{}, {1}}).Winner(); expected != result {
		t.Errorf("Incorrect game winner: expected %v, got %v", expected, result)
	}
	if expected, result := -1, (day22.CombatGame{{1}, {1}}).Winner(); expected != result {
		t.Errorf("Incorrect game winner: expected %v, got %v", expected, result)
	}
}

func TestPart1(t *testing.T) {
	if expected, result := 306, day22.Part1(day22.ExampleInput1); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}

func TestPlayRecursiveGame(t *testing.T) {
	game1 := day22.ReadInput(`Player 1:
43
19

Player 2:
2
29
14`)
	gc1 := 0
	if expected, result := 0, game1.PlayRecursiveGame(&gc1); expected != result {
		t.Errorf("Incorrect recursive game: expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	if expected, result := 291, day22.Part2(day22.ExampleInput1); expected != result {
		t.Errorf("Incorrect part 2: expected %v, got %v", expected, result)
	}
}

func BenchmarkPart2(t *testing.B) {
	day22.Part2(day22.Input)
}
