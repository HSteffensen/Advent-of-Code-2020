package day20

import (
	"fmt"
	"strconv"
	"strings"
)

type ImageTile struct {
	ID                int
	grid              [10][10]bool
	rotation          int
	flippedVertically bool
	edgeNumbers       [4]int // top, right, bottom, left
	Neighbors         []*ImageTile
}

func (tile ImageTile) String() string {
	result := make([]string, 0, 1000)
	result = append(result, fmt.Sprintf("\nTile %v:\n", tile.ID))
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			item := "."
			if tile.At(x, y) {
				item = "#"
			}
			result = append(result, item)
		}
		result = append(result, "\n")
	}

	return strings.Join(result, "")
}

func ReverseInt(value, bits int) int {
	result := 0
	for i := 0; i < bits; i++ {
		result <<= 1
		j := 1 << i
		if value&j != 0 {
			result++
		}
	}
	return result
}

func NewImageTile(input string) *ImageTile {
	lines := strings.Split(input, "\n")
	line1 := lines[0]
	lines = lines[1:]

	idString := line1[5 : len(line1)-1]
	idNumber, err := strconv.Atoi(idString)
	if err != nil {
		panic(fmt.Sprintf("Bad NewImageTile input string: expected \"Tile ####:\" as first line, got %#v.", line1))
	}
	grid := [10][10]bool{}
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = lines[i][j] == '#'
		}
	}
	edges := [4]int{}
	//top
	for x, y := 0, 0; x < 10; x++ {
		edges[0] <<= 1
		if grid[y][x] {
			edges[0]++
		}
	}
	//right
	for x, y := 9, 0; y < 10; y++ {
		edges[1] <<= 1
		if grid[y][x] {
			edges[1]++
		}
	}
	//bottom
	for x, y := 9, 9; x >= 0; x-- {
		edges[2] <<= 1
		if grid[y][x] {
			edges[2]++
		}
	}
	//left
	for x, y := 0, 9; y >= 0; y-- {
		edges[3] <<= 1
		if grid[y][x] {
			edges[3]++
		}
	}
	return &ImageTile{ID: idNumber, grid: grid, edgeNumbers: edges, Neighbors: make([]*ImageTile, 0, 4)}
}

func (tile ImageTile) Top() int {
	if tile.flippedVertically {
		return ReverseInt(tile.edgeNumbers[(2+tile.rotation)%4], 10)
	}
	return tile.edgeNumbers[(0+tile.rotation)%4]
}

func (tile ImageTile) Right() int {
	if tile.flippedVertically {
		return ReverseInt(tile.edgeNumbers[(1+tile.rotation)%4], 10)
	}
	return tile.edgeNumbers[(1+tile.rotation)%4]
}

func (tile ImageTile) Bottom() int {
	if tile.flippedVertically {
		return tile.edgeNumbers[(0+tile.rotation)%4]
	}
	return ReverseInt(tile.edgeNumbers[(2+tile.rotation)%4], 10)
}

func (tile ImageTile) Left() int {
	if tile.flippedVertically {
		return tile.edgeNumbers[(3+tile.rotation)%4]
	}
	return ReverseInt(tile.edgeNumbers[(3+tile.rotation)%4], 10)
}

func (tile ImageTile) At(x, y int) bool {
	if tile.flippedVertically {
		y = 9 - y
	}
	for i := 0; i < tile.rotation; i++ {
		x, y = 9-y, x
	}
	return tile.grid[y][x]
}

func (tile *ImageTile) RotateAnticlockwise(times int) {
	if tile.flippedVertically {
		times = -times
	}
	tile.rotation = (tile.rotation + times) % 4
	if tile.rotation < 0 {
		tile.rotation += 4
	}
}

func (tile *ImageTile) FlipVertically() {
	tile.flippedVertically = !tile.flippedVertically
}

func (tile *ImageTile) MakeTopEqual(target int) {
	for i, edge := range tile.edgeNumbers {
		if edge == target {
			tile.flippedVertically = false
			tile.rotation = i
			return
		} else if ReverseInt(edge, 10) == target {
			tile.flippedVertically = true
			tile.rotation = (i + 2) % 4
			return
		}
	}
	panic(fmt.Sprintf("Bad tile flop target: Tile %v has no edge to match %b", tile.ID, target))
}

func (tile *ImageTile) MakeRightEqual(target int) {
	tile.MakeTopEqual(target)
	tile.RotateAnticlockwise(-1)
}

