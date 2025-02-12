package bubblesort

func BubbleSort(lista []int) []int {
	n := len(lista)
	var swapped bool
	for i := range n {
		swapped = false
		for j := range n-i-1 {
			if lista[j] > lista[j+1]{
				lista[j], lista[j+1] = lista[j+1], lista[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return lista
}
