package main

import (
	"bufio"
	"fmt"
	"github.com/Ninlgde/ProgrammingPearls/src/go/bitmap"
	"io"
	"os"
	"strconv"
	"strings"
)

//const BITSPERWORD uint = 32
//const SHIFT uint = 5
//const MASK uint = 0x1f
const N uint = 10000000

func main() {
	//twostepmain()
	unittest()
}

func twostepmain() {
	a := make([]int32, 1+N/bitmap.BITSPERWORD/2) // only half mem

	half := int(N) / 2

	fi, err := os.Open("data.dat")
	check2(err)

	defer fi.Close()

	reader1 := bufio.NewReader(fi)

	for {
		line, err := reader1.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.Replace(line, "\n", "", -1)
		v, _ := strconv.Atoi(line)
		if v < half {
			bitmap.Set(a, uint(v))
		}
	}

	fw, err := os.Create("result2.dat")
	check2(err)

	defer fw.Close()

	w := bufio.NewWriter(fw)

	for i := 0; i < half; i++ {
		if bitmap.Test(a, uint(i)) != 0 {
			w.WriteString(strconv.Itoa(i))
			w.WriteString("\n")
		}
	}

	// 第二遍扫描文件
	reader2 := bufio.NewReader(fi)

	for {
		line, err := reader2.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.Replace(line, "\n", "", -1)
		v, _ := strconv.Atoi(line)
		if v >= half {
			bitmap.Set(a, uint(v-half))
		}
	}

	for i := 0; i < half; i++ {
		if bitmap.Test(a, uint(i)) != 0 {
			w.WriteString(strconv.Itoa(i + half))
			w.WriteString("\n")
		}
	}

	w.Flush()
}

func onestepmain() {
	a := make([]int32, 1+N/bitmap.BITSPERWORD)

	fi, err := os.Open("data.dat")
	check2(err)

	defer fi.Close()

	reader := bufio.NewReader(fi)

	//x := make([]int32, N)
	//i := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//fmt.Println(line)
		line = strings.Replace(line, "\n", "", -1)
		v, _ := strconv.Atoi(line)
		//x[i] = int32(v)
		//i++
		bitmap.Set(a, uint(v))
	}

	fw, err := os.Create("result1.dat")
	check2(err)

	defer fw.Close()

	w := bufio.NewWriter(fw)

	for i := 0; i < int(N); i++ {
		if bitmap.Test(a, uint(i)) != 0 {
			w.WriteString(strconv.Itoa(i))
			w.WriteString("\n")
		}
	}

	w.Flush()
}

func unittest() {
	var idx uint
	a := make([]int32, 1+N/bitmap.BITSPERWORD)
	idx = 88
	bitmap.Set(a, idx)
	bitmap.Set(a, 87)
	fmt.Println("\n--- test idx ---\n\n", bitmap.Test(a, idx), "\n")
	fmt.Println("\n--- test idx ---\n\n", bitmap.Test(a, 87), "\n")
	bitmap.Clr(a, idx)
	fmt.Println("\n--- test idx ---\n\n", bitmap.Test(a, idx), "\n")
	fmt.Println("\n--- test idx ---\n\n", bitmap.Test(a, 87), "\n")
}

func check2(e error) {
	if e != nil {
		panic(e)
	}
}
