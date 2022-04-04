package myalgo

func IsPrime(n int) bool {
	if n <= 0 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func DivisorEnumeration(n int) []int {
	if n <= 0 {
		return []int{}
	}
	var ret []int
	for i := 1; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		ret = append(ret, i)
		if i != n/i {
			ret = append(ret, n/i)
		}
	}
	return ret
}

func PrimeFactrization(n int) []int {
	if n <= 0 {
		return []int{}
	}
	var ret []int
	for i := 2; i*i <= n; i++ {
		// 素数でなければskip
		for n%i == 0 {
			n = n / i
			ret = append(ret, i)
		}
	}
	if n >= 2 {
		ret = append(ret, n)
	}
	return ret
}
