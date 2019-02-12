package search

func Search1a(x []int, t int) int {
	for i, v := range x {
		if v == t {
			return i
		}
	}
	return -1
}

func Search1b(x []int, t int) int {
	for i := 0; i < len(x); i++ {
		if x[i] == t {
			return i
		}
	}
	return -1
}

func Search1c(x []int, t int) int {
	n := len(x)
	for i := 0; i < n; i++ {
		if x[i] == t {
			return i
		}
	}
	return -1
}

func Search2(x []int, t int) int {
	n := len(x) - 1
	hold := x[n]
	x[n] = t
	i := 0
	// 少了for循环的比较所以快的一匹
	for ; ; i++ {
		if x[i] == t {
			break
		}
	}
	x[n] = hold
	if i == n {
		if hold == t {
			return n
		}
		return -1
	}
	return i
}

func Search3(x []int, t int) int {
	n := len(x) - 1
	hold := x[n]
	x[n] = t
	i := 0
	// 少了for循环的比较所以快的一匹
	for ; ; i += 8 {
		if x[i] == t {
			break
		}
		if x[i+1] == t {
			i += 1
			break
		}
		if x[i+2] == t {
			i += 2
			break
		}
		if x[i+3] == t {
			i += 3
			break
		}
		if x[i+4] == t {
			i += 4
			break
		}
		if x[i+5] == t {
			i += 5
			break
		}
		if x[i+6] == t {
			i += 6
			break
		}
		if x[i+7] == t {
			i += 7
			break
		}
	}
	x[n] = hold
	if i == n {
		if hold == t {
			return n
		}
		return -1
	}
	return i
}
