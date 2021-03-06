package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{3, 2, 7, 6}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{3, 2, 6}
	expect := 1
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 1
	A := []int{7}
	expect := -1
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 7
	A := []int{4, 9, 2, 1, 18, 3, 0}
	expect := 0
	runSample(t, n, A, expect)
}
