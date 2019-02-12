package intset

type IntSet interface {
	IntSetImp(maxelements int, maxval int)
	Insert(t int)
	Size() int
	Report() []int
	Has(t int) bool
	Remove(t int)
}

func NewIntSet(maxelements int, maxval int) IntSet {
	set := newIntSetSTL(maxelements, maxval)
	return &set
}

func NewIntSetArray(maxelements int, maxval int) IntSet {
	set := newIntSetArray(maxelements, maxval)
	return &set
}

func NewIntSetBins(maxelements int, maxval int) IntSet {
	set := newIntSetBins(maxelements, maxval)
	return &set
}

func NewIntSetBitVec(maxelements int, maxval int) IntSet {
	set := newIntSetBitVec(maxelements, maxval)
	return &set
}

func NewIntSetBST(maxelements int, maxval int) IntSet {
	set := newIntSetBST(maxelements, maxval)
	return &set
}

func NewIntSetList(maxelements int, maxval int) IntSet {
	set := newIntSetList(maxelements, maxval)
	return &set
}
