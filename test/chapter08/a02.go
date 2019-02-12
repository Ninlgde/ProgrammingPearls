package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/search"
)

func main() {
	var x = []int{31, -41, 59, 26, -53, 58, 97, -93, -23, 84}

	fmt.Println(search.MaxRange2(x))
}
