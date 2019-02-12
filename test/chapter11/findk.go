package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/search"
	"github.com/Ninlgde/ProgrammingPearls/src/go/utils"
)

func main() {
	slice := utils.GenerateSlice(10)
	//slice := []int {143, 382, 19, 164, 673, 508, -22, 296, -351, 442}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	search.SearchK(slice, 8)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
	fmt.Println(slice[8])
}
