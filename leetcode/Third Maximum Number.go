func thirdMax(nums []int) int {
    const m = -9223372036854775808
    var max1, max2, max3 = m, m, m
    for _, num := range nums {
        if num == max1 || num == max2 || num == max3 {continue}
        if num > max1 {
           max3 = max2
           max2 = max1
           max1 = num
        } else if num > max2 {
           max3 = max2
           max2 = num
        } else if num > max3 {
           max3 = num
        }
    }
    if max3 == m {return max1}
    return max3
}
