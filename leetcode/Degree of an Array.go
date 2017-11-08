func findShortestSubArray(nums []int) int {
    left, right, count := make(map[int]int), make(map[int]int), make(map[int]int)
    for i, x := range nums {
        if _, ok := left[x]; !ok {
            left[x] = i
        }
        right[x] = i
        l, ok := count[x] 
        if ok {
            count[x] = l + 1
        } else {
            count[x] = 1
        }
    }     
    ans := len(nums)
    degree := -1
    for _, v := range count {
        if v > degree {
            degree = v
        }
    }
    for x := range count {
        current := right[x] - left[x] + 1
        if count[x] == degree && current < ans {
            ans = current
        }
    }
    return ans
}
