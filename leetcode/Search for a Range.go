func searchRange(nums []int, target int) []int {
    n := len(nums)
    var i = sort.SearchInts(nums, target)
    if i >= n || nums[i] != target {return []int{-1, -1}}
    j := i 
    for j < n && nums[j] == target {
        j++
    }
    return []int{i, j-1}
}
