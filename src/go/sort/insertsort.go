package sort

func ISort1(x []int) {
	for i := 0; i < len(x); i++ {
		for j := i; j > 0 && x[j-1] > x[j]; j-- {
			x[j-1], x[j] = x[j], x[j-1]
		}
	}
}

func ISort3(x []int) {
	for i := 0; i < len(x); i++ {
		t := x[i]
		j := i
		for ; j > 0 && x[j-1] > t; j-- {
			x[j] = x[j-1]
		}
		x[j] = t
	}
}
