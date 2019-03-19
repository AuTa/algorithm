package leetcode0273


import (
	"strings"
)

var EngMap = make(map[int]string, 27)

func initEngMap() {
	EngMap[1] = "One"
	EngMap[2] = "Two"
	EngMap[3] = "Three"
	EngMap[4] = "Four"
	EngMap[5] = "Five"
	EngMap[6] = "Six"
	EngMap[7] = "Seven"
	EngMap[8] = "Eight"
	EngMap[9] = "Nine"
	EngMap[10] = "Ten"
	EngMap[11] = "Eleven"
	EngMap[12] = "Twelve"
	EngMap[13] = "Thirteen"
	EngMap[14] = "Fourteen"
	EngMap[15] = "Fifteen"
	EngMap[16] = "Sixteen"
	EngMap[17] = "Seventeen"
	EngMap[18] = "Eighteen"
	EngMap[19] = "Nineteen"
	EngMap[20] = "Twenty"
	EngMap[30] = "Thirty"
	EngMap[40] = "Forty"
	EngMap[50] = "Fifty"
	EngMap[60] = "Sixty"
	EngMap[70] = "Seventy"
	EngMap[80] = "Eighty"
	EngMap[90] = "Ninety"
}

const (
	One      int = 1
	Thousand     = 1000 * One
	Million      = 1000 * Thousand
	Billion      = 1000 * Million
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	initEngMap()
	var res string

	for _, i := range []int{One, Thousand, Million, Billion} {
		if num < i {
			break
		}
		remainder := num / i % 1000
		rS := remainderToWords(remainder)
		var strs []string
		if rS != "" {
			switch i {
			case Thousand:
				strs = []string{rS, "Thousand"}
			case Million:
				strs = []string{rS, "Million"}
			case Billion:
				strs = []string{rS, "Billion"}
			default:
				strs = []string{rS}
			}
			if res != "" {
				strs = append(strs, res)
			}
			res = strings.Join(strs, " ")
		}
	}
	return res
}

func remainderToWords(r int) string {
	var res string
	if x := r % 100; x != 0 {
		if v, ok := EngMap[x]; ok {
			res = v
		} else {
			if x := r % 10; x != 0 {
				res += EngMap[x]
			}
			if x := (r - r%10) % 100; x != 0 {
				res = strings.Join([]string{EngMap[x], res}, " ")
			}
		}
	}

	if x := r / 100; x != 0 {
		var strs []string
		if res != "" {
			strs = []string{EngMap[x], "Hundred", res}
		} else {
			strs = []string{EngMap[x], "Hundred"}
		}
		res = strings.Join(strs, " ")
	}
	return res
}
