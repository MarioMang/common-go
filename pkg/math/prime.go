package math

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i < n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func IsNotPrimce(n int) bool {
	return !IsPrime(n)
}
