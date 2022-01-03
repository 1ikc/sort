package sort

func CountSort(raw Raw) Raw {
	// 待排序数组存在负数，可以在排序前统一增加一个大的正数，排序后再统一扣减
	// 假设待排序数组最大值是　5*(10^4), 大正数定义为50001
	bucketLen := raw.GetMaxValue() + 1 + 50001
	bucket := make(Raw, bucketLen)
	sortIdx := 0

	for _, v := range raw {
		v += 50001
		bucket[v]++
	}

	for i := 0; i < bucketLen; i++ {
		for bucket[i] > 0 {
			raw[sortIdx] = i - 50001
			sortIdx++
			bucket[i]--
		}
	}

	return raw
}