package leetcode0848

func shiftingLetters(S string, shifts []int) string {
	s := []byte(S)
	var sum int
	for i := len(shifts) - 1; i >= 0; i-- {
		sum = (sum + shifts[i]) % 26
		s[i] = (s[i]+byte(sum)-'a')%26 + 'a'
	}
	return string(s)
}
