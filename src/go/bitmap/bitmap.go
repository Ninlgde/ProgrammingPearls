package bitmap

const BITSPERWORD uint = 32
const SHIFT uint = 5
const MASK uint = 0x1f

func Set(a []int32, i uint) {
	a[i>>SHIFT] |= (1 << (i & MASK))
}

func Clr(a []int32, i uint) {
	a[i>>SHIFT] &= ^(1 << (i & MASK))
}

func Test(a []int32, i uint) int32 {
	return a[i>>SHIFT] & (1 << (i & MASK))
}
