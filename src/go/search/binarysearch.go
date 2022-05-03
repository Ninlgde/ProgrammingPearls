package search

import (
	"fmt"
)

func BinarySearch(a []int, target int) int {
	begin := 0
	end := len(a) - 1

	for begin <= end {
		mid := begin + (end-begin)/2
		if a[mid] > target {
			end = mid - 1
		} else if a[mid] < target {
			begin = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func FastBinarySearch(a []int, target int) int {
	begin := -1
	end := len(a)

	for begin+1 != end {
		mid := begin + (end-begin)/2
		if a[mid] < target {
			begin = mid
		} else {
			end = mid
		}
	}
	if end < len(a) && a[end] == target {
		return end
	}
	return -1
}

func FirstBinarySearch(a []int, target int) int {
	begin := 0
	end := len(a) - 1

	first := -1
	for begin <= end {
		mid := begin + (end-begin)/2
		if a[mid] > target {
			end = mid - 1
		} else if a[mid] < target {
			begin = mid + 1
		} else {
			first = mid
			end = mid - 1
		}
	}
	return first
}

func RecursionBinarySearch(a []int, target int, begin int, end int) int {

	if begin > end {
		return -1
	}

	mid := begin + (end-begin)/2
	if a[mid] > target {
		return RecursionBinarySearch(a, target, begin, mid-1)
	} else if a[mid] < target {
		return RecursionBinarySearch(a, target, mid+1, end)
	} else {
		return mid
	}
	return -1
}

func mastbe(a []int, l int, u int, t int) {
	if a[l] > t || a[u] < t {
		panic(fmt.Errorf("MUST BE ERROR", t))
	}
}

func cantbe(a []int, l int, u int, t int) {
	if a[l] < t && a[u] > t {
		panic(fmt.Errorf("CANT BE ERROR", t))
	}
}
