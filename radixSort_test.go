package sort

import "testing"

func TestRadixSort(t *testing.T) {
	r1 := Raw([]int{})
	r1Expected := r1
	if !Diff(RadixSort(r1), r1Expected) {
		t.Fatalf("[fail] RadixSort: %v", r1)
	}
	t.Logf("[pass] RadixSort: %v", r1)

	r2 := Raw([]int{1, 5, 7, 8, 3, 4, 6})
	r2Expected := Raw([]int{1, 3, 4, 5, 6, 7, 8})
	if !Diff(RadixSort(r2), r2Expected) {
		t.Fatalf("[fail] RadixSort: %v", r2)
	}
	t.Logf("[pass] RadixSort: %v", r2)

	r3 := Raw([]int{5, 1, 1, 2, 0, 0})
	r3Expected := Raw([]int{0, 0, 1, 1, 2, 5})
	if !Diff(RadixSort(r3), r3Expected) {
		t.Fatalf("[fail] RadixSort: %v", r3)
	}
	t.Logf("[pass] RadixSort: %v", r3)

	r4 := Raw([]int{-1, -4, 0})
	r4Expected := Raw([]int{-4, -1, 0})
	if !Diff(RadixSort(r4), r4Expected) {
		t.Fatalf("[fail] RadixSort: %v", r4)
	}
	t.Logf("[pass] RadixSort: %v", r4)
}