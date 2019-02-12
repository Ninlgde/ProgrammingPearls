package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/sort"
	"github.com/Ninlgde/ProgrammingPearls/src/go/utils"
)

func main() {
	slice := utils.GenerateSlice(100000)
	//slice := []int {-69, -259, 477, -232, 528, 353, -503, 343, -16, 108, -474, 353, -62, -114, 212, 484, -223, 10, -169, 16, -802, 265, -320, -205, -104, 980, -136, 395, 275, -519}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	sort.QSort4(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
