func shortestToChar(S string, C byte) []int {
    n := len(S)
    res := make([]int, n)
    var current = 888888
    barr := []byte(S) 
    for i, x := range barr{
        if x == C {
            current = i
        } 
        if current == 888888 {
            res[i] = 888888
        } else {
            res[i] = i - current
        }
    }
    current = 888888
    for i:=n-1; i>=0; i-- {
        if barr[i] == C {
            current = i
        }
        if current == 888888 {continue}
        if current - i < res[i] {
            res[i] = current - i
        }
    }
    return res
}
