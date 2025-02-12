package bubblesort

import (
	"sort"
	"testing"
	c "CustomSort/support"
)

// Benchmark para a implementação padrão
func BenchmarkStandardSort(b *testing.B) {
	b.ResetTimer() // Reseta o timer do benchmark
	for i := 0; i < b.N; i++ {
		slice := c.GenerateRandomList(1000)
		sort.Ints(slice)
	}
}

// Benchmark para a implementação personalizada
func BenchmarkCustomInserSort(b *testing.B) {
	b.ResetTimer() // Reseta o timer do benchmark
	for i := 0; i < b.N; i++ {
		slice := c.GenerateRandomList(1000)
		BubbleSort(slice)
	}
}
