func majorityElement(nums []int) int {
    var res = nums[0]
    var count int
    for _, num := range nums {
        if num == res {
            count++
        } else {
            count--
            if count == 0 {
                res = num
                count = 1
            }
        }
    }
    return res
}
