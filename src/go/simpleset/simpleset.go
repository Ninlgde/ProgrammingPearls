package simpleset

// 单线程的set

type Set interface {
	Add(e interface{})
	Contains(e interface{}) bool
	Remove(e interface{})
	Cardinality() int
	Iter() <-chan interface{}
}

type mapSet struct {
	S map[interface{}]bool
}

func NewSet() Set {
	return &mapSet{make(map[interface{}]bool)}
}

func (set *mapSet) Add(e interface{}) {
	if !set.S[e] {
		set.S[e] = true
	}
}

func (set *mapSet) Contains(e interface{}) bool {
	return set.S[e]
}

func (set *mapSet) Remove(e interface{}) {
	delete(set.S, e)
}

func (set *mapSet) Cardinality() int {
	return len(set.S)
}

func (set *mapSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for elem := range set.S {
			ch <- elem
		}
		close(ch)
	}()
	return ch
}
