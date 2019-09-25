func maximumProduct(nums []int) int {
    var n = len(nums)
    sort.Ints(nums)
    var cand1 = nums[n-1] * nums[n-2] * nums[n-3] 
    var cand2 = nums[n-1] * nums[0] * nums[1]
    if cand1 > cand2 {return cand1}
    return cand2
}
