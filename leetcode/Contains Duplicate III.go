func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {return false}
	var dic = make(map[int]int)
	var w = t + 1

	subDuplicate := func(x, j, w int) bool {
		if val, ok := dic[j]; ok {
			current := x - val
			if current < 0 {
				current = -current
			}
			if current < w {return true}
		}
		return false
	}

	for i, x := range nums {
		var m = x / w
		if _, ok := dic[m]; ok {
			return true
		}
		if subDuplicate(x, m-1, w) {return true}
		if subDuplicate(x, m+1, w) {return true}
		dic[m] = x
		if i >= k {delete(dic, nums[i-k]/w)}
	}
	return false
}
