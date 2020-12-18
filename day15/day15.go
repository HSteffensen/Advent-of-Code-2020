package day15

type GameHistory struct {
	Time           uint32
	NextNumber     uint32
	MostRecentTime [30000000]uint32
}

func (gh *GameHistory) Reset(input []int) {
	gh.Time = uint32(len(input))
	gh.NextNumber = uint32(input[len(input)-1])
	for i, value := range input[:len(input)-1] {
		gh.MostRecentTime[value] = uint32(i + 1)
	}
}

func (gh *GameHistory) Step() {
	lastSeen := gh.MostRecentTime[gh.NextNumber]
	nextNum := uint32(0)
	if lastSeen > 0 {
		nextNum = gh.Time - lastSeen
	}

	gh.MostRecentTime[gh.NextNumber] = gh.Time
	gh.Time++
	gh.NextNumber = nextNum
}

func RunUntil(input []int, endTime int) int {
	game := GameHistory{}
	game.Reset(input)

	uEndTime := uint32(endTime)
	for game.Time < uEndTime {
		game.Step()
	}

	return int(game.NextNumber)
}
