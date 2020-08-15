package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, q := readThreeNums(reader)
	E := make([][]int, m)

	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}
	fmt.Println(solve(n, m, q, E))
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const INF = 1 << 62
const MOD = 1000000007
const inv2 = (MOD + 1) / 2

func add(a int, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	if a < 0 {
		a += MOD
	}
	return a
}

func mul(a int, b int) int {
	A := int64(a)
	B := int64(b)

	return int((A * B) % MOD)
}

func calc(a1, d, n int) int {
	return mul(mul(n, inv2), add(mul(2, a1), mul(add(n, -1), d)))
}

type Frac struct {
	x, y int64
}

func NewFrac(x, y int64) Frac {
	if y < 0 {
		return Frac{-x, -y}
	}
	return Frac{x, y}
}

func (this Frac) Less(that Frac) bool {
	return this.x*that.y <= this.y*that.x
}

type Line struct {
	m, c int64
}

func (this Line) Intersect(that Line) Frac {
	return NewFrac(this.c-that.c, that.m-this.m)
}

type Lines []Line

func (lines Lines) Len() int {
	return len(lines)
}

func (lines Lines) Less(i, j int) bool {
	a, b := lines[i], lines[j]
	if a.m != b.m {
		return a.m < b.m
	}
	return a.c < b.c
}

func (lines Lines) Swap(i, j int) {
	lines[i], lines[j] = lines[j], lines[i]
}

func solve(n int, m int, q int, E [][]int) int {
	hv := make([]int, n)

	for _, e := range E {
		v, u, w := e[0], e[1], e[2]
		v--
		u--
		hv[v] = max(hv[v], w)
		hv[u] = max(hv[u], w)
	}

	var ans int

	d := make([]int64, n)
	for i := 0; i < n; i++ {
		d[i] = -INF
	}
	d[0] = 0
	nd := make([]int64, n)

	for val := 0; val < m; val++ {
		var mx int64

		for i := 0; i < n; i++ {
			mx = max2(mx, d[i])
		}
		if val > 0 {
			ans = add(ans, int(mx%MOD))
		}
		copy(nd, d)

		for _, e := range E {
			v, u, w := e[0], e[1], e[2]
			v--
			u--
			nd[v] = max2(nd[v], d[u]+int64(w))
			nd[u] = max2(nd[u], d[v]+int64(w))
		}
		copy(d, nd)
	}

	fin := make([]Line, n)

	for i := 0; i < n; i++ {
		fin[i] = Line{int64(hv[i]), d[i]}
	}

	sort.Sort(Lines(fin))

	var j int

	for i := 1; i <= n; i++ {
		if i == n || fin[i].m > fin[i-1].m {
			fin[j] = fin[i-1]
			j++
		}
	}

	fin = fin[:j]

	ch := make([]Line, j)
	var p int
	for _, cur := range fin {
		for p > 1 && cur.Intersect(ch[p-1]).Less(ch[p-1].Intersect(ch[p-2])) {
			p--
		}

		ch[p] = cur
		p++
	}

	var prv int64
	Q := int64(q - m)

	for i := 0; i < p-1; i++ {
		f := ch[i].Intersect(ch[i+1])

		if f.x < 0 {
			continue
		}

		lst := min2(Q, f.x/f.y)

		if lst < prv {
			continue
		}

		ans = add(ans, calc(int((ch[i].c+ch[i].m*prv)%MOD), int(ch[i].m%MOD), int(lst-prv+1)))

		prv = lst + 1
	}

	ans = add(ans, calc(int((ch[p-1].c+ch[p-1].m*prv)%MOD), int(ch[p-1].m%MOD), int(Q-prv+1)))

	return ans
}

func min2(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func max2(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
