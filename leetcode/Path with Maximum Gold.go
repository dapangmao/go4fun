func max(x, y int) int {
    if y > x {return y}
    return x
}

func getMaximumGold(board [][]int) int {
    res := 0
    n, m := len(board), len(board[0])
	var dfs func(y, x, i int) 
	dfs = func(y, x, i int) {
		if y < 0 || x < 0 ||  x >= m || y == n || board[y][x] <= 0 {
            res = max(res, i)
            return
        }
        i += board[y][x]
        board[y][x] = -board[y][x]
		dfs(y, x+1, i)
        dfs(y, x-1, i)
        dfs(y+1, x, i)
        dfs(y-1, x, i) 
        board[y][x] = -board[y][x]
	}
    for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
            if board[i][j] == 0 {continue}
			dfs(i, j, 0) 
		}
	}
    return res
}
