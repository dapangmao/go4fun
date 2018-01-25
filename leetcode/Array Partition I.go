func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    var res int
    for i, x := range nums {
        if i%2==0 {res += x}
    }
    return res
}
