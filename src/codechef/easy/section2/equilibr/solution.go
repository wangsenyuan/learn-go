package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)
	readNum(scanner)
	fmt.Println(solve(n))
}

const MOD = 1000000007

func solve(n int) int {
	x := pow(2, n-1)
	y := x - n
	if y < 0 {
		y += MOD
	}

	xi := inverse(x)

	return int((int64(y) * int64(xi)) % MOD)
}

func inverse(n int) int {
	return pow(n, MOD-2)
}

func pow(a, b int) int {
	x := int64(1)
	y := int64(a)

	for b > 0 {
		if b&1 == 1 {
			x = (x * y) % MOD
		}
		y = (y * y) % MOD
		b >>= 1
	}
	return int(x)
}
