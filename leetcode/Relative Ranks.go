func findRelativeRanks(nums []int) []string {
	max := -1
	for _, num := range nums {
		if num > max {max = num}
	}
	var dp = make([]int, max+1)
	for i, num := range nums {
		dp[num] = i+1
	}
	var res = make([]string, len(nums))
	var top = []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	var j = 4
	for i:=max; i>=0; i-- {
		if dp[i] == 0 {continue}
		if len(top) > 0 {
			res[dp[i]-1] = top[0]
			top = top[1:]
		} else {
            res[dp[i]-1] = fmt.Sprint(j)
			j++
		}
	}
	return res
}