func (tile *ImageTile) MakeTopRightUnique(tileList []*ImageTile) {
	counts := CountSideUniqueness(tileList)
	for _, edge := range tile.edgeNumbers {
		if counts[edge] == 1 {
			tile.MakeTopEqual(edge)
			break
		}
	}
	if counts[tile.Right()] != 1 {
		tile.RotateAnticlockwise(-1)
	}

	if counts[tile.Right()] != 1 || counts[tile.Top()] != 1 {
		panic("Failed to make tile top right unique")
	}
}

func (tile ImageTile) MatchesEdge(other *ImageTile) int {
	for i, edge1 := range tile.edgeNumbers {
		edge1Reverse := ReverseInt(edge1, 10)
		for _, edge2 := range other.edgeNumbers {
			if edge2 == edge1 || edge2 == edge1Reverse {
				return i
			}
		}
	}
	return -1
}

func (tile ImageTile) NeighborWithEdge(target int) *ImageTile {
	targetReverse := ReverseInt(target, 10)
	for _, neighbor := range tile.Neighbors {
		for _, edge := range neighbor.edgeNumbers {
			if edge == target || edge == targetReverse {
				return neighbor
			}
		}
	}
	panic("Failed to find neighbor with the given edge")
}

func CountSideUniqueness(tiles []*ImageTile) map[int]int {
	counts := make(map[int]int)

	for _, tile := range tiles {
		for _, edge := range tile.edgeNumbers {
			counts[edge]++
			counts[ReverseInt(edge, 10)]++
		}
	}

	return counts
}

func CountCounts(counts map[int]int) map[int]int {
	tally := make(map[int]int)

	for _, count := range counts {
		tally[count]++
	}

	return tally
}

func FindNeighbors(tiles []*ImageTile) []*ImageTile {
	for i := range tiles {
		tiles[i].Neighbors = make([]*ImageTile, 0, 4)
	}
	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			if tiles[i].MatchesEdge(tiles[j]) >= 0 {
				tiles[i].Neighbors = append(tiles[i].Neighbors, tiles[j])
				tiles[j].Neighbors = append(tiles[j].Neighbors, tiles[i])
			}
		}
	}
	return tiles
}

func Part1(input string) int {
	tiles := ReadInput(input)
	tiles = FindNeighbors(tiles)
	result := 1
	corners := make([]*ImageTile, 0, 4)
	for _, tile := range tiles {
		if len(tile.Neighbors) == 2 {
			result *= tile.ID
			corners = append(corners, tile)
		}
	}
	if count := len(corners); count != 4 {
		panic(fmt.Sprintf("Wrong number of corner tiles: should be 4 but is %v", count))
	}
	return result
}

func StringTileRow(tiles []*ImageTile) string {
	resultRows := make([][]string, 16)
	for _, tile := range tiles {
		tileRows := strings.Split(strings.Trim(tile.String(), "\n"), "\n")
		for i, row := range tileRows {
			resultRows[i] = append(resultRows[i], row)
		}
	}

	result := make([]string, 16)
	for i, row := range resultRows {
		if len(row) > 0 {
			result[i] = strings.Join(row, " ") + "\n"
		}
	}

	return strings.Join(result, "")
}

func StringTileGrid(tileGrid [][]*ImageTile) string {

	lines := make([]string, len(tileGrid))
	for i, line := range tileGrid {
		rowLines := StringTileRow(line)
		ignoreFirstLine := strings.SplitN(rowLines, "\n", 2)[1]
		lines[i] = ignoreFirstLine
	}
	return strings.Join(lines, "\n")
}

func PuzzleTogether(tiles []*ImageTile, sideLength int) [][]*ImageTile {
	tiles = FindNeighbors(tiles)
	outGrid := make([][]*ImageTile, sideLength)

	var firstTile *ImageTile
	for _, tile := range tiles {
		if len(tile.Neighbors) == 2 {
			firstTile = tile
			break
		}
	}
	outGrid[0] = make([]*ImageTile, sideLength)
	outGrid[0][0] = firstTile // consider (0,0) as the top right corner

	firstTile.MakeTopRightUnique(tiles)
	prevTile := firstTile
	for x := 1; x < sideLength; x++ {
		nextTile := prevTile.NeighborWithEdge(prevTile.Left())
		nextTile.MakeRightEqual(prevTile.Left())
		outGrid[0][x] = nextTile
		prevTile = nextTile
	}

	for y := 1; y < sideLength; y++ {
		outGrid[y] = make([]*ImageTile, sideLength)
		upTile := outGrid[y-1][0]
		prevTile := upTile.NeighborWithEdge(upTile.Bottom())
		prevTile.MakeTopEqual(upTile.Bottom())
		outGrid[y][0] = prevTile
		for x := 1; x < sideLength; x++ {
			nextTile := prevTile.NeighborWithEdge(prevTile.Left())
			nextTile.MakeRightEqual(prevTile.Left())
			outGrid[y][x] = nextTile
			prevTile = nextTile
		}
	}

	for _, line := range outGrid {
		for _, tile := range line {
			tile.FlipVertically()
			tile.RotateAnticlockwise(2)
		}
	}

	return outGrid
}

