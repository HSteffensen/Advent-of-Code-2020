package day16

import "strings"

type TicketField struct {
	Name string
	Min1 int
	Max1 int
	Min2 int
	Max2 int
}

type Ticket []int

func (tf TicketField) InRange(value int) bool {
	return (value >= tf.Min1 && value <= tf.Max1) || (value >= tf.Min2 && value <= tf.Max2)
}

func TicketErrorRate(ticket Ticket, fields []TicketField) int {
	total := 0
	for _, value := range ticket {
		valueOk := false
		for _, field := range fields {
			if field.InRange(value) {
				valueOk = true
				break
			}
		}
		if !valueOk {
			total += value
		}
	}
	return total
}

func Part1(input string) int {
	ticketFields := ReadInputFields(input)
	tickets := ReadNearbyTickets(input)
	total := 0
	for _, ticket := range tickets {
		total += TicketErrorRate(ticket, ticketFields)
	}
	return total
}

func IsValidTicket(ticket Ticket, fields []TicketField) bool {
	result := true
	for _, value := range ticket {
		valueOk := false
		for _, field := range fields {
			if field.InRange(value) {
				valueOk = true
				break
			}
		}
		result = result && valueOk
	}
	return result
}

func DetermineFields(fields []TicketField, tickets []Ticket) []TicketField {
	fieldMasks := make(map[int]TicketField)
	for i, field := range fields {
		fieldMasks[1<<i] = field
	}
	goodTickets := make([]Ticket, 0, len(tickets))
	for _, ticket := range tickets {
		if IsValidTicket(ticket, fields) {
			goodTickets = append(goodTickets, ticket)
		}
	}

	positionPossibilities := make([]int, len(fieldMasks))
	for i := range positionPossibilities {
		positionPossibilities[i] = (1 << len(fields)) - 1
		for _, ticket := range goodTickets {
			for j, mask := range fieldMasks {
				if positionPossibilities[i]&j != 0 && !mask.InRange(ticket[i]) {
					positionPossibilities[i] &= (-1 ^ j)
				}
			}
		}
	}

	resultFields := make([]TicketField, len(fields))

	for done := false; !done; {
		done = true
		for i, posBits := range positionPossibilities {
			if posBits == 0 {
				continue
			}
			if field, ok := fieldMasks[posBits]; !ok {
				done = false
				continue
			} else {
				resultFields[i] = field
				for j := range positionPossibilities {
					positionPossibilities[j] &= (-1 ^ posBits)
				}
			}
		}
	}

	return resultFields
}

func Part2(input, prefix string) int {
	fields := ReadInputFields(input)
	yourTicket := ReadYourTicket(input)
	nearbyTickets := ReadNearbyTickets(input)

	fieldOrder := DetermineFields(fields, append(nearbyTickets, yourTicket))

	result := 1
	for i, field := range fieldOrder {
		if strings.HasPrefix(field.Name, prefix) {
			result *= yourTicket[i]
		}
	}
	return result
}
