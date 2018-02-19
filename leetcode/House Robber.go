func rob(nums []int) int {
    n := len(nums)
    if n == 0 {return 0}
    if n == 1 {return nums[0]}
    if nums[1] < nums[0] {nums[1] = nums[0]} 
    for i:=2; i<n; i++ {
        if nums[i-1] > nums[i-2] + nums[i] {
            nums[i] = nums[i-1]
        } else {
            nums[i] += nums[i-2]
        }
    }
    return nums[n-1]
}
