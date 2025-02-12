package main

import (
	l "CustomSort/support"
	"CustomSort/insertsort"
	"CustomSort/bubblesort"
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"time"
)

func main() {
	testeFunc()
}

func testeFunc() {
	fmt.Println("Standard sorting lib")

	values := []int{7, 2, 1, 1, 6, 4, 8, 10, 3, 9}
	fmt.Println("Before Sorting:", values)
	sort.Ints(values)
	fmt.Println("After Sorting:", values)
	fmt.Println("---------------------------------")

	fmt.Println("My IndexSorting implementation")

	valuesList := l.List(7, 2, 1, 1, 6, 4, 8, 10, 3, 9)
	fmt.Println("Before Sorting:", valuesList)
	insertsort.InsertSort(valuesList, true)
	fmt.Println("After Sorting:", valuesList)
	fmt.Println("---------------------------------")

	if slices.Equal(valuesList, values) {
		fmt.Println("Listas iguais")
		return
	}

	fmt.Println("My BubbleSorting implementation")

	valuesList = l.List(7, 2, 1, 1, 6, 4, 8, 10, 3, 9)
	fmt.Println("Before Sorting:", valuesList)
	bubblesort.BubbleSort(valuesList)
	fmt.Println("After Sorting:", valuesList)
	fmt.Println("---------------------------------")

	if slices.Equal(valuesList, values) {
		fmt.Println("Listas iguais")
		return
	}


	fmt.Println("Listas diferentes")
}

func Shuffle(valuesList []int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range valuesList {
		j := rand.Intn(len(valuesList))
		valuesList[i], valuesList[j] = valuesList[j], valuesList[i]
	}
}
