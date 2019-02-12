package intset

type listNode struct {
	val  int
	next *listNode
}

type intSetList struct {
	n        int
	head     *listNode
	sentinel *listNode
}

func newIntSetList(maxelements int, maxval int) intSetList {
	set := intSetList{0, nil, nil}
	set.IntSetImp(maxelements, maxval)
	return set
}

func (set *intSetList) IntSetImp(maxelements int, maxval int) {
	set.sentinel = &listNode{maxval, nil}
	set.head = set.sentinel
}

func (set *intSetList) Insert(t int) {
	set.head = lrinsert(set, set.head, t)
}

func (set *intSetList) Size() int {
	return set.n
}

func (set *intSetList) Report() []int {
	result := make([]int, set.Size())
	i := 0
	for p := set.head; p != set.sentinel; p = p.next {
		result[i] = p.val
		i++
	}
	return result
}

func (set *intSetList) Has(t int) bool {
	for p := set.head; p != set.sentinel; p = p.next {
		if p.val == t {
			return true
		}
	}
	return false
}

func (set *intSetList) Remove(t int) {
	p := set.head
	var pre *listNode
	for ; p != set.sentinel; p = p.next {
		if p.val == t {
			pre.next = p.next
		}
		pre = p
	}
}

func lrinsert(set *intSetList, p *listNode, t int) *listNode {
	if p.val < t {
		p.next = lrinsert(set, p.next, t)
	} else if p.val > t {
		p = &listNode{t, p}
		set.n++
	}
	return p
}
