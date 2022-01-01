package sort

func InsertSort(raw Raw) Raw {
	n := len(raw)

	for i := 1; i < n; i++ {
		tmp := raw[i]
		j := i
		for ; j > 0 && tmp < raw[j - 1]; j-- {
			raw.Swap(j, j - 1)
		}
		if j != i {
			raw[j] = tmp
		}
	}

	return raw
}