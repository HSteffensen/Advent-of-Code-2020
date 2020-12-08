package day3_test

import (
	"HSteffensen/AoC2020/day3"
	"reflect"
	"testing"
)

func TestReadTerrainGrid(t *testing.T) {
	exampleTerrain := day3.ReadTerrainGrid(day3.ExampleInput)
	expectedLine1 := []day3.TerrainFeature{day3.Open, day3.Open, day3.Tree, day3.Tree, day3.Open, day3.Open, day3.Open, day3.Open, day3.Open, day3.Open, day3.Open}
	expectedLine2 := []day3.TerrainFeature{day3.Tree, day3.Open, day3.Open, day3.Open, day3.Tree, day3.Open, day3.Open, day3.Open, day3.Tree, day3.Open, day3.Open}
	if !reflect.DeepEqual(exampleTerrain[0], expectedLine1) {
		t.Errorf("Incorrect terrain grid line 1: expected %v, got %v", expectedLine1, exampleTerrain[0])
	}
	if !reflect.DeepEqual(exampleTerrain[1], expectedLine2) {
		t.Errorf("Incorrect terrain grid line 1: expected %v, got %v", expectedLine2, exampleTerrain[1])
	}
}

func TestTreesOnAngle(t *testing.T) {
	exampleTerrain := day3.ReadTerrainGrid(day3.ExampleInput)
	result := exampleTerrain.TreesOnAngle(3, 1)
	expected := 7
	if result != expected {
		t.Errorf("Incorrect Tree count: expected %v, got %v", expected, result)
	}
}
