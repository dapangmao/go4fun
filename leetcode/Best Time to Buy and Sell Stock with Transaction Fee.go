func maxProfit(prices []int, fee int) int {
    n := len(prices)
    cash, hold := 0, -prices[0]
    var currentCash, currentHold int
    for i:=1; i<n; i++ {
        currentCash = hold + prices[i] - fee
        if currentCash > cash {
            cash = currentCash
        }
        currentHold = cash - prices[i]
        if currentHold > hold {
            hold = currentHold
        }
    }
    return cash
}
