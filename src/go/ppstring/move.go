package ppstring

import "github.com/Ninlgde/ProgrammingPearls/src/go/utils"

func Move1(s []uint8, rotdist int) {
	for i := 0; i < utils.GCD(rotdist, len(s)); i++ {
		tmp := s[i]
		j := i
		for {
			k := j + rotdist
			if k >= len(s) {
				k -= len(s)
			}
			if k == i {
				break
			}
			s[j] = s[k]
			j = k
		}
		s[j] = tmp
	}
}

func Move2(s []uint8, rotdist int) []uint8 {
	i := rotdist
	p := rotdist
	j := len(s) - p
	for i != j {
		if i > j {
			swap(s, p-i, p, j)
			i -= j
		} else {
			swap(s, p-i, p+j-i, i)
			j -= i
		}
	}
	swap(s, p-i, p, i)
	return s
}

func swap(s []uint8, a int, b int, m int) {
	for i := 0; i < m; i++ {
		s[a+i], s[b+i] = s[b+i], s[a+i]
	}
}

func Move3(s []uint8, rotdist int) []uint8 {
	reverse(s, 0, rotdist-1)
	reverse(s, rotdist, len(s)-1)
	reverse(s, 0, len(s)-1)
	return s
}

func Move4(s []uint8, a int, b int) {
	reverse(s, 0, a-1)
	reverse(s, a, len(s)-b-1)
	reverse(s, len(s)-b, len(s)-1)
	reverse(s, 0, len(s)-1)
}

func reverse(s []uint8, a int, b int) {
	for i, j := a, b; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
