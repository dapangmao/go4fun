func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    dis := make([]int, n)
    for i:=0; i<n; i++ {dis[i] = 2 << 61}
    dis[src] = 0
    prev := make([]int, n)
    for i:=0; i<n; i++ {prev[i] = 2 << 61}
    prev[src] = 0
    for i:=0; i<=K; i++ {
        for _, flight := range flights {
            current := prev[flight[0]] + flight[2]
            if current < dis[flight[1]] {
                dis[flight[1]] = current
            }
        }
        prev = dis
    }
    if dis[dst] < 2 << 61 {return dis[dst]}
    return -1
}
