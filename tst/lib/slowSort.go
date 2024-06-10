package lib

func SlowSort(base []int, order SortOrder) []int {
	var containerValue int
	var containerKey int

	skipList := make(map[int]interface{})
	result := make([]int, len(base))

	for index, _ := range result {
		firstElement := true

		for k, v := range base {
			_, ok := skipList[k]

			if ok {
				continue
			}

			var isBetter bool
			switch order {
			case ASC:
				isBetter = v < containerValue
			case DESC:
				isBetter = v > containerValue
			default:
				isBetter = v < containerValue
			}

			if isBetter || firstElement {
				containerValue = v
				containerKey = k
			}

			firstElement = false
		}

		skipList[containerKey] = nil
		result[index] = containerValue
	}

	return result
}
