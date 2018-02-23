func containsDuplicate(nums []int) bool {
    dic := make(map[int]int)
    for _, num := range nums {
        dic[num]++
        if dic[num] > 1 {return true}
    }
    return false
}
