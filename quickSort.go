package sort

import "math/rand"

func QuickSort(raw Raw) Raw {
	var quickSort func(raw Raw, i, j int)
	quickSort = func(raw Raw, i, j int) {
		start, end := i, j
		if start >= end {
			return
		}

		// 随机选择中点 将中点移到当前区间第一个位置
		middle := rand.Intn(end - start + 1) + start
		raw.Swap(i, middle)

		pivot := raw[i]
		for start < end {
			// 定位 比中点值小的值
			for start < end && raw[end] >= pivot {
				end--
			}
			// 定位 比中点值大的值
			for start < end && raw[start] <= pivot {
				start++
			}
			raw.Swap(start, end)
 		}

 		// 恢复中点原来的位置
 		middle = start
		raw.Swap(i, middle)

		quickSort(raw, i, middle - 1)
		quickSort(raw, middle + 1, j)
	}

	quickSort(raw, 0, len(raw) - 1)
	return raw
}