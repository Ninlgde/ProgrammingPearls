package sampling

import (
	"github.com/Ninlgde/ProgrammingPearls/src/go/utils"
	"github.com/deckarep/golang-set"
)

func GenSets(m int, n int) []int {
	result := make([]int, 0, m)

	S := mapset.NewSet()
	for S.Cardinality() < m {
		S.Add(utils.BigRand() % n)
	}
	for val := range S.Iter() {
		result = append(result, val.(int))
	}
	return result
}
