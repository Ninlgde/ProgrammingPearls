package intset

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

var niltreenode = &treeNode{}

type intSetBST struct {
	n    int
	root *treeNode
}

func newIntSetBST(maxelements int, maxval int) intSetBST {
	set := intSetBST{0, nil}
	set.IntSetImp(maxelements, maxval)
	return set
}

func (set *intSetBST) IntSetImp(maxelements int, maxval int) {
	set.root = niltreenode
}

func (set *intSetBST) Insert(t int) {
	set.root = trinsert(set, set.root, t)
}

func (set *intSetBST) Size() int {
	return set.n
}

func (set *intSetBST) Report() []int {
	result := make([]int, set.Size())
	ch1 := traverse(set, set.root)
	i := 0
	for {
		val, ok := <-ch1
		if ok {
			result[i] = val
			i++
		} else {
			break
		}
	}
	return result
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
	if p == niltreenode {
		p = &treeNode{t, niltreenode, niltreenode}
		set.n++
	} else if t < p.val {
		p.left = trinsert(set, p.left, t)
	} else if t > p.val {
		p.right = trinsert(set, p.right, t)
	}
	return p
}

func walk(set *intSetBST, p *treeNode, ch chan int) {
	if p == niltreenode {
		return
	}
	walk(set, p.left, ch)
	ch <- p.val
	walk(set, p.right, ch)
}

func traverse(set *intSetBST, p *treeNode) <-chan int {
	ch := make(chan int)
	go func() {
		walk(set, p, ch)
		close(ch)
	}()
	return ch
}
