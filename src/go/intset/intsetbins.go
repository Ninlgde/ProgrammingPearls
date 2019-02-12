package intset

type binNode struct {
	val  int
	next *binNode
}

type intSetBins struct {
	n        int
	bins     int
	maxval   int
	bin      []*binNode
	sentinel *binNode
}

func newIntSetBins(maxelements int, maxval int) intSetBins {
	set := intSetBins{0, maxelements, maxval, nil, nil}
	set.IntSetImp(maxelements, maxval)
	return set
}

func (set *intSetBins) IntSetImp(maxelements int, maxval int) {
	set.bin = make([]*binNode, maxelements)
	set.sentinel = &binNode{maxval, nil}
	for i := 0; i < maxelements; i++ {
		set.bin[i] = set.sentinel
	}
}

func (set *intSetBins) Insert(t int) {
	i := t / (1 + set.maxval/set.bins)
	set.bin[i] = brinsert(set, set.bin[i], t)
}

func (set *intSetBins) Size() int {
	return set.n
}

func (set *intSetBins) Report() []int {
	result := make([]int, set.Size())
	j := 0
	for i := 0; i < set.bins; i++ {
		for p := set.bin[i]; p != set.sentinel; p = p.next {
			result[j] = p.val
			j++
		}
	}
	return result
}

func (set *intSetBins) Has(t int) bool {
	i := t / (1 + set.maxval/set.bins)
	for p := set.bin[i]; p != set.sentinel; p = p.next {
		if p.val == t {
			return true
		}
	}
	return false
}

func (set *intSetBins) Remove(t int) {
	i := t / (1 + set.maxval/set.bins)
	p := set.bin[i]
	var pre *binNode
	for ; p != set.sentinel; p = p.next {
		if p.val == t {
			if p == set.bin[i] {
				set.bin[i] = p.next
			} else {
				pre.next = p.next
			}
			set.n--
			break
		}
		pre = p
	}
}

func brinsert(set *intSetBins, p *binNode, t int) *binNode {
	if p.val < t {
		p.next = brinsert(set, p.next, t)
	} else if p.val > t {
		p = &binNode{t, p}
		set.n++
	}
	return p
}
