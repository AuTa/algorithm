package leetcode1009

func bitwiseComplement(N int) int {
	var i = 1
	for i < N {
		i = i<<1 + 1
	}
	return N ^ i
}
