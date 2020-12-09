package day7_test

import (
	"HSteffensen/AoC2020/day7"
	"reflect"
	"testing"
)

func TestReadInput(t *testing.T) {
	data1 := day7.ReadInput("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	t.Log(data1)
	for bagName, bag := range data1 {
		if bag.Name != bagName {
			t.Errorf("Incorrect bag name in TestReadInput 1: %#v != %#v", bagName, bag.Name)
		}
	}
	if !reflect.DeepEqual(data1["bright white"].HeldBy, []string{"light red"}) {
		t.Errorf("Incorrect HeldBy in TestReadInput 1: %v", data1["bright white"].HeldBy)
	}
}

func TestBagsWhichCanContain(t *testing.T) {
	data1 := day7.ReadInput("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	if len(data1.BagsWhichCanContain("light red")) != 0 {
		t.Error("Incorrect number of bags light red")
	}
	if len(data1.BagsWhichCanContain("bright white")) != 1 {
		t.Error("Incorrect number of bags bright white")
	}
	data2 := day7.ReadInput(day7.ExampleInput)
	if len(data2.BagsWhichCanContain("shiny gold")) != 4 {
		t.Error("Incorrect number of bags shiny gold")
	}
}

func TestBagsWithin(t *testing.T) {
	data1 := day7.ReadInput("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	if val, expected := data1.BagsWithin("light red"), 3; val != expected {
		t.Errorf("Incorrect number of bags shiny gold: expected %v, got %v", expected, val)
	}
	if val, expected := data1.BagsWithin("bright white"), 0; val != expected {
		t.Errorf("Incorrect number of bags shiny gold: expected %v, got %v", expected, val)
	}
	data2 := day7.ReadInput(day7.ExampleInput)
	if val, expected := data2.BagsWithin("shiny gold"), 32; val != expected {
		t.Errorf("Incorrect number of bags shiny gold: expected %v, got %v", expected, val)
	}
}
