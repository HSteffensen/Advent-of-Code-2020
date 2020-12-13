package day13

import "strconv"

func TimeUntilBus(now, busID int) int {
	return busID - (now % busID)
}

func Part1(input string) int {
	time, busses := ReadInput(input)
	lowestWait, lowestID := time, 0

	for _, busID := range busses {
		if busID == "x" {
			continue
		}
		busVal, err := strconv.Atoi(busID)
		if err != nil {
			panic("Bad bus ID for day13")
		}
		if busTime := TimeUntilBus(time, busVal); busTime < lowestWait {
			lowestWait = busTime
			lowestID = busVal
		}
	}
	return lowestWait * lowestID
}

func FindDeparture(start, first, second, offset int) int {
	for ; (start+offset)%second != 0; start += first {
	}
	return start
}

func Part2(input string) int {
	_, busses := ReadInput(input)
	startTime, product := 0, 1
	for i, busID := range busses {
		if busID == "x" {
			continue
		}
		busVal, err := strconv.Atoi(busID)
		if err != nil {
			panic("Bad bus ID for day13")
		}

		startTime = FindDeparture(startTime, product, busVal, i)
		product *= busVal
	}
	return startTime
}
