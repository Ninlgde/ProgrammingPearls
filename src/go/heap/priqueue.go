package heap

type PriQueue struct {
	heap *minHeap
}

func NewPriQueue(cap int) *PriQueue {
	pq := &PriQueue{}
	pq.heap = newMinHeap(cap)
	return pq
}

func (pq *PriQueue) Insert(t int) {
	pq.heap.Insert(t)
}

func (pq *PriQueue) ExtractMin() int {
	return pq.heap.Delete()
}
