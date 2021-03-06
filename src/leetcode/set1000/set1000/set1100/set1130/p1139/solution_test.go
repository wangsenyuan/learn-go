package p1139

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := largest1BorderedSquare(grid)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
	}
	runSample(t, grid, 9)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 1, 0, 0},
	}
	runSample(t, grid, 1)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{0, 0, 0},
	}
	runSample(t, grid, 0)
}
