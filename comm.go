package sort

type Raw []int

func (r Raw) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func Diff(r1, r2 Raw) bool {
	n := len(r1)
	if n != len(r2) {
		return false
	}

	for i := 0; i < n; i++ {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}