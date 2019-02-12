package sort

func SelectSort(x []int) []int {
	for i := 0; i < len(x)-1; i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] < x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	return x
}
