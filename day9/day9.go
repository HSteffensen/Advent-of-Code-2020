package day9

func ValidNextNumber(history []int, next int) bool {
	cache := make(map[int]int)
	for i, v := range history {
		cache[v] = i
	}
	for i, v := range history {
		j, exists := cache[next-v]
		if exists && i != j {
			return true
		}
	}
	return false
}

func FirstInvalidIndex(numbers []int, historySize int) int {
	for i := historySize; i <= len(numbers); i++ {
		if !ValidNextNumber(numbers[i-historySize:i], numbers[i]) {
			return i
		}
	}
	return -1
}

func swapIfSmaller(smallest, second int) (int, int) {
	if second < smallest {
		return second, smallest
	}
	return smallest, second
}

func sumSmallestLargest(numbers []int) int {
	smallest, largest := swapIfSmaller(numbers[0], numbers[1])
	for _, n := range numbers[2:] {
		n, largest = swapIfSmaller(n, largest)
		smallest, n = swapIfSmaller(smallest, n)
	}
	return smallest + largest
}

func FindEncryptionWeakness(numbers []int, historySize int) int {
	target := numbers[FirstInvalidIndex(numbers, historySize)]

	i, j := 0, 2
	total := numbers[0] + numbers[1]
	for i+j < len(numbers) {
		switch {
		case total == target:
			return sumSmallestLargest(numbers[i:j])
		case total < target:
			total += numbers[j]
			j++
		case total > target:
			total -= numbers[i]
			i++
		}
	}
	return -1
}
