func numRabbits(answers []int) int {
    dic := make(map[int]int)
    for _, a := range answers {
        dic[a]++
    }
    var res int
    for k, v := range dic {
        var current = v / (k+1)
        if v % (k+1) != 0 {current += 1}
        res += current * (k+1)
    }
    return res
}
