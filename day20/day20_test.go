package day20_test

import (
	"HSteffensen/AoC2020/day20"
	"strings"
	"testing"
)

func TestReverseInt(t *testing.T) {
	if expected, result := 0b1, day20.ReverseInt(1, 1); expected != result {
		t.Errorf("Incorrect reverse int: expected %b, got %b", expected, result)
	}
	if expected, result := 0b10, day20.ReverseInt(0b01, 2); expected != result {
		t.Errorf("Incorrect reverse int: expected %b, got %b", expected, result)
	}
	if expected, result := 0b00110100, day20.ReverseInt(0b00101100, 8); expected != result {
		t.Errorf("Incorrect reverse int: expected %b, got %b", expected, result)
	}
}

func TestNewImageTile(t *testing.T) {
	tile1 := day20.NewImageTile(strings.Split(day20.ExampleInput1, "\n\n")[0])
	if expected, result := 2311, tile1.ID; expected != result {
		t.Errorf("Incorrect new image tile id: expected %v, got %v", expected, result)
	}
	if expected, result := 0b0011010010, tile1.Top(); expected != result {
		t.Errorf("Incorrect new image tile top: expected %v, got %v", expected, result)
	}
	if expected, result := 0b0001011001, tile1.Right(); expected != result {
		t.Errorf("Incorrect new image tile right: expected %v, got %v", expected, result)
	}
	if expected, result := 0b0011100111, tile1.Bottom(); expected != result {
		t.Errorf("Incorrect new image tile bottom: expected %v, got %v", expected, result)
	}
	if expected, result := 0b0111110010, tile1.Left(); expected != result {
		t.Errorf("Incorrect new image tile left: expected %v, got %v", expected, result)
	}
}

func TestImageTileRotate(t *testing.T) {
	tile1 := day20.NewImageTile(strings.Split(day20.ExampleInput1, "\n\n")[0])
	if tile1.At(0, 0) || tile1.At(9, 0) || !tile1.At(9, 9) || tile1.At(0, 9) {
		t.Error("Incorrect new image tile corners: expected bottom right to be the only 'true'")
	}
	if !tile1.At(0, 5) || tile1.At(4, 0) || tile1.At(9, 4) || tile1.At(5, 9) {
		t.Error("Incorrect new image tile sides: expected left to be the only 'true'")
	}

	tile1.RotateAnticlockwise(1)
	if expected, result := 2311, tile1.ID; expected != result {
		t.Errorf("Incorrect new image tile id: expected %v, got %v", expected, result)
	}
	if expected, result := 0b0001011001, tile1.Top(); expected != result {
		t.Errorf("Incorrect new image tile top: expected %v, got %v", expected, result)
	}
	if expected, result := 0b1110011100, tile1.Right(); expected != result {
		t.Errorf("Incorrect new image tile right: expected %v, got %v", expected, result)
	}
	if expected, result := 0b0111110010, tile1.Bottom(); expected != result {
		t.Errorf("Incorrect new image tile bottom: expected %v, got %v", expected, result)
	}
	if expected, result := 0b0100101100, tile1.Left(); expected != result {
		t.Errorf("Incorrect new image tile left: expected %v, got %v", expected, result)
	}
	if tile1.At(0, 0) || !tile1.At(9, 0) || tile1.At(9, 9) || tile1.At(0, 9) {
		t.Error("Incorrect new image tile corners: expected top right to be the only 'true'")
	}
	if tile1.At(0, 5) || tile1.At(4, 0) || tile1.At(9, 4) || !tile1.At(5, 9) {
		t.Error("Incorrect new image tile sides: expected bottom to be the only 'true'")
	}

	tile1.RotateAnticlockwise(1)
	if expected, result := 2311, tile1.ID; expected != result {
		t.Errorf("Incorrect new image tile id: expected %v, got %v", expected, result)
	}
	if expected, result := 0b1110011100, tile1.Top(); expected != result {
		t.Errorf("Incorrect new image tile top: expected %v, got %v", expected, result)
	}
	if expected, result := 0b0100111110, tile1.Right(); expected != result {
		t.Errorf("Incorrect new image tile right: expected %v, got %v", expected, result)
	}
	if expected, result := 0b0100101100, tile1.Bottom(); expected != result {
		t.Errorf("Incorrect new image tile bottom: expected %v, got %v", expected, result)
	}
	if expected, result := 0b1001101000, tile1.Left(); expected != result {
		t.Errorf("Incorrect new image tile left: expected %v, got %v", expected, result)
	}
	if !tile1.At(0, 0) || tile1.At(9, 0) || tile1.At(9, 9) || tile1.At(0, 9) {
		t.Error("Incorrect new image tile corners: expected top left to be the only 'true'")
	}
	if tile1.At(0, 5) || tile1.At(4, 0) || !tile1.At(9, 4) || tile1.At(5, 9) {
		t.Error("Incorrect new image tile sides: expected right to be the only 'true'")
	}
}

