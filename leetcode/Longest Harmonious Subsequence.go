func findLHS(nums []int) int {
    m := make(map[int]int)
    for _, x := range(nums) {
        if _, ok := m[x]; ok {
            m[x]++
        } else {
            m[x] = 1
        }
    }
    var res int
    for k := range m {
        if val, ok := m[k-1]; ok {
            option := m[k] + val
            if option > res {
                res = option
            }
        }
    }
    return res
}
