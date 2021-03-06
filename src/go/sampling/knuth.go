package sampling

import "github.com/Ninlgde/ProgrammingPearls/src/go/utils"

func GenKnuth(m int, n int) []int {
	result := make([]int, 0, m)
	for i := 0; i < n; i++ {
		if (utils.BigRand() % (n - i)) < m {
			result = append(result, i)
			m--
		}
	}
	return result
}

func GenKnuthFaster(m int, n int) []int {
	result := make([]int, m)
	max := m
	for i := 0; i < n; i++ {
		if (utils.BigRand() % (n - i)) < m {
			result[max-m] = i
			m--
		}
	}
	return result
}
