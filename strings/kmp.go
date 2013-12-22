package strings

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func ExKMP(s string) (lcp []int) {
	n := len(s)
	lcp = make([]int, n)
	if n > 0 {
		lcp[0] = n
	}
	if n > 1 {
		for i := 1; i < n; i++ {
			if s[i-1] == s[i] {
				lcp[1]++
			} else {
				break
			}
		}
	}
	k := 1
	for i := 2; i < n; i++ {
		len1, len2 := k+lcp[k], lcp[i-k]
		if len2 < len1-i {
			lcp[i] = len2
		} else {
			j := max(len1-i, 0)
			for i+j < n && s[j] == s[i+j] {
				j++
			}
			lcp[i] = j
			k = i
		}
	}
	return
}
