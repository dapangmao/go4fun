func missingNumber(nums []int) int {
    var sum, n = 0, len(nums)
    for _, x := range nums {sum += x}
    return (n+1) *n/2 - sum
}
