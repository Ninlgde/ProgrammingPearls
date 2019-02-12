package intset

import "github.com/Ninlgde/ProgrammingPearls/src/go/search"

type intSetArray struct {
	n int
	x []int
}

func newIntSetArray(maxelements int, maxval int) intSetArray {
	set := intSetArray{0, make([]int, maxelements+1)}
	set.IntSetImp(maxelements, maxval)
	return set
}

func (set *intSetArray) IntSetImp(maxelements int, maxval int) {
	set.x[0] = maxval
}

func (set *intSetArray) Insert(t int) {
	i := 0
	for ; set.x[i] < t; i++ {
	}
	if set.x[i] == t {
		return
	}
	for j := set.n; j >= i; j-- {
		set.x[j+1] = set.x[j]
	}
	set.x[i] = t
	set.n++
}

func (set *intSetArray) Size() int {
	return set.n
}

func (set *intSetArray) Report() []int {
	result := make([]int, set.Size())
	for i := 1; i <= set.n; i++ {
		result[i-1] = set.x[i]
	}
	return result
}

func (set *intSetArray) Has(t int) bool {
	return search.BinarySearch(set.x[1:set.n+1], t) != -1
}

func (set *intSetArray) Remove(t int) {
	i := 0
	for ; set.x[i] < t; i++ {
	}
	if set.x[i] == t {
		for j := i; j < set.n; j++ {
			set.x[j] = set.x[j+1]
		}
		set.n--
	}
}
