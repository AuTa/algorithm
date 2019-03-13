package leetcode0720

import (
	"math/rand"
	"sort"
)

type Solution struct {
	bMap map[int]int
	size int
}

func Constructor(N int, blacklist []int) Solution {
	bSize := len(blacklist)
	sort.Ints(blacklist)
	bMap := make(map[int]int, bSize)
	used := make(map[int]struct{}, 0)
	offset := 0
	for i := bSize - 1; i >= 0; i-- {
		n := blacklist[i]
		if n < (N - bSize) {
			for {
				o := N - bSize + offset
				if _, ok := used[o]; ok {
					offset++
				} else {
					bMap[n] = o
					used[o] = struct{}{}
					break
				}
			}
		} else {
			bMap[n] = n
			used[n] = struct{}{}
		}
	}
	return Solution{bMap, N - bSize}
}

func (this *Solution) Pick() int {
	n := rand.Intn(this.size)
	if v, ok := this.bMap[n]; ok {
		return v
	}
	return n
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
