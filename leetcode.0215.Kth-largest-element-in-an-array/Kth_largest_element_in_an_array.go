package leetcode0215

import "container/heap"

type IntHeap []int

var _ heap.Interface = new(IntHeap)

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	// 最小堆
	// 只保留到第 k 位的数
	intHeap := (IntHeap)(make([]int, 0, k))
	h := &intHeap
	heap.Init(h)
	for _, item := range nums {
		if h.Len() < k {
			heap.Push(h, item)
		} else if item <= (*h)[0] {
			continue
		} else {
			heap.Pop(h)
			heap.Push(h, item)
		}
	}
	return (*h)[0]
}
