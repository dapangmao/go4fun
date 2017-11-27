func pivotIndex(nums []int) int {
    n := len(nums)
    if n == 0 {return -1}
    left, right := make([]int, n), make([]int, n)
    left[0], right[n-1] = nums[0], nums[n-1]
    for i:=1; i<n-1; i++ {
        left[i] += left[i-1] + nums[i]
    }
    for i:=n-2; i>=0; i-- {
        right[i] += right[i+1] + nums[i]
    }
    for i:=0; i<n; i++ {
        l := 0
        if i > 0 {
            l = left[i-1]
        }
        r := 0
        if i < n-1 {
            r = right[i+1]
        }
        if l == r {
            return i 
        }
    }
    return -1
}
