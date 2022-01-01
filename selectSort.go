package sort

func SelectSort(raw Raw) Raw {
	n := len(raw)

	for i := 0; i < n; i++ {
		min := i

		for j := i + 1; j < n; j++ {
			if raw[j] < raw[min] {
				min = j
			}
		}

		if min != i {
			raw.Swap(i, min)
		}
	}

	return raw
}