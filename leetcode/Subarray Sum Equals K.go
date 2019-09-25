func subarraySum(nums []int, k int) int {
    dic := make(map[int]int)
    dic[0] = 1
    var current, res int
    for _, x := range nums {
        current += x
        if val, ok := dic[current-k]; ok {res += val}
        dic[current]++

    }
    return res
}
