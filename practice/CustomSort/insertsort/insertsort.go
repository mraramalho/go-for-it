package insertsort

func InsertSort(valuesList []int, ascendingOrder bool) {
	for i := 1; i < len(valuesList); i++ {
		j := i
		if ascendingOrder {
			for j > 0 && valuesList[j] < valuesList[j-1] {
				valuesList[j], valuesList[j-1] = valuesList[j-1], valuesList[j]
				j--
			}
		} else {
			for j > 0 && valuesList[j] > valuesList[j-1] {
				valuesList[j], valuesList[j-1] = valuesList[j-1], valuesList[j]
				j--
			}
		}
	}
}