func TestMatchingEdges(t *testing.T) {
	tile1 := day20.NewImageTile(`Tile 2311:
	..##.#..#.
	##..#.....
	#...##..#.
	####.#...#
	##.##.###.
	##...#.###
	.#.#.#..##
	..#....#..
	###...#.#.
	..###..###`)
	tile2 := day20.NewImageTile(`Tile 3079:
	#.#.#####.
	.#..######
	..#.......
	######....
	####.#..#.
	.#...#.##.
	#.#####.##
	..#.###...
	..#.......
	..#.###...`)

	if tile1.MatchesEdge(tile2) == -1 {
		t.Errorf("Incorrect matches edge: Tile %v should match an edge with Tile %v", tile1.ID, tile2.ID)
	}

	tiles := []*day20.ImageTile{tile1, tile2}
	tilesMatched := day20.FindNeighbors(tiles)
	if len(tilesMatched[0].Neighbors) != 1 {
		t.Errorf("Incorrect neighbors: Tile %v should have 1 neighbor", tilesMatched[0].ID)
	} else if tilesMatched[0].Neighbors[0].ID != tilesMatched[1].ID {
		t.Errorf("Incorrect neighbors: Tile %v should be neighbors with Tile %v", tilesMatched[0].ID, tilesMatched[1].ID)
	}
	tiles = day20.ReadInput(day20.ExampleInput1)
	tiles = day20.FindNeighbors(tiles)
	if len(tiles[0].Neighbors) != 3 {
		t.Errorf("Incorrect neighbors: Tile %v should have 3 neighbors", tilesMatched[0].ID)
	}
	if len(tiles[1].Neighbors) != 2 {
		t.Errorf("Incorrect neighbors: Tile %v should have 3 neighbors", tilesMatched[0].ID)
	}
}

func TestPart1(t *testing.T) {
	if expected, result := 20899048083289, day20.Part1(day20.ExampleInput1); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}

func TestMakeTopRightUnique(t *testing.T) {
	tiles := day20.ReadInput(day20.ExampleInput1)

	tiles[8].MakeRightEqual(0b1010111110)
	if tiles[8].Right() != 0b1010111110 {
		t.Error("Failed to make tile.right equal desired value.")
	}

	tiles[8].MakeTopRightUnique(tiles)
	if tiles[8].Top() != 0b1010111110 && tiles[8].Right() != 0b0100001000 {
		t.Error("Failed to make tile.top and tile.right equal desired value.")
	}

	tiles = day20.FindNeighbors(tiles)
	if len(tiles[8].Neighbors) != 2 {
		t.Error("Unexpected number of neighbors of [8]")
	}
	tileLeft := tiles[8].NeighborWithEdge(tiles[8].Left())
	if tileLeft.ID != 2311 {
		t.Errorf("Unexpected tile: %v", tileLeft.ID)
	}

	if len(tiles[1].Neighbors) != 2 {
		t.Error("Unexpected number of neighbors of [1]")
	}
	tiles[1].MakeTopRightUnique(tiles)
	tileLeft = tiles[1].NeighborWithEdge(tiles[1].Left())
	if tileLeft.ID != 2311 {
		t.Errorf("Unexpected tile: %v", tileLeft.ID)
	}
}

func TestPuzzleTogether(t *testing.T) {
	tiles := day20.ReadInput(day20.ExampleInput1)
	puzzle := day20.PuzzleTogether(tiles, 3)
	t.Log("Puzzled:")
	t.Log("\n" + day20.StringTileGrid(puzzle))
	image := day20.TileGridToImage(puzzle)
	t.Log("Imaged:")
	t.Log("\n" + day20.StringImage(image))
	t.Error()
}

func TestRotateImage(t *testing.T) {
	image := [][]bool{{true, false, true}, {true, false, false}, {false, false, false}}
	image2 := day20.RotateImage(image)
	if !image2[0][0] || image2[0][1] || image2[1][0] || image2[1][1] {
		t.Log("\n" + day20.StringImage(image2))
		t.Error("Incorrect rotate image")
	}
	image3 := day20.FlipImage(image)
	if image3[0][0] || image3[0][1] || !image3[1][0] || image3[1][1] {
		t.Log("\n" + day20.StringImage(image3))
		t.Error("Incorrect flip image")
	}
}

func TestSerpentAt(t *testing.T) {
	if !day20.HasSeaSerpentAt(day20.SeaSerpent, 0, 0) {
		t.Error("There should be a serpent here")
	}
	if expected, result := 1, day20.CountSeaSerpents(day20.SeaSerpent); expected != result {
		t.Errorf("There should be exactly %v serpent(s) here, counted %v", expected, result)
	}
}

func TestAllImageOrientations(t *testing.T) {
	image := [][]bool{{true, false, true}, {true, false, false}, {false, false, false}}
	for _, image := range day20.AllImageOrientations(image) {
		t.Log("\n" + day20.StringImage(image))
	}
	t.Error()
}

func TestPart2(t *testing.T) {
	if expected, result := 273, day20.Part2(day20.ExampleInput1, 3); expected != result {
		t.Errorf("Incorrect part 2: expected %v, got %v", expected, result)
	}
}
