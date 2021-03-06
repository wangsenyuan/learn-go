package p797

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)

	var dfs func(v int, m int, path []int)
	res := make([][]int, 0, n)

	dfs = func(v int, m int, path []int) {
		if v == n-1 {
			res = append(res, copySlice(path[:m]))
			return
		}
		for _, w := range graph[v] {
			path[m] = w
			dfs(w, m+1, path)
		}
	}

	path := make([]int, n)
	path[0] = 0
	dfs(0, 1, path)

	return res
}

func copySlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}
