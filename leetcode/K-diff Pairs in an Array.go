func findPairs(nums []int, k int) int {
    if k < 0 {return 0}
    var dic = make(map[int]bool)
    var res int
    for _, x := range nums {
        if k != 0 {
            if _, ok := dic[x]; ok {continue}
            if _, ok := dic[x+k]; ok {res++}
            if _, ok := dic[x-k]; ok {res++}
            dic[x] = false
        } else {
            if val, ok := dic[x]; ok {
                if !val {res++}
                dic[x] = true
            } else {
                dic[x] = false
            }
        }
    }
    return res
}
