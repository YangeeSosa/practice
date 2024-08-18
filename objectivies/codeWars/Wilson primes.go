package kata

func factorial(n int, mod int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = (result * i) % mod
	}
	return result
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func AmIWilson(n int) bool {
	if !isPrime(n) {
		return false
	}

	mod := n * n
	nFactorial := factorial(n-1, mod)

	return (nFactorial+1)%mod == 0
}
