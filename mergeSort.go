package sort

func MergeSort(raw Raw) Raw {
	n := len(raw)
	if n < 2 {
		return raw
	}

	middle := n / 2
	return merge(MergeSort(raw[:middle]), MergeSort(raw[middle:]))
}

func merge(r1, r2 Raw) Raw {
	ret := Raw{}
	for len(r1) > 0 && len(r2) > 0 {
		if r1[0] > r2[0] {
			ret = append(ret, r2[0])
			r2 = r2[1:]
		} else {
			ret = append(ret, r1[0])
			r1 = r1[1:]
		}
	}

	if len(r1) > 0 {
		ret = append(ret, r1...)
	}

	if len(r2) > 0 {
		ret = append(ret, r2...)
	}

	return ret
}