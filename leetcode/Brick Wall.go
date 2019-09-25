func leastBricks(wall [][]int) int {
    var n = len(wall)
    var max int
    var dic = prep(wall)
    for _, val := range dic {
        if val > max {max = val}
    }
    return n - max
}

func prep(wall [][]int) map[int]int {
    var res = make(map[int]int)
    for _, w := range wall {
        var current int
        for i:=0; i<len(w)-1; i++ {
            current += w[i]
            res[current]++
        }
    }
    return res
}
