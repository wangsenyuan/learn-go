package p1400

func canConstruct(s string, k int) bool {
	n := len(s)
	if n < k {
		return false
	}
	if n == k {
		return true
	}

	cnt := make([]int, 26)

	for i := 0; i < n; i++ {
		cnt[int(s[i]-'a')]++
	}

	var odd int

	for i := 0; i < 26; i++ {
		odd += cnt[i] & 1
	}

	return odd <= k
}
