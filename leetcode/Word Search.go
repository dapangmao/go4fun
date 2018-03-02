func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	var dfs func(y, x, i int) bool
	dfs = func(y, x, i int) bool {
		if i == len(word) { return true}
		if y<0 || x<0 || x == m || y == n || board[y][x] != word[i] {return false}
		board[y][x] ^= 25
		i++
		if dfs(y, x+1, i) || dfs(y, x-1, i) || dfs(y+1, x, i) || dfs(y-1, x, i) {
			return true
		}
		board[y][x] ^= 25
		return false
	}
	
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if dfs(i, j, 0) {return true}
		}
	}
	return false
}
