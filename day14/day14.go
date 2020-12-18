package day14

import (
	"fmt"
	"strconv"
	"strings"
)

type Mask struct {
	Zeroes int
	Ones   int
}

func (m Mask) ApplyTo(value int) int {
	return (value & (-1 ^ m.Zeroes)) | m.Ones
}

func ReadMask(maskString string) Mask {
	mask, counter := Mask{}, 1
	for i := len(maskString) - 1; i >= 0; i-- {
		switch maskString[i] {
		case '0':
			mask.Zeroes |= counter
		case '1':
			mask.Ones |= counter
		case 'X':
		default:
			panic(fmt.Sprintf("Bad mask input: %#v at position %v is %#v", maskString, i, maskString[i]))
		}
		counter <<= 1
	}
	return mask
}

func SingleMask(maskString string) []Mask {
	return []Mask{ReadMask(maskString)}
}

func helperMultiMask(maskString string) []string {
	if len(maskString) <= 0 {
		return []string{""}
	}

	createdStrings := helperMultiMask(maskString[1:])
	currentBitStrings := []string{maskString[0:1]}
	if maskString[0] == 'X' {
		currentBitStrings = []string{"0", "1"}
	} else if maskString[0] == '0' {
		currentBitStrings = []string{"X"}
	}
	resultStrings := make([]string, 0, len(createdStrings)*len(currentBitStrings))

	for _, firstBit := range currentBitStrings {
		for _, suffix := range createdStrings {
			resultStrings = append(resultStrings, firstBit+suffix)
		}
	}

	return resultStrings
}

func MultiMask(maskString string) []Mask {
	allMaskStrings := helperMultiMask(maskString)
	allMasks := make([]Mask, len(allMaskStrings))

	for i, maskStr := range allMaskStrings {
		allMasks[i] = ReadMask(maskStr)
	}

	return allMasks
}

func Version1Data(data map[int]int, position, value int, mask Mask) {
	data[position] = mask.ApplyTo(value)
}

func Version2Data(data map[int]int, position, value int, mask Mask) {
	data[mask.ApplyTo(position)] = value
}

func BuildData(input string, masker func(string) []Mask, dataSetter func(map[int]int, int, int, Mask)) map[int]int {
	lines := strings.Split(input, "\n")
	data := make(map[int]int)
	currentMasks := []Mask{{}}

	for i, line := range lines {
		splitLine := strings.Split(line, " = ")
		kind, value := splitLine[0], splitLine[1]
		switch kind[0:4] {
		case "mask":
			currentMasks = masker(value)
		case "mem[":
			position, err := strconv.Atoi(kind[4 : len(kind)-1])
			if err != nil {
				panic(fmt.Sprintf("Bad day14 input line mem position on line %v: %#v", i, kind))
			}
			inValue, err := strconv.Atoi(value)
			if err != nil {
				panic(fmt.Sprintf("Bad day14 input line value on line %v: %#v", i, value))
			}
			for _, mask := range currentMasks {
				dataSetter(data, position, inValue, mask)
			}
		default:
			panic(fmt.Sprintf("Bad day14 input line start on line %v: %#v", i, kind))
		}
	}

	return data
}

func Part1(input string) int {
	data := BuildData(input, SingleMask, Version1Data)
	total := 0

	for _, value := range data {
		total += value
	}

	return total
}

func Part2(input string) int {
	data := BuildData(input, MultiMask, Version2Data)
	total := 0

	for _, value := range data {
		total += value
	}

	return total
}
