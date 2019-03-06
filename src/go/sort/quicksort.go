package sort

import (
	"math/rand"
	"time"
)

func QSort1(x []int) []int {
	if len(x) < 2 {
		return x
	}
	l := 0
	u := len(x) - 1
	m := 0
	for i := l + 1; i <= u; i++ {
		if x[i] < x[l] {
			m++
			x[m], x[i] = x[i], x[m]
		}
	}
	x[m], x[l] = x[l], x[m]
	QSort1(x[:m])
	QSort1(x[m+1:])

	return x
}

func QSort2(x []int) []int {
	if len(x) < 2 {
		return x
	}
	l := 0
	u := len(x) - 1
	m := u + 1
	for i := u; i >= l; i-- {
		if x[i] >= x[l] {
			m--
			x[m], x[i] = x[i], x[m]
		}
	}
	QSort2(x[:m])
	QSort2(x[m+1:])

	return x
}

func QSort3(x []int) []int {
	if len(x) < 2 {
		return x
	}
	l := 0
	u := len(x) - 1
	t := x[l]
	i := l
	j := u + 1
	for {
		for i++; i <= u && x[i] < t; i++ {
		}
		for j--; x[j] > t; j-- {
		}
		if i > j {
			break
		}
		x[i], x[j] = x[j], x[i]
	}
	x[l], x[j] = x[j], x[l]
	QSort3(x[:j])
	QSort3(x[j+1:])
	return x
}

func QSort4(x []int) []int {
	if len(x) < 50 {
		return x
	}
	l := 0
	u := len(x) - 1
	r := rand.Int() % len(x)
	x[r], x[l] = x[l], x[r]
	t := x[l]
	i := l
	j := u + 1
	for {
		for i++; i <= u && x[i] < t; i++ {
		}
		for j--; x[j] > t; j-- {
		}
		if i > j {
			break
		}
		x[i], x[j] = x[j], x[i]
	}
	x[l], x[j] = x[j], x[l]
	QSort4(x[:j])
	QSort4(x[j+1:])
	return x
}

func QuickSort(x []int) []int {
	rand.Seed(time.Now().UnixNano())
	QSort4(x)
	ISort3(x)
	return x
}
