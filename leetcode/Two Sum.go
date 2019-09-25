func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, x := range nums {
        if v, ok := m[x]; ok {
            return []int{v, i}
        } else {
            m[target-x] = i
        }
    }
    return make([]int, 2)
}
