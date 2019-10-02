func permute(nums []int) [][]int {
	if len(nums) == 0 {return [][]int{{}}}
	var res [][]int
	n := len(nums)
	last := nums[n-1]
	for _, x := range permute(nums[:n-1]) {
		for i:=0; i<=len(x); i++ {
			current := make([]int, len(x)+1)
			copy(current, x[:i])
			current[i] = last
			copy(current[i+1:], x[i:])
			res = append(res, current)
		}
	}
	return res
}

