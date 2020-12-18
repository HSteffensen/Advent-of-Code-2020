package aocCommon

func AssertSameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	itemAppearsTimes := make(map[string]int, len(x))
	for _, i := range x {
		itemAppearsTimes[i]++
	}

	for _, i := range y {
		if _, ok := itemAppearsTimes[i]; !ok {
			return false
		}

		itemAppearsTimes[i]--

		if itemAppearsTimes[i] == 0 {
			delete(itemAppearsTimes, i)
		}
	}

	if len(itemAppearsTimes) == 0 {
		return true
	}

	return false
}
