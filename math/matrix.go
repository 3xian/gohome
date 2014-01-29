package math

func Mul(a [][]int64, b [][]int64) [][]int64 {
	var (
		n = len(a)
		m = len(b[0])
		t = len(b) // len(a[0]) == len(b) 才有意义
		c = make([][]int64, n)
	)
	for i := 0; i < n; i++ {
		c[i] = make([]int64, m)
		for j := 0; j < m; j++ {
			for k := 0; k < t; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func Copy(a [][]int64) [][]int64 {
	b := make([][]int64, len(a))
	for i := range b {
		b[i] = make([]int64, len(a[i]))
		copy(b[i], a[i])
	}
	return b
}

func Pow(a [][]int64, k int64) [][]int64 {
	if k <= 0 {
		panic("k <= 0")
	}
	var (
		pd = Copy(a)
		pw = Copy(a)
	)
	for k--; k > 0; k /= 2 {
		if k&1 == 1 {
			pd = Mul(pd, pw)
		}
		if k > 1 {
			pw = Mul(pw, pw)
		}
	}
	return pd
}
