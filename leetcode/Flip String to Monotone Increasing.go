func minFlipsMonoIncr(S string) int {
    ba := []byte(S)
    l := len(ba)
    onesLeft, onesRight := make([]int, l), make([]int, l)
    var current int
    for i := range ba {
        onesLeft[i] = current
        if S[i] == '1' { current++}
    }
    current = 0
    for i:=l-1; i>=0; i-- {
        onesRight[i] = current
        if ba[i] == '0' { current++}
    }
    minNum := 342432423432
    for i := range ba {
        current = onesLeft[i] + onesRight[i]        
        if current < minNum {minNum = current}
    }
    return minNum
}
