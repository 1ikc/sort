package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	r1 := Raw([]int{})
	r1Expected := r1
	if !Diff(BubbleSort(r1), r1Expected) {
		t.Fatalf("[fail] BubbleSort: %v", r1)
	}
	t.Logf("[pass] BubbleSort: %v", r1)

	r2 := Raw([]int{1, 5, 7, 8, 3, 4, 6})
	r2Expected := Raw([]int{1, 3, 4, 5, 6, 7, 8})
	if !Diff(BubbleSort(r2), r2Expected) {
		t.Fatalf("[fail] BubbleSort: %v", r2)
	}
	t.Logf("[pass] BubbleSort: %v", r2)

	r3 := Raw([]int{1, 5, 7, 8, 3, 4})
	r3Expected := Raw([]int{1, 3, 4, 5, 7, 8})
	if !Diff(BubbleSort(r3), r3Expected) {
		t.Fatalf("[fail] BubbleSort: %v", r3)
	}
	t.Logf("[pass] BubbleSort: %v", r3)
}