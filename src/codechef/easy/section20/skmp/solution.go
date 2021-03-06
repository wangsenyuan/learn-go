package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		S, _ := reader.ReadString('\n')
		P, _ := reader.ReadString('\n')
		R := solve(S[:len(S)-1], P[:len(P)-1])
		fmt.Println(R)
	}
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

func solve(S string, P string) string {
	cnt := make([]int, 26)

	for i := 0; i < len(S); i++ {
		cnt[int(S[i]-'a')]++
	}

	for i := 0; i < len(P); i++ {
		cnt[int(P[i]-'a')]--
	}

	// all cnt should be >= 0
	res := make([]byte, len(S))
	var k int
	var x int

	for byte(x+'a') < P[0] {
		for cnt[x] > 0 {
			res[k] = byte(x + 'a')
			k++
			cnt[x]--
		}
		x++
	}

	if byte(x+'a') == P[0] {
		var j int
		for j < cnt[x] && j < len(P) && byte(x+'a') == P[j] {
			j++
		}
		if j == cnt[x] || j < len(P) && byte(x+'a') < P[j] {
			// safe to place x at begining
			for cnt[x] > 0 {
				res[k] = byte(x + 'a')
				k++
				cnt[x]--
			}
			x++
		}
		// then pat P
		for i := 0; i < len(P); i++ {
			res[k] = P[i]
			k++
		}
	}

	for x < 26 {
		for cnt[x] > 0 {
			res[k] = byte(x + 'a')
			k++
			cnt[x]--
		}
		x++
	}

	return string(res)
}
