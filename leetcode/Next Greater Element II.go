func nextGreaterElements(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    for i:=0; i<n; i++ {res[i] = -1}
    var stack []int
    for i:=0; i<n*2; i++ {
        j := i % n
        for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[j] {
            pop, lastIdx := 0, len(stack)-1
            pop, stack = stack[lastIdx], stack[:lastIdx]
            res[pop] = nums[j]
        }
        stack = append(stack, j)
    }
    return res
}
