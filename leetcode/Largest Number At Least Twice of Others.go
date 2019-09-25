func dominantIndex(nums []int) int {
	var max, max2 = -12423421, -12423421
	var i = -1
	for k, num := range nums {
		if num > max {
			max, max2 = num, max
			i = k
		} else if num > max2 {
			max2 = num
		}
	}
	if max >= max2*2 {
		return i
	}
	return -1
}
