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

var hasSetSeed = false

func BigRandSeed(seed int64) int {
	rand.Seed(seed)
	return rand.Int()
}

func BigRand() int {
	if hasSetSeed {
		return rand.Int()
	}
	hasSetSeed = true
	return BigRandSeed(time.Now().UnixNano())
}

func RandInt(l int, u int) int {
	return BigRand()%(u-l+1) + l
}
