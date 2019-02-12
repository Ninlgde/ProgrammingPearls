package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/ppstring"
)

func main() {
	a := "abcdefgh"
	fmt.Println("before: ", a)

	b := []uint8(a)

	ppstring.Move1(b, 3)

	fmt.Println("after: ", string(b))

	//b1 := b[9-4:]
	//move(b1, 3)
	//fmt.Println("after: ", string(b1))
}

func move(s []uint8, i int) []uint8 {
	var tmp uint8
	var last int
	for j := 0; j < i; j++ {
		tmp = s[j]
		for k := 0; k < 1+len(s)/i; k++ {
			a := j + i*k
			if j+i*(k+1) < len(s) {
				s[a] = s[j+i*(k+1)]
				last = j + i*(k+1)
			}
		}
		s[last] = tmp
	}
	return s
}
