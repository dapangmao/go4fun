func networkDelayTime(times [][]int, N int, K int) int {
    graph := make(map[int][][2]int)
    for _, time := range times {
        u, v, w: = time[0], time[1], time[2]
        graph[u] = append(graph[u], [2]int{v, w})
    }
    dist := make([]int, N+1)
    for i:=1; i<=N; i++ {
        dist[i] = 4444444
    }
    var dfs func(node, elapsed int)
    dfs = func(node, elapsed int) {
        if elapsed >= dist[node] {return}
        dist[node] = elapsed
        for _, pair := range graph[node] {
            dfs(pair[0], elapsed+pair[1])
        }
    }
    dfs(K, 0)
    ans := -1
    for _, v := range dist[1:]{
        if v == 4444444 {return -1}
        if v > ans {ans = v}
    }
    return ans
}
