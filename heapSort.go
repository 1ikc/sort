package sort

func HeapSort(raw Raw) Raw {
	n := len(raw)

	// 堆化
	var heapify func(raw Raw, i, length int)
	heapify = func(raw Raw, i, length int) {
		// 定位 当前结点的左右子结点 堆是一种特殊的完全二叉树
		left := 2 * i + 1
		right := 2 * i + 2
		largest := i

		// 判断当前结点是否需要下沉
		if left < length && raw[left] > raw[largest] {
			largest = left
		}
		if right < length && raw[right] > raw[largest] {
			largest = right
		}
		if largest != i {
			// 当前结点下沉 子结点上浮 向下继续递归
			raw.Swap(i, largest)
			heapify(raw, largest, length)
		}
	}

	// 构建大顶堆
	for i := n / 2; i >= 0; i-- {
		heapify(raw, i, n)
	}

	// 重建堆
	for i := n - 1; i >= 0; i-- {
		raw.Swap(i, 0)
		n -= 1
		heapify(raw, 0, n)
	}

	return raw
}