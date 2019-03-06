package sort

import (
	"github.com/Ninlgde/ProgrammingPearls/src/go/utils"
	"testing"
)

func makeASlice(n int) []int {
	slice := utils.GenerateSlice(10000)
	return slice
}

func BenchmarkQSort1(b *testing.B) {
	l := makeASlice(b.N)
	for i := 0; i <= b.N; i++ {
		tmp := make([]int, len(l))
		QSort1(tmp)
	}
}

func BenchmarkQSort2(b *testing.B) {
	l := makeASlice(b.N)
	for i := 0; i <= b.N; i++ {
		tmp := make([]int, len(l))
		QSort2(tmp)
	}
}

func BenchmarkQSort3(b *testing.B) {
	l := makeASlice(b.N)
	for i := 0; i <= b.N; i++ {
		tmp := make([]int, len(l))
		QSort3(tmp)
	}
}

func BenchmarkQSort4(b *testing.B) {
	l := makeASlice(b.N)
	for i := 0; i <= b.N; i++ {
		tmp := make([]int, len(l))
		QSort4(tmp)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	l := makeASlice(b.N)
	for i := 0; i <= b.N; i++ {
		tmp := make([]int, len(l))
		QuickSort(tmp)
	}
}
