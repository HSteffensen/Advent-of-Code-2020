package day8

type Program []Instruction

type Instruction struct {
	OpCode string
	Arg    int
}

var Operations = map[string]func(int, int, int) (int, int){
	"nop": func(acc, value, pos int) (int, int) {
		return acc, pos + 1
	},
	"acc": func(acc, value, pos int) (int, int) {
		return acc + value, pos + 1
	},
	"jmp": func(acc, value, pos int) (int, int) {
		return acc, pos + value
	},
}

func (p Program) Step(acc, pos int) (int, int) {
	return Operations[p[pos].OpCode](acc, p[pos].Arg, pos)
}

func (p Program) loopHelper(acc, pos int) (int, int, map[int]bool) {
	visited := make(map[int]bool)
	for acc, pos := 0, 0; ; {
		if pos < 0 || pos >= len(p) || visited[pos] {
			return acc, pos, visited
		}
		visited[pos] = true
		acc, pos = p.Step(acc, pos)
	}
}

func (p Program) Loops() bool {
	_, pos, _ := p.loopHelper(0, 0)
	return pos >= 0 && pos < len(p)
}

func (p Program) FindLoopAcc() int {
	acc, _, _ := p.loopHelper(0, 0)
	return acc
}

func (p Program) FixLoopAcc() int {
	swapJmpNop := func(i int) bool {
		switch p[i].OpCode {
		case "nop":
			p[i].OpCode = "jmp"
			return true
		case "jmp":
			p[i].OpCode = "nop"
			return true
		default:
			return true
		}
	}
	for i := range p {
		if swapJmpNop(i) {
			acc, pos, _ := p.loopHelper(0, 0)
			if pos == len(p) {
				return acc
			}
			swapJmpNop(i)
		}
	}
	panic("Failed to fix Program loop")
}
