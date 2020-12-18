package day15

type GameHistory struct {
	Time           int
	NextNumber     int
	MostRecentTime map[int]int
}

func (gh *GameHistory) Reset(input []int) {
	gh.Time = len(input)
	gh.NextNumber = input[len(input)-1]
	gh.MostRecentTime = make(map[int]int)
	for i, value := range input[:len(input)-1] {
		gh.MostRecentTime[value] = i
	}
}

func (gh *GameHistory) Step() {
	lastSeen, beenSeen := gh.MostRecentTime[gh.NextNumber]
	nextNum := gh.Time - 1 - lastSeen
	if !beenSeen {
		nextNum = 0
	}

	gh.MostRecentTime[gh.NextNumber] = gh.Time - 1
	gh.Time++
	gh.NextNumber = nextNum
}

// func (gh GameHistory) FullHistory() []int {
// 	return append(gh.History, gh.NextNumber)
// }

// func (gh GameHistory) Time() int {
// 	return len(gh.History) + 1
// }

func RunUntil(input []int, endTime int) int {
	game := GameHistory{}
	game.Reset(input)

	for game.Time < endTime {
		game.Step()
	}

	return game.NextNumber
}
