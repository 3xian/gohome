package math

func Primes(n int) (primes []int, fact []int) {
	fact = make([]int, n+1)
	for i := 2; i <= n; i++ {
		if fact[i] == 0 {
			for j := i; j <= n; j += i {
				fact[j] = i
			}
			primes = append(primes, i)
		}
	}
	return
}
