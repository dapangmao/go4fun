func findMaxLength(nums []int) int {
    var n = len(nums)
    if n < 0 {return 0}
    var dic = make(map[int]int)
    dic[0] = -1
    res, count := 0, 0
    for i:=0; i<n; i++ {
        if nums[i] == 0 {
            count--
        } else {
            count++
        }
        if val, ok := dic[count]; ok {
            var current = i - val
            if current > res {res = current}
        } else {
            dic[count] = i
        }
    }
    return res
}
