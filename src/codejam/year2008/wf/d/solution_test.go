package main

import "testing"

func runSample(t *testing.T, N, M int, G []string, expect int) {
	g := make([][]byte, N)
	for i := 0; i < N; i++ {
		g[i] = []byte(G[i])
	}
	res := solve(N, M, g)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, M, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 2, 2
	G := []string{
		"T.",
		"T#",
	}
	expect := 2
	runSample(t, N, M, G, expect)
}

func TestSample2(t *testing.T) {
	N, M := 4, 4
	G := []string{
		"T##.",
		"##.#",
		".#T#",
		"####",
	}
	expect := 24
	runSample(t, N, M, G, expect)
}

func TestSample3(t *testing.T) {
	N, M := 4, 4
	G := []string{
		"T##.",
		"##.#",
		".#T#",
		"T###",
	}
	expect := 24
	runSample(t, N, M, G, expect)
	//  [[0 1 2 -1] [1 2 -1 2] [-1 1 4 1] [3 2 1 2]]
	// [[0 1 2 -1] [1 2 -1 2] [-1 3 4 1] [3 2 1 2]]
	// [[0 1 2 -1] [1 2 -1 2] [-1 3 4 1] [5 4 1 2]]
}

func TestSample4(t *testing.T) {
	N, M := 5, 5
	G := []string{
		"T#T.#",
		"..#.#",
		"#.###",
		"###.#",
		"T###T",
	}
	expect := 49
	runSample(t, N, M, G, expect)
}
