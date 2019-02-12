package sampling

import (
	"github.com/Ninlgde/ProgrammingPearls/src/go/sort"
	"github.com/Ninlgde/ProgrammingPearls/src/go/utils"
)

func GenShuf(m int, n int) []int {
	result := make([]int, m)
	x := make([]int, n)
	for i := range x {
		x[i] = i
	}
	for i := 0; i < m; i++ {
		j := utils.RandInt(i, n-1)
		x[i], x[j] = x[j], x[i]
		result[i] = x[i]
	}
	sort.QuickSort(result)
	return result
}
