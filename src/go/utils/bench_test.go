package utils

import (
	"math/rand"
	"testing"
)

func randPair() (int, int) {
	rand.Seed(0)
	return rand.Int(), rand.Int()
}

func BenchmarkGCD(b *testing.B) {
	m, n := randPair()
	//fmt.Println(m, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GCD(m, n)
	}
}

func BenchmarkGCD2(b *testing.B) {
	m, n := randPair()
	//fmt.Println(m, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GCD2(m, n)
	}
}
