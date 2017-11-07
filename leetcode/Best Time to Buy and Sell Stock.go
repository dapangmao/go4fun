func maxProfit(prices []int) int {
    res := 0
    lowest := 324324
    for _, x := range prices {
        if x > lowest && x - lowest > res {
            res = x - lowest
        }
        if x < lowest {
            lowest = x
        }
    }
    return res
}
