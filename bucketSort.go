package sort

func BucketSort(raw Raw) Raw {
	n := len(raw)
	bucket := make([]Raw, n)

	for i := range raw {
		// 待排序数组存在负数，可以在排序前统一增加一个大的正数，排序后再统一扣减
		// 假设待排序数组最大值是　5*(10^4), 大正数定义为50001
		raw[i] += 50001
	}

	max := raw.GetMaxValue()
	for _, v := range raw {
		// 分配桶
		idx := v * (n - 1) / max
		bucket[idx] = append(bucket[idx], v - 50001)
	}

	pos := 0
	for i := 0; i < n; i++ {
		bucketLen := len(bucket[i])
		if bucketLen > 0 {
			// 是否稳定，取决于桶内排序选择的排序算法是否稳定
			// 此处选择插入排序，因此是稳定排序
			InsertSort(bucket[i])
			copy(raw[pos:], bucket[i])
			pos += bucketLen
		}
	}


	return raw
}