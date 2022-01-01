package sort

import "testing"

func TestShellSort(t *testing.T) {
	r1 := Raw([]int{})
	r1Expected := r1
	if !Diff(ShellSort(r1), r1Expected) {
		t.Fatalf("[fail] ShellSort: %v", r1)
	}
	t.Logf("[pass] ShellSort: %v", r1)

	r2 := Raw([]int{1, 5, 7, 8, 3, 4, 6})
	r2Expected := Raw([]int{1, 3, 4, 5, 6, 7, 8})
	if !Diff(ShellSort(r2), r2Expected) {
		t.Fatalf("[fail] ShellSort: %v", r2)
	}
	t.Logf("[pass] ShellSort: %v", r2)

	r3 := Raw([]int{5, 1, 1, 2, 0, 0})
	r3Expected := Raw([]int{0, 0, 1, 1, 2, 5})
	if !Diff(ShellSort(r3), r3Expected) {
		t.Fatalf("[fail] ShellSort: %v", r3)
	}
	t.Logf("[pass] ShellSort: %v", r3)
}