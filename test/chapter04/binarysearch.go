package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/search"
	"sync"
)

func main() {

	var a = []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 13}

	var t = []int{5, 2, 6, 13, 12, 99}

	var ch = make(chan []int, len(t))
	var wg = new(sync.WaitGroup)

	for i := 0; i < len(t); i++ {
		go goandfind(a, t[i], ch, wg)
		wg.Add(1)
	}

	wg.Wait()
	// 关闭通道
	close(ch)

	for v := range ch {
		fmt.Println("find ", v[0], " at ", v[1])
	}

}

func goandfind(a []int, t int, ch chan []int, wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			r := []int{t, -1}
			ch <- r
			//wg.Done()
			fmt.Println("error catched", err)
		}
	}()
	defer wg.Done()
	r := []int{t, search.BinarySearch(a, t)}
	ch <- r
}
