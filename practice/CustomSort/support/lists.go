package support

import (
	"math/rand"
	"time"
)

func List(args ...int) []int {
	list := make([]int, len(args))
	copy(list, args)
	return list
}

func GenerateRandomList(numValues int) []int {
	slice := make([]int, numValues)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(len(slice))
	}
	return slice
}
