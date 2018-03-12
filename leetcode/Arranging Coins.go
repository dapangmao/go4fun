func arrangeCoins(n int) int {
    for i:=1; ;i++ {
        if n < i*(i+1)/2 {return i-1}
    }
    return 0
}
