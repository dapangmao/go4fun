var visited = map[int][]*TreeNode{1: []*TreeNode{&TreeNode{0, nil, nil}}}

func allPossibleFBT(N int) []*TreeNode {
    if _, ok := visited[N]; !ok {
        var res []*TreeNode
        for x:=0; x<N; x++ {
            y := N - 1 - x 
            for _, left := range allPossibleFBT(x) {
                for _, right := range allPossibleFBT(y) {
                    res = append(res, &TreeNode{0, left, right})
                }
            }
        }
        visited[N] = res  
    }
    return visited[N]  
}
