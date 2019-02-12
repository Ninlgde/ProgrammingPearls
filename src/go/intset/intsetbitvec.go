package intset

import "github.com/Ninlgde/ProgrammingPearls/src/go/bitmap"

type intSetBitVec struct {
	n    int
	hi   uint
	bits []int32
}

func newIntSetBitVec(maxelements int, maxval int) intSetBitVec {
	set := intSetBitVec{0, uint(maxval), nil}
	set.IntSetImp(maxelements, maxval)
	return set
}

func (set *intSetBitVec) IntSetImp(maxelements int, maxval int) {
	set.bits = make([]int32, set.hi/bitmap.BITSPERWORD)
	for i := 0; i < maxval; i++ {
		bitmap.Clr(set.bits, uint(i))
	}
}

func (set *intSetBitVec) Insert(t int) {
	if bitmap.Test(set.bits, uint(t)) != 0 {
		return
	}
	bitmap.Set(set.bits, uint(t))
	set.n++
}

func (set *intSetBitVec) Size() int {
	return set.n
}

func (set *intSetBitVec) Report() []int {
	result := make([]int, set.Size())
	j := 0
	var i uint
	for i = 0; i <= set.hi; i++ {
		if bitmap.Test(set.bits, i) != 0 {
			result[j] = int(i)
			j++
		}
	}
	return result
}

func (set *intSetBitVec) Has(t int) bool {
	return bitmap.Test(set.bits, uint(t)) != 0
}

func (set *intSetBitVec) Remove(t int) {
	bitmap.Clr(set.bits, uint(t))
}
