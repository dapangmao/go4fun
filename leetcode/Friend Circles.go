func findCircleNum(M [][]int) int {
    n, m := len(M), len(M[0])
    var visited = make(map[int]bool)

    graph := make(map[int][]int)
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if j > i {continue}  // don't give up those orpahns 
            if M[i][j] == 1 {
                graph[i] = append(graph[i], j)  
                graph[j] = append(graph[j], i) 
            }
        }
    }
    
    var dfs func(i int, graph map[int][]int)
    dfs = func(i int, graph map[int][]int) {
        if _, ok := visited[i]; ok {return}
        visited[i] = true
        for _, j := range graph[i] {
            dfs(j, graph)
        }
    }
    
    var res int
    for k, _ := range graph {
        if _, ok := visited[k]; ok {continue}
        dfs(k, graph)
        res++
    }
    return res
}
