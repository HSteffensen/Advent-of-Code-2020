package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func SeatNumber(boardingPass string) (row, col, seatId int) {
	binarify := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
	binPass := binarify.Replace(boardingPass)
	seatID64, err := strconv.ParseInt(binPass, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	seatId = int(seatID64)
	row = seatId >> 3
	col = seatId & 7
	return
}

func MaxSeatId(boardingPasses []string) int {
	largest := 0
	for _, pass := range boardingPasses {
		_, _, seatId := SeatNumber(pass)
		if seatId > largest {
			largest = seatId
		}
	}
	return largest
}

func MySeatId(boardingPasses []string) int {
	seatsTaken := make([]bool, MaxSeatId(boardingPasses)+1)
	for _, pass := range boardingPasses {
		_, _, seatId := SeatNumber(pass)
		if seatId > len(seatsTaken) {
			seatsTaken = seatsTaken[:seatId]
		}
		seatsTaken[seatId] = true
	}
	for i, taken := range seatsTaken[1:] {
		i++
		if !taken && seatsTaken[i-1] && seatsTaken[i+1] {
			return i
		}
	}
	return -1
}
