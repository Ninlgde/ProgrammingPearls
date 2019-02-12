package intset

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

type intSetBST struct {
	n        int
	root     *treeNode
	sentinel *treeNode
}

func newIntSetBST(maxelements int, maxval int) intSetBST {
	set := intSetBST{0, nil, nil}
	set.IntSetImp(maxelements, maxval)
	return set
}

func (set *intSetBST) IntSetImp(maxelements int, maxval int) {
	set.sentinel = &treeNode{maxval, nil, nil}
	set.root = set.sentinel
}

func (set *intSetBST) Insert(t int) {
	set.root = trinsert(set, set.root, t)
}

func (set *intSetBST) Size() int {
	return set.n
}

func (set *intSetBST) Report() []int {
	result := make([]int, set.Size())
	var vn *int
	*vn = 0
	traverse(set, set.root, result, vn)
	return nil
}

func (set *intSetBST) Has(t int) bool {
	p := set.root
	for p != nil && p.val != t {
		if p.val > t {
			p = p.left
		} else if p.val < t {
			p = p.right
		}
	}
	return p != nil && p.val == t
}

func (set *intSetBST) Remove(t int) {
	//todo: bst delete
}

func trinsert(set *intSetBST, p *treeNode, t int) *treeNode {
	if p == set.sentinel {
		p = &treeNode{t, nil, nil}
		set.n++
	} else if t < p.val {
		p.left = trinsert(set, p.left, t)
	} else if t > p.val {
		p.right = trinsert(set, p.right, t)
	}
	return p
}

func traverse(set *intSetBST, p *treeNode, v []int, vn *int) {
	if p == set.sentinel {
		return
	}
	traverse(set, p.left, v, vn)
	v[*vn] = p.val
	*vn++
	traverse(set, p.right, v, vn)
}
