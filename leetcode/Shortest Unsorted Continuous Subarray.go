func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    var should = make([]int, n)
    copy(should, nums)
    sort.Ints(nums)
    var i, j = 0, n-1
    for ; i<n; i++ {
        if nums[i] != should[i] {break}
    }
    for ; j>=0; j-- {
        if nums[j] != should[j] {break}
    }
    var res = j - i + 1
    if res > 0 {return res}
    return 0
}
