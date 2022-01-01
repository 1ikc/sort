package sort

func BubbleSort(raw Raw) Raw {
	n := len(raw)
	if n == 0 {
		return raw
	}

	// 记录本轮次是否有数据交换
	flag := true

	for i := 0; i < n && flag; i++ {
		flag = false
		for j := n - 1; j > i; j-- {
			if raw[j - 1] > raw[j] {
				raw.Swap(j - 1, j)
				flag = true
			}
		}
	}

	return raw
}