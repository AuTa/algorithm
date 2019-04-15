package leetcode0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charCount := [26]int{}
	for _, r := range s {
		charCount[r-'a']++
	}
	for _, r := range t {
		charCount[r-'a']--
	}
	for _, c := range charCount {
		if c != 0 {
			return false
		}
	}
	return true
}
