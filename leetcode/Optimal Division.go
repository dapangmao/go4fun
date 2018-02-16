func optimalDivision(nums []int) string {
    n := len(nums)
    if n == 0 {return ""}
    var res = strconv.Itoa(nums[0])
    if n == 1 {return res}
    if n == 2 {return fmt.Sprintf("%v/%v", res, nums[1])}
    res += "/("
	for i:=1; i<n-1; i++ {
		res += strconv.Itoa(nums[i]) + "/"
	}
    return fmt.Sprintf("%v%v)", res, nums[n-1])  
}
