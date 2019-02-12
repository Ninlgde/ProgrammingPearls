package main

import (
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/utils"
	"math/rand"
	"time"
)

func main() {
	//set := intset.NewIntSetBST(10, 1000)
	//
	//set.Insert(1)
	//set.Insert(2)
	//set.Insert(3)
	//set.Insert(4)
	//
	//fmt.Println(set.Report())
	//fmt.Println(set.Size())
	//
	//set.Insert(1)
	//set.Insert(2)
	//set.Insert(3)
	//set.Insert(4)
	//
	//fmt.Println(set.Report())
	//fmt.Println(set.Size())
	//
	//fmt.Println(set.Has(3))
	//fmt.Println(set.Has(5))
	//
	//set.Remove(3)
	//fmt.Println(set.Report())
	//fmt.Println(set.Size())

	t1 := time.Now()
	for i := 0; i < 8000000; i++ {
		utils.BigRand()
		//fmt.Println(time.Now().UnixNano(), " : ",utils.BigRand())
	}
	elapsed := time.Since(t1)
	fmt.Println("search1a elapsed: ", elapsed)

	t2 := time.Now()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8000000; i++ {
		rand.Int()
	}
	elapsed2 := time.Since(t2)
	fmt.Println("search1a elapsed: ", elapsed2)
}
