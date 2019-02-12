package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/search"
)

func main() {

	var a = []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 13}

	//fmt.Println(binarysearch(a, 5))
	//fmt.Println(binarysearch(a, 2))
	//fmt.Println(binarysearch(a, 6))
	//fmt.Println(binarysearch(a, 13))
	//fmt.Println(binarysearch(a, 12))

	var t = []int{5, 2, 6, 13, 12, 99}
	for i := 0; i < len(t); i++ {
		goandfind4(a, t[i])
	}

}

func goandfind4(a []int, t int) {
	begin := 0
	end := len(a) - 1
	fmt.Println("find ", t, " at ", search.RecursionBinarySearch(a, t, begin, end))
}
