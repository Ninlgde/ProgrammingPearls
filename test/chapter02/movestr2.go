package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/ppstring"
)

func main() {
	a := "abcdefghi"
	fmt.Println("before: ", a)

	b := []uint8(a)

	ppstring.Move2(b, 4)

	fmt.Println("after: ", string(b))
}
