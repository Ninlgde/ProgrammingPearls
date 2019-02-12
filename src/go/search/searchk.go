package search

import (
	"math/rand"
	"time"
)

func SearchK(x []int, k int) {
	if len(x) < 2 {
		return
	}
	l := 0
	u := len(x) - 1
	rand.Seed(time.Now().UnixNano())
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
	if j > k {
		SearchK(x[:j], k)
	} else if j < k {
		SearchK(x[j+1:], k-j) // k要减去j，因为范围被缩小了原来找第k个。。如果找后半区 就要找第k-j个
	}
}
