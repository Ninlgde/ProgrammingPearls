package heap

type minHeap struct {
	n   int
	cap int
	x   []int
}

func newMinHeap(cap int) *minHeap {
	h := &minHeap{0, cap, nil}
	h.Create(cap)
	return h
}

func (heap *minHeap) Create(cap int) bool {
	heap.x = make([]int, cap+1)
	return true
}

func (heap *minHeap) Insert(t int) {
	heap.n++
	heap.x[heap.n] = t
	heap.shiftup()
}

func (heap *minHeap) Delete() int {
	r := heap.x[1]
	heap.n--
	heap.x[1] = heap.x[heap.n]
	heap.shiftdown()
	return r
}

func (heap *minHeap) shiftup() {
	x := heap.x
	n := heap.n
	i := n
	for {
		if i == 1 {
			break
		}
		p := i / 2
		if x[p] <= x[i] {
			break
		}
		x[p], x[i] = x[i], x[p]
		i = p
	}
}

func (heap *minHeap) shiftdown() {
	x := heap.x
	n := heap.n
	i := 1
	for {
		c := 2 * i
		if x[c] > x[n] {
			break
		}
		if x[c+1] <= x[n] {
			if x[c+1] < x[c] {
				c++
			}
		}
		if x[i] <= x[c] {
			break
		}
		x[c], x[i] = x[i], x[c]
		i = c
	}
}
