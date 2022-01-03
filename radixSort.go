package sort

func RadixSort(raw Raw) Raw {
	n := len(raw)

	getMaxDigits := func (raw Raw) int {
		maxBit := 1
		key := 10
		for _, v := range raw {
			for v >= key {
				key *= 10
				maxBit++
			}
		}
		return maxBit
	}

	for i := range raw {
		raw[i] += 50001
	}
	maxBit := getMaxDigits(raw)

	radix := 1
	tmp := make(Raw, n)
	var count [10]int
	for i := 0; i < maxBit; i++ {
		for j := 0; j < 10; j++ {
			count[j] = 0
		}
		for j := 0; j < n; j++ {
			k := (raw[j] / radix) % 10
			count[k]++
		}
		for j := 1; j < 10; j++ {
			count[j] += count[j - 1]
		}
		for j := n - 1; j >= 0; j-- {
			k := (raw[j] / radix) % 10
			tmp[count[k] - 1] = raw[j]
			count[k]--
		}
		for j := 0; j < n; j++ {
			raw[j] = tmp[j]
		}
		radix *= 10
	}

	for i := range raw {
		raw[i] -= 50001
	}

	return raw
}