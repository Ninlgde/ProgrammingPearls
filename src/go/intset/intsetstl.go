package intset

import (
	"github.com/Ninlgde/ProgrammingPearls/src/go/sort"
	"github.com/deckarep/golang-set"
)

type intSetSTL struct {
	S mapset.Set
}

func newIntSetSTL(maxelements int, maxval int) intSetSTL {
	set := intSetSTL{S: mapset.NewSet()}
	set.IntSetImp(maxelements, maxval)
	return set
}

func (set *intSetSTL) IntSetImp(maxelements int, maxval int) {
	// do nothing
}

func (set *intSetSTL) Insert(t int) {
	set.S.Add(t)
}

func (set *intSetSTL) Size() int {
	return set.S.Cardinality()
}

func (set *intSetSTL) Report() []int {
	result := make([]int, set.Size())
	i := 0
	for val := range set.S.Iter() {
		result[i] = val.(int)
		i++
	}
	sort.QuickSort(result)
	return result
}

func (set *intSetSTL) Has(t int) bool {
	return set.S.Contains(t)
}

func (set *intSetSTL) Remove(t int) {
	set.S.Remove(t)
}
