package leetcode0239

func maxSlidingWindow(nums []int, k int) []int {
	var q []int
	var output []int
	for i := 0; i < len(nums); i++ {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i-k == q[0] {
			q = q[1:]
		}
		if i+1 >= k {
			output = append(output, nums[q[0]])
		}
	}
	return output
}
