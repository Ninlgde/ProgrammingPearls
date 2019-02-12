package search

import (
	"github.com/Ninlgde/ProgrammingPearls/src/go/utils"
)

func MaxRange(x []int) int {
	return MaxRange5(x)
}

func MaxRange1(x []int) int {
	maxsofar := 0

	for i := range x {
		for j := i; j < len(x); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += x[k]
			}
			maxsofar = utils.Max(maxsofar, sum)
		}
	}
	return maxsofar
}

func MaxRange2(x []int) int {
	maxsofar := 0

	for i := range x {
		sum := 0
		for j := i; j < len(x); j++ {
			sum += x[j]
			maxsofar = utils.Max(maxsofar, sum)
		}
	}
	return maxsofar
}

func MaxRange3(x []int) int {

	cumarr := make(map[int]int)
	cumarr[-1] = 0
	for i := range x {
		cumarr[i] = cumarr[i-1] + x[i]
	}

	maxsofar := 0

	for i := range x {
		//var sum float64 = 0
		for j := i; j < len(x); j++ {
			sum := cumarr[j] - cumarr[i-1]
			maxsofar = utils.Max(maxsofar, sum)
		}
	}

	return maxsofar
}

func MaxRange4(x []int) int {
	return maxsum3(x)
}

func maxsum3(x []int) int {
	if len(x) == 0 {
		return 0
	}
	if len(x) == 1 {
		return utils.Max(0, x[0])
	}
	l := 0
	u := len(x)
	m := (l + u) / 2
	lmax := 0
	rmax := 0
	sum := 0
	for i := m; i >= l; i-- {
		sum += x[i]
		lmax = utils.Max(lmax, sum)
	}
	sum = 0
	for i := m + 1; i < u; i++ {
		sum += x[i]
		rmax = utils.Max(rmax, sum)
	}
	return max(lmax+rmax, maxsum3(x[l:m]), maxsum3(x[m:u]))
}

func max(a int, b int, c int) int {
	return utils.Max(a, utils.Max(b, c))
}

func MaxRange5(x []int) int {
	maxsofar := 0
	maxendinghere := 0
	for i := range x {
		maxendinghere = utils.Max(maxendinghere+x[i], 0)
		maxsofar = utils.Max(maxsofar, maxendinghere)
	}

	return maxsofar
}
