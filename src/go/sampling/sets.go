package sampling

import (
	"github.com/Ninlgde/ProgrammingPearls/src/go/utils"
	"github.com/deckarep/golang-set"
)

func GenSets(m int, n int) []int {
	result := make([]int, m)

	S := mapset.NewSet()
	for S.Cardinality() < m {
		S.Add(utils.BigRand() % n)
	}
	j := 0
	for val := range S.Iter() {
		result[j] = val.(int)
		j++
	}
	return result
}
