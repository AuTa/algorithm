package leetcode0334

func increasingTriplet(nums []int) bool {
	var min, max *int
	if len(nums) < 3 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if i == 0 {
			min = &n
		} else if i == 1 {
			if n > *min {
				max = &n
			} else {
				min = &n
			}
		} else {
			if max != nil && n > *max {
				return true
			} else if n < *min {
				min = &n
			} else if n > *min && (max == nil || (max != nil && n < *max)) {
				max = &n
			}
		}
	}
	return false
}
