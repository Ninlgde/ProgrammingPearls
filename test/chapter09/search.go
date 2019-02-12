package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/search"
	"time"
)

func main() {
	N := 1000000000
	x := make([]int, N)
	for i := range x {
		x[i] = i
	}

	t1 := time.Now()
	r1 := search.Search1a(x, N-3)
	elapsed := time.Since(t1)
	fmt.Println("search1a elapsed: ", elapsed, "result: ", r1)

	t2 := time.Now()
	r2 := search.Search1b(x, N-3)
	elapsed = time.Since(t2)
	fmt.Println("search1b elapsed: ", elapsed, "result: ", r2)

	t3 := time.Now()
	r3 := search.Search1c(x, N-3)
	elapsed = time.Since(t3)
	fmt.Println("search1c elapsed: ", elapsed, "result: ", r3)

	t4 := time.Now()
	r4 := search.Search2(x, N-3)
	elapsed = time.Since(t4)
	fmt.Println("search2 elapsed: ", elapsed, "result: ", r4)

	t5 := time.Now()
	r5 := search.Search3(x, N-3)
	elapsed = time.Since(t5)
	fmt.Println("search3 elapsed: ", elapsed, "result: ", r5)
}
