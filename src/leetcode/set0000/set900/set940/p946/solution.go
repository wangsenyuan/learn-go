package p946

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	if n == 0 {
		return true
	}
	stack := make([]int, n)
	var p int
	var j int
	for i := 0; i < n; i++ {
		stack[p] = pushed[i]
		p++
		for p > 0 && stack[p-1] == popped[j] {
			j++
			p--
		}
	}
	return j == n
}

func validateStackSequences1(pushed []int, popped []int) bool {
	n := len(pushed)
	if n == 0 {
		return true
	}

	stack := make([]int, n)

	var process func(p, i, j int) bool

	process = func(p, i, j int) bool {
		if j == n {
			return i == n
		}

		if p > 0 && p <= n && stack[p-1] == popped[j] {
			// pop
			return process(p-1, i, j+1)
		}
		if p == n || i == n {
			return false
		}

		stack[p] = pushed[i]
		return process(p+1, i+1, j)
	}

	return process(0, 0, 0)
}
