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

func readOneNum(scanner *bufio.Scanner) (a int) {
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

	t := readOneNum(scanner)

	for t > 0 {
		n := readOneNum(scanner)
		res := solve(n)

		fmt.Printf("%d", res[0])

		for i := 1; i < n; i++ {
			fmt.Printf(" %d", res[i])
		}
		fmt.Println()
		t--
	}
}

func solve(n int) []int {
	if n < 2 {
		panic("require n at least 2")
	}
	res := make([]int, n)
	var i int
	for i+1 < n {
		res[i] = i + 2
		res[i+1] = i + 1
		i += 2
	}

	if i+1 == n {
		res[i] = i - 1
		res[i-1] = i + 1
	}

	return res
}
