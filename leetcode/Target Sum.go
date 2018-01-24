func findTargetSumWays(nums []int, S int) int {
    var dic = make(map[int]int)
    if len(nums) == 0 {return 0}
    dic[nums[0]]++
    dic[-nums[0]]++
    for i:=1; i<len(nums); i++ {
        var current = make(map[int]int)
        for k, v := range dic {
            current[k+nums[i]] += v
            current[k-nums[i]] += v
        }
        dic = current
    }
    if val, ok := dic[S]; ok {return val}
    return 0
}
