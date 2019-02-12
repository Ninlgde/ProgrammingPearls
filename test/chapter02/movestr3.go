package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/ppstring"
)

func main() {
	a := "abcdefgh"
	fmt.Println("before: ", a)

	b := []uint8(a)

	ppstring.Move3(b, 3)

	fmt.Println("after: ", string(b))

	a4 := "abcdefgh"
	fmt.Println("before: ", a4)

	b4 := []uint8(a4)

	ppstring.Move4(b4, 2, 2)

	fmt.Println("after: ", string(b4))
}
