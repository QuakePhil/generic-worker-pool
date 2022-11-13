package primes

// https://en.wikipedia.org/wiki/Primality_test#Simple_methods
func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n <= 1 || n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
