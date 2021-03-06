package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, K := readTwoNums(scanner)
		rks := make([][]int, K)
		for i := 0; i < K; i++ {
			rks[i] = readNNums(scanner, 2)
		}
		res := solve(N, K, rks)
		fmt.Printf("%d", len(res))
		for i := 0; i < len(res); i++ {
			fmt.Printf(" %d %d", res[i][0], res[i][1])
		}
		fmt.Println()
	}
}

func solve(N int, K int, rks [][]int) [][]int {
	if K == N {
		return nil
	}
	row := make([]bool, N)
	col := make([]bool, N)

	for _, rk := range rks {
		x, y := rk[0], rk[1]
		row[x-1] = true
		col[y-1] = true
	}
	var u = 0
	res := make([][]int, N-K)
	for i, j := 0, 0; i < N; i++ {
		if !row[i] {
			//this row has no rook
			for col[j] {
				j++
			}
			res[u] = []int{i + 1, j + 1}
			u++
			row[i] = true
			col[j] = true
		}
	}

	return res
}
