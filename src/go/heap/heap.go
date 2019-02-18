package heap

type Heap interface {
	Create(cap int) bool
	Insert(t int)
	Delete() int
}

func NewHeap(cap int, isMax bool) Heap {
	if isMax {
		return newMaxHeap(cap)
	} else {
		return newMinHeap(cap)
	}
}
