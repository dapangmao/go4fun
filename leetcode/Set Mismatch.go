func findErrorNums(nums []int) []int {
    var res = []int{-1, -1}
    var dic = make(map[int]int)
    var sum, should = 0, 0
    for i, x := range nums {
        dic[x]++
        should += i+1
        sum += x
        if dic[x] > 1 {res[0] = x}
    }
    res[1] = res[0] + should - sum
    return res
}
