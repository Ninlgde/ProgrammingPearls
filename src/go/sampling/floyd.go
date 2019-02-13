package sampling

import (
	"github.com/Ninlgde/ProgrammingPearls/src/go/simpleset"
	"github.com/Ninlgde/ProgrammingPearls/src/go/utils"
	"github.com/deckarep/golang-set"
)

func GenFloyd(m int, n int) []int {
	result := make([]int, m)

	S := mapset.NewSet()

	for i := n - m; i < n; i++ {
		t := utils.BigRand() % (i + 1)
		if S.Contains(t) {
			S.Add(i) // t in S
		} else {
			S.Add(t) // t not in S
		}
	}

	j := 0
	for val := range S.Iter() {
		result[j] = val.(int)
		j++
	}
	return result
}

func GenFloydFaster(m int, n int) []int {
	result := make([]int, m)

	S := simpleset.NewSet()

	for i := n - m; i < n; i++ {
		t := utils.BigRand() % (i + 1)
		if S.Contains(t) {
			S.Add(i) // t in S
		} else {
			S.Add(t) // t not in S
		}
	}

	j := 0
	for val := range S.Iter() {
		result[j] = val.(int)
		j++
	}
	return result
}
