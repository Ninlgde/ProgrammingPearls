package main

import (
	"bufio"
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/sampling"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const NN int = 10000000

func main() {
	var k int = 8000000
	rand.Seed(time.Now().UnixNano())

	t1 := time.Now()
	a := sampling.GenKnuth(k, NN)
	elapsed := time.Since(t1)
	fmt.Println("GenKnuth elapsed: ", elapsed, "result len: ", len(a))

	t1 = time.Now()
	a = sampling.GenKnuthFaster(k, NN)
	elapsed = time.Since(t1)
	fmt.Println("GenKnuthFaster elapsed: ", elapsed, "result len: ", len(a))

	t1 = time.Now()
	a = sampling.GenFloyd(k, NN)
	elapsed = time.Since(t1)
	fmt.Println("GenFloyd elapsed: ", elapsed, "result len: ", len(a))

	t1 = time.Now()
	a = sampling.GenFloydFaster(k, NN)
	elapsed = time.Since(t1)
	fmt.Println("GenFloydFaster elapsed: ", elapsed, "result len: ", len(a))

	f, err := os.Create("data.dat")
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)

	for i := 0; i < k; i++ {
		//fmt.Print(strconv.Itoa(int(a[i])), "\n")
		w.WriteString(strconv.Itoa(int(a[i])))
		w.WriteString("\n")
	}

	w.Flush()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
