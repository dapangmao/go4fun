func minSubArrayLen(s int, nums []int) int {
    dic := make(map[int]int)
    dic[0] = -1
    var current int
    max := math.MaxInt32
    for i, num := range nums {
        current += num
        if j, ok := dic[current-s]; ok {
            if i-j < max {max = i-j}
        } 
        dic[current] = i
    }
    if max == math.MaxInt32 {return 0}
    return max
}
