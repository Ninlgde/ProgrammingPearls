package utils

import (
	"math/rand"
	"time"
)

func GCD(i int, j int) int {
	for i != j {
		if i > j {
			i -= j
		} else {
			j -= i
		}
	}
	return i
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func BigRandSeed(seed int64) int {
	rand.Seed(seed)
	return rand.Int()
}

func BigRand() int {
	return BigRandSeed(time.Now().UnixNano())
}

func RandInt(l int, u int) int {
	return BigRand()%(u-l+1) + l
}
