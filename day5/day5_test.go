package day5_test

import (
	"HSteffensen/AoC2020/day5"
	"testing"
)

func TestSeatNumber(t *testing.T) {
	row, col, seatId := day5.SeatNumber("FBFBBFFRLR")
	expRow, expCol, expId := 44, 5, 357
	if row != expRow {
		t.Errorf("Incorrect row: expected %v, got %v", expRow, row)
	}
	if col != expCol {
		t.Errorf("Incorrect col: expected %v, got %v", expCol, col)
	}
	if seatId != expId {
		t.Errorf("Incorrect seatId: expected %v, got %v", expId, seatId)
	}
}
