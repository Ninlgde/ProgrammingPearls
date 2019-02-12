package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/intset"
)

func main() {
	set := intset.NewIntSetBST(10, 1000)

	set.Insert(1)
	set.Insert(2)
	set.Insert(3)
	set.Insert(4)

	fmt.Println(set.Report())
	fmt.Println(set.Size())

	set.Insert(1)
	set.Insert(2)
	set.Insert(3)
	set.Insert(4)

	fmt.Println(set.Report())
	fmt.Println(set.Size())

	fmt.Println(set.Has(3))
	fmt.Println(set.Has(5))

	set.Remove(3)
	fmt.Println(set.Report())
	fmt.Println(set.Size())
}
