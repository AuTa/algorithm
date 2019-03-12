package leetcode0060

import "strconv"

func getPermutation(n int, k int) string {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i + 1
	}
	return sliceK("", s, n, k-1)
}

func sliceK(ans string, s []int, n, k int) string {
	if n == 1 {
		return ans + strconv.Itoa(s[0])
	}
	preF := factorial(n - 1)
	index := (k) / preF
	nextK := (k) % preF
	sn := s[index]
	for i := index + 1; i < n; i++ {
		s[i-1] = s[i]
	}
	return sliceK(ans+strconv.Itoa(sn), s, n-1, nextK)
}

var f = make(map[int]int, 0)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	if v, ok := f[n]; ok {
		return v
	}
	f[n] = n * factorial(n-1)
	return f[n]
}
