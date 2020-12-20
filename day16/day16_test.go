package day16_test

import (
	"HSteffensen/AoC2020/day16"
	"reflect"
	"testing"
)

func TestReadInputFields(t *testing.T) {
	if expected, result := (day16.TicketField{"class", 1, 3, 5, 7}), day16.ReadInputFields("class: 1-3 or 5-7")[0]; expected != result {
		t.Errorf("Incorrect read input fields: expected %v, got %v", expected, result)
	}
	if expected, result := (day16.TicketField{"class", 1, 3, 5, 7}), day16.ReadInputFields(day16.ExampleInput1)[0]; expected != result {
		t.Errorf("Incorrect read input fields: expected %v, got %v", expected, result)
	}
	if expected, result := 3, len(day16.ReadInputFields(day16.ExampleInput1)); expected != result {
		t.Errorf("Incorrect read input fields length: expected %v, got %v", expected, result)
	}
}

func TestReadYourTicket(t *testing.T) {
	if expected, result := (day16.Ticket{7, 1, 14}), day16.ReadYourTicket(day16.ExampleInput1); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect your ticket: expected %v, got %v", expected, result)
	}
	if expected, result := (day16.Ticket{97, 61, 53, 101, 131, 163, 79, 103, 67, 127, 71, 109, 89, 107, 83, 73, 113, 59, 137, 139}), day16.ReadYourTicket(day16.Input); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect your ticket: expected %v, got %v", expected, result)
	}
}

func TestReadNearbyTickets(t *testing.T) {
	if expected, result := (day16.Ticket{7, 3, 47}), day16.ReadNearbyTickets(day16.ExampleInput1)[0]; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect nearby ticket: expected %v, got %v", expected, result)
	}
	if expected, result := 4, len(day16.ReadNearbyTickets(day16.ExampleInput1)); expected != result {
		t.Errorf("Incorrect nearby ticket count: expected %v, got %v", expected, result)
	}
	if expected, result := (day16.Ticket{609, 949, 551, 408, 648, 723, 627, 428, 905, 665, 248, 557, 630, 169, 494, 583, 680, 658, 677, 385}), day16.ReadNearbyTickets(day16.Input)[1]; !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect nearby ticket: expected %v, got %v", expected, result)
	}
}

func TestTicketFieldInRange(t *testing.T) {
	ticketFields := []day16.TicketField{{"class", 1, 3, 5, 7}, {"row", 6, 11, 33, 44}, {"seat", 13, 40, 45, 50}}

	if ticket, value := ticketFields[0], 1; !ticket.InRange(value) {
		t.Errorf("Incorrect ticketfield in range: %v should be in range of %v", value, ticket)
	}
	if ticket, value := ticketFields[0], 2; !ticket.InRange(value) {
		t.Errorf("Incorrect ticketfield in range: %v should be in range of %v", value, ticket)
	}
	if ticket, value := ticketFields[0], 3; !ticket.InRange(value) {
		t.Errorf("Incorrect ticketfield in range: %v should be in range of %v", value, ticket)
	}
	if ticket, value := ticketFields[0], 5; !ticket.InRange(value) {
		t.Errorf("Incorrect ticketfield in range: %v should be in range of %v", value, ticket)
	}
	if ticket, value := ticketFields[0], 6; !ticket.InRange(value) {
		t.Errorf("Incorrect ticketfield in range: %v should be in range of %v", value, ticket)
	}
	if ticket, value := ticketFields[0], 7; !ticket.InRange(value) {
		t.Errorf("Incorrect ticketfield in range: %v should be in range of %v", value, ticket)
	}

	if ticket, value := ticketFields[0], 0; ticket.InRange(value) {
		t.Errorf("Incorrect ticketfield in range: %v should not be in range of %v", value, ticket)
	}
	if ticket, value := ticketFields[0], 4; ticket.InRange(value) {
		t.Errorf("Incorrect ticketfield in range: %v should not be in range of %v", value, ticket)
	}
	if ticket, value := ticketFields[0], 8; ticket.InRange(value) {
		t.Errorf("Incorrect ticketfield in range: %v should not be in range of %v", value, ticket)
	}
}

func TestTicketErrorRate(t *testing.T) {
	ticketFields := []day16.TicketField{{"class", 1, 3, 5, 7}, {"row", 6, 11, 33, 44}, {"seat", 13, 40, 45, 50}}
	if expected, result := 0, day16.TicketErrorRate(day16.Ticket{7, 3, 47}, ticketFields); expected != result {
		t.Errorf("Incorrect ticket error: expected %v, got %v", expected, result)
	}
	if expected, result := 4, day16.TicketErrorRate(day16.Ticket{40, 4, 50}, ticketFields); expected != result {
		t.Errorf("Incorrect ticket error: expected %v, got %v", expected, result)
	}
	if expected, result := 55, day16.TicketErrorRate(day16.Ticket{55, 2, 20}, ticketFields); expected != result {
		t.Errorf("Incorrect ticket error: expected %v, got %v", expected, result)
	}
	if expected, result := 12, day16.TicketErrorRate(day16.Ticket{38, 16, 12}, ticketFields); expected != result {
		t.Errorf("Incorrect ticket error: expected %v, got %v", expected, result)
	}
}

func TestPart1(t *testing.T) {
	if expected, result := 71, day16.Part1(day16.ExampleInput1); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}

func TestDetermineFields(t *testing.T) {
	ticketFields := []day16.TicketField{{"class", 0, 1, 4, 19}, {"row", 0, 5, 8, 19}, {"seat", 0, 13, 16, 19}}
	nearbyTickets := []day16.Ticket{{3, 9, 18}, {15, 1, 5}, {5, 14, 9}}
	expected := []day16.TicketField{ticketFields[1], ticketFields[0], ticketFields[2]}
	if result := day16.DetermineFields(ticketFields, nearbyTickets); !reflect.DeepEqual(expected, result) {
		t.Errorf("Incorrect determined fields")
	}
}

func TestPart2(t *testing.T) {
	if expected, result := 12, day16.Part2(day16.ExampleInput2, "class"); expected != result {
		t.Errorf("Incorrect part 1: expected %v, got %v", expected, result)
	}
}