func TileGridToImage(tileGrid [][]*ImageTile) [][]bool {
	outGrid := make([][]bool, len(tileGrid)*8)
	for i := range outGrid {
		outGrid[i] = make([]bool, len(tileGrid[0])*8)
	}

	for ty, tileRow := range tileGrid {
		for tx, tile := range tileRow {
			for gy := 0; gy < 8; gy++ {
				for gx := 0; gx < 8; gx++ {
					outGrid[ty*8+gy][tx*8+gx] = tile.At(gx+1, gy+1)
				}
			}
		}
	}

	return outGrid
}

func StringImage(image [][]bool) string {
	rowStrings := make([]string, len(image))
	for y, row := range image {
		rowString := make([]string, len(row))
		for x, value := range row {
			rowString[x] = "."
			if value {
				rowString[x] = "#"
			}
		}
		rowStrings[y] = strings.Join(rowString, "")
	}
	return strings.Join(rowStrings, "\n")
}

func RotateImage(image [][]bool) [][]bool {
	sideLength := len(image)
	result := make([][]bool, sideLength)
	for y := range result {
		result[y] = make([]bool, sideLength)
	}
	for y, row := range image {
		for x, value := range row {
			result[sideLength-1-x][y] = value
		}
	}
	return result
}

func FlipImage(image [][]bool) [][]bool {
	result := make([][]bool, len(image))
	for y := range image {
		ry := len(image) - 1 - y
		result[ry] = image[y]
	}
	return result
}

func AllImageOrientations(image [][]bool) [8][][]bool {
	result := [8][][]bool{}
	result[0] = image
	for i := 1; i <= 3; i++ {
		result[i] = RotateImage(result[i-1])
	}
	result[4] = FlipImage(image)
	for i := 5; i <= 7; i++ {
		result[i] = RotateImage(result[i-1])
	}
	return result
}

func seaSerpent() [][]bool {
	serpentString := `
..................#.
#....##....##....###
.#..#..#..#..#..#...
`
	serpent := make([][]bool, 3)
	for y := range serpent {
		serpent[y] = make([]bool, 20)
	}
	for y, row := range strings.Split(strings.Trim(serpentString, "\n"), "\n") {
		for x, char := range row {
			if char == '#' {
				serpent[y][x] = true
			}
		}
	}
	return serpent
}

var SeaSerpent = seaSerpent()

func HasSeaSerpentAt(image [][]bool, x, y int) bool {
	serpentHeight, serpentWidth := 3, 20
	imageHeight, imageWidth := len(image), len(image[0])
	if y+serpentHeight <= imageHeight && x+serpentWidth <= imageWidth {
		for dy := 0; dy < serpentHeight; dy++ {
			for dx := 0; dx < serpentWidth; dx++ {
				if SeaSerpent[dy][dx] && !image[y+dy][x+dx] {
					return false
				}
			}
		}
		return true
	}
	return false
}

func CountSeaSerpents(image [][]bool) int {
	count := 0
	for y, row := range image {
		for x := range row {
			if HasSeaSerpentAt(image, x, y) {
				count++
			}
		}
	}
	return count
}

func CountTrues(image [][]bool) int {
	count := 0
	for _, row := range image {
		for _, value := range row {
			if value {
				count++
			}
		}
	}
	return count
}

func Part2(input string, sideLength int) int {
	tiles := ReadInput(input)
	puzzle := PuzzleTogether(tiles, sideLength)
	image := TileGridToImage(puzzle)
	total := CountTrues(image)
	serpents := 0
	for _, orientation := range AllImageOrientations(image) {
		serpents = CountSeaSerpents(orientation)
		if serpents > 0 {
			break
		}
	}
	if serpents == 0 {
		panic("There should be serpents here")
	}
	return total - serpents*15
}
