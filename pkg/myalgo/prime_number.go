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
