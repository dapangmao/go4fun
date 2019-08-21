func findJudge(N int, trust [][]int) int {
    if N == 1 && len(trust) == 0 {return 1}
    pool := make([]int, N+1)
    for _, t := range trust {
        x, y := t[0], t[1]
        pool[x] = -1
        pool[y] != -1 {
            pool[y]++
        }
    }
    for p, count := range pool {
        if count == N-1 {
            return p
        }
    }
    return -1
}
