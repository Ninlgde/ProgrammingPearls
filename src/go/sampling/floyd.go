package sampling

import (
	"github.com/Ninlgde/ProgrammingPearls/src/go/utils"
	"github.com/deckarep/golang-set"
)

func GenFloyd(m int, n int) []int {
	result := make([]int, 0, m)

	S := mapset.NewSet()

	for i := n - m; i < n; i++ {
		t := utils.BigRand() % (i + 1)
		if S.Contains(t) {
			S.Add(i) // t in S
		} else {
			S.Add(t) // t not in S
		}
	}

	for val := range S.Iter() {
		result = append(result, val.(int))
	}
	return result
}
