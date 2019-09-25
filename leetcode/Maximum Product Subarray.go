func maxProduct(nums []int) int {
    var n = len(nums)
    minDP, maxDP := make([]int, n), make([]int, n)
    copy(minDP, nums)
    copy(maxDP, nums)
    for i:=1; i<n; i++ {
        op1 := minDP[i-1] * minDP[i]
        op2 := maxDP[i-1] * maxDP[i]
        if op1 > maxDP[i] {maxDP[i] = op1}
        if op2 > maxDP[i] {maxDP[i] = op2}
        if op1 < minDP[i] {minDP[i] = op1}
        if op2 < minDP[i] {minDP[i] = op2}
    }
    var res = math.MinInt32
    for _, x := range maxDP {
        if x > res {res = x}
    }
    return res
}
