func climbStairs(n int) int {
    if n == 1 {return 1}
    prevprev, prev, current := 1, 2, 2
    for i:=2; i<n; i++ {
        current = prev + prevprev
        prevprev = prev
        prev = current
    }
    return current
}
