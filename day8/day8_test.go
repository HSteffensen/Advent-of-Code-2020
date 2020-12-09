package day8_test

import (
	"HSteffensen/AoC2020/day8"
	"testing"
)

func TestReadInput(t *testing.T) {
	data1 := day8.ReadInput("acc +1")
	if expected := (day8.Instruction{"acc", 1}); data1[0] != expected {
		t.Errorf("Incorrect Instruction: expected %v, got %v", expected, data1[0])
	}
	data2 := day8.ReadInput("jmp -1")
	if expected := (day8.Instruction{"jmp", -1}); data2[0] != expected {
		t.Errorf("Incorrect Instruction: expected %v, got %v", expected, data2[0])
	}
	data3 := day8.ReadInput(day8.ExampleInput)
	if expected := (day8.Instruction{"nop", 0}); data3[0] != expected {
		t.Errorf("Incorrect Instruction: expected %v, got %v", expected, data3[0])
	}
}

func TestFindLoopAcc(t *testing.T) {
	data1 := day8.ReadInput(day8.ExampleInput)
	if actual, expected := data1.FindLoopAcc(), 5; actual != expected {
		t.Errorf("Incorrect acc value: expected %v, got %v", expected, actual)
	}
}

func TestFixLoopAcc(t *testing.T) {
	data1 := day8.ReadInput("jmp +0")
	if actual, expected := data1.FixLoopAcc(), 0; actual != expected {
		t.Errorf("Incorrect acc value: expected %v, got %v", expected, actual)
	}
	data2 := day8.ReadInput(day8.ExampleInput)
	if actual, expected := data2.FixLoopAcc(), 8; actual != expected {
		t.Errorf("Incorrect acc value: expected %v, got %v", expected, actual)
	}
}
