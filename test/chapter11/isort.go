package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/sort"
	"github.com/Ninlgde/ProgrammingPearls/src/go/utils"
)

func main() {

	slice := utils.GenerateSlice(100000)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	sort.ISort3(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
