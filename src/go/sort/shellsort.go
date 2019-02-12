package sort

func ShellSort(x []int) []int {
	n := len(x)
	h := 1
	for ; h < n; h = 3*h + 1 {
	}
	for {
		h /= 3
		if h < 1 {
			break
		}
		for i := h; i < n; i++ {
			for j := i; j >= h; j -= h {
				if x[j-h] < x[j] {
					break
				}
				x[j-h], x[j] = x[j], x[j-h]
			}
		}
	}
	return x
}
