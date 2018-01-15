func containsNearbyDuplicate(nums []int, k int) bool {
    var m = make(map[int]int)
    for i, x := range nums {
        if val, ok := m[x]; ok {
            if i-val <= k {return true}
        } 
        m[x] = i
    }
    return false
}
