package day1

func FindTripleSummingToN(numbers []int, targetN int) (int, int, int) {
	cache := make(map[int][2]int)
	for i, v1 := range numbers {
		for _, v2 := range numbers[:i] {
			cache[v1+v2] = [2]int{v1, v2}
		}
	}
	for _, v := range numbers {
		pair, exists := cache[targetN-v]
		if exists {
			return v, pair[0], pair[1]
		}
	}
	return 0, 0, 0
}
