package sort

import "math"

func ShellSort(raw Raw) Raw {
	gap := 1
	n := len(raw)

	// 确定增量区间大小
	for gap < n {
		gap = gap * 3 + 1
	}

	for gap > 0 {
		for i := gap; i < n; i++ {
			tmp := raw[i]
			j := i - gap
			for ; j >= 0 && raw[j] > tmp; j -= gap {
				raw.Swap(j + gap, j)
			}
			raw[j + gap] = tmp
		}
		gap = int(math.Floor(float64(gap / 3)))
	}

	return raw
}