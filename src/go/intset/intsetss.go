package intset

import (
	"github.com/Ninlgde/ProgrammingPearls/src/go/simpleset"
	"github.com/Ninlgde/ProgrammingPearls/src/go/sort"
)

type intSetSimpleSet struct {
	S simpleset.Set
}

func newIntSetSS(maxelements int, maxval int) intSetSimpleSet {
	set := intSetSimpleSet{S: simpleset.NewSet()}
	set.IntSetImp(maxelements, maxval)
	return set
}

func (set *intSetSimpleSet) IntSetImp(maxelements int, maxval int) {
	// do nothing
}

func (set *intSetSimpleSet) Insert(t int) {
	set.S.Add(t)
}

func (set *intSetSimpleSet) Size() int {
	return set.S.Cardinality()
}

func (set *intSetSimpleSet) Report() []int {
	result := make([]int, set.Size())
	i := 0
	for val := range set.S.Iter() {
		result[i] = val.(int)
		i++
	}
	sort.QuickSort(result)
	return result
}

func (set *intSetSimpleSet) Has(t int) bool {
	return set.S.Contains(t)
}

func (set *intSetSimpleSet) Remove(t int) {
	set.S.Remove(t)
}
