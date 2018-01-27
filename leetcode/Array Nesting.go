func arrayNesting(nums []int) int {
	var res int
	for i:=0; i<len(nums); i++ {
		var current = dfs(i, nums)
		if current > res {res = current}
	}
	return res
}

func dfs(i int, path []int) int {
	if path[i] < 0 {return 0}
	var current = path[i]
	path[i] = -1
	return 1 + dfs(current, path)
}
