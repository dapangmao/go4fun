func powerfulIntegers(x int, y int, bound int) []int {
    var res []int
    var set = make(map[int]bool)
    for i:=0; i<17; i++ {
        for j:=0; j<17; j++ {
            current := pow(x, i) + pow(y, j)
            if !set[current]  && current <= bound {
                res = append(res, current)
                set[current] = true
            }
        }
    }
    return res
}

func pow(i, j int) int {
    var res = 1
    for ; j > 0; j-- {
        res = i*res
        if res >= math.MaxInt32 || res <= math.MinInt32 {return 1}
    }
    return res
}
