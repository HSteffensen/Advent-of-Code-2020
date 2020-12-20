package day17

type Point3Dlike interface {
	Plus(Point3Dlike) Point3Dlike
	Minus(Point3Dlike) Point3Dlike
	Neighbors() []Point3Dlike
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) Plus(other Point3Dlike) Point3Dlike {
	o := other.(Point3D)
	p.X += o.X
	p.Y += o.Y
	p.Z += o.Z
	return p
}

func (p Point3D) Minus(other Point3Dlike) Point3Dlike {
	o := other.(Point3D)
	p.X -= o.X
	p.Y -= o.Y
	p.Z -= o.Z
	return p
}

func (p Point3D) Neighbors() []Point3Dlike {
	resultList := make([]Point3Dlike, 0, 26)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				resultList = append(resultList, p.Plus(Point3D{dx, dy, dz}))
			}
		}
	}
	return resultList
}

type Point4D struct {
	X, Y, Z, W int
}

func (p Point4D) Plus(other Point3Dlike) Point3Dlike {
	o := other.(Point4D)
	p.X += o.X
	p.Y += o.Y
	p.Z += o.Z
	p.W += o.W
	return p
}

func (p Point4D) Minus(other Point3Dlike) Point3Dlike {
	o := other.(Point4D)
	p.X -= o.X
	p.Y -= o.Y
	p.Z -= o.Z
	p.W -= o.W
	return p
}

func (p Point4D) Neighbors() []Point3Dlike {
	resultList := make([]Point3Dlike, 0, 80)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					resultList = append(resultList, p.Plus(Point4D{dx, dy, dz, dw}))
				}
			}
		}
	}
	return resultList
}

type CubeGrid map[Point3Dlike]bool

func NewCubeGridWithXYPlane(grid [][]bool) CubeGrid {
	cg := make(CubeGrid)
	for y, row := range grid {
		for x, value := range row {
			if value {
				point := Point3D{x, y, 0}
				cg[point] = true
			}
		}
	}
	return cg
}

func NewHyperCubeGridWithXYPlane(grid [][]bool) CubeGrid {
	cg := make(CubeGrid)
	for y, row := range grid {
		for x, value := range row {
			if value {
				point := Point4D{x, y, 0, 0}
				cg[point] = true
			}
		}
	}
	return cg
}

func (cg CubeGrid) Step() CubeGrid {
	counts := make(map[Point3Dlike]int)
	nextGrid := make(CubeGrid)
	for pos, active := range cg {
		if active {
			for _, neighbor := range pos.Neighbors() {
				counts[neighbor]++
			}
		}
	}
	for pos, count := range counts {
		switch count {
		case 3:
			nextGrid[pos] = true
		case 2:
			if cg[pos] {
				nextGrid[pos] = true
			}
		}
	}
	return nextGrid
}

func (cg CubeGrid) CountActive() int {
	count := 0
	for _, active := range cg {
		if active {
			count++
		}
	}
	return count
}

func Part1(input [][]bool, stepCount int) int {
	cg := NewCubeGridWithXYPlane(input)
	for i := 0; i < stepCount; i++ {
		cg = cg.Step()
	}
	return cg.CountActive()
}

func Part2(input [][]bool, stepCount int) int {
	cg := NewHyperCubeGridWithXYPlane(input)
	for i := 0; i < stepCount; i++ {
		cg = cg.Step()
	}
	return cg.CountActive()
}
