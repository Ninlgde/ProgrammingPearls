package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const NN uint = 10000000

func main() {
	var k int = 8000000
	rand.Seed(time.Now().UnixNano())

	a := make([]int32, NN)
	for i := 0; i < int(NN); i++ {
		a[i] = int32(i)
	}

	for i := 0; i < k; i++ {
		var idx int
		idx = int(rand.Int31n(int32(int(NN)-i)) + int32(i))
		a[i], a[idx] = a[idx], a[i]
	}

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
