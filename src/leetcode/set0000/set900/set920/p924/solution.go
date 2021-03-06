package p924

import "sort"

func minMalwareSpread(graph [][]int, initial []int) int {
	if len(initial) == 0 {
		return -1
	}
	n := len(graph)
	uf := NewUFSet(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	sort.Ints(initial)

	bags := make([]int, n)

	for i := 0; i < len(initial); i++ {
		x := initial[i]
		px := uf.Find(x)
		bags[px]++
	}
	cand := initial[0]
	save := -1

	for i := 0; i < len(initial); i++ {
		x := initial[i]
		p := uf.Find(x)
		if bags[p] == 1 {
			if uf.Size(p) > save {
				save = uf.Size(p)
				cand = x
			}
		}
	}

	return cand
}

type UF struct {
	set  []int
	size []int
}

func NewUFSet(n int) *UF {
	set := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
		size[i] = 1
	}
	return &UF{set, size}
}

func (this *UF) Find(x int) int {
	set := this.set
	var loop func(x int) int
	loop = func(x int) int {
		if x != set[x] {
			set[x] = loop(set[x])
		}
		return set[x]
	}
	return loop(x)
}

func (this *UF) Union(a, b int) {
	pa := this.Find(a)
	pb := this.Find(b)
	if pa == pb {
		return
	}
	size := this.size
	if size[pa] >= size[pb] {
		this.set[pb] = pa
		size[pa] += size[pb]
	} else {
		this.set[pa] = pb
		size[pb] += size[pa]
	}
}

func (this UF) Size(x int) int {
	return this.size[x]
}
