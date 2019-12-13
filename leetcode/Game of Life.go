var options = [][]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}

func gameOfLife(board [][]int)  {
	n, m := len(board), len(board[0])
	clone := make(map[int][]int, n)
	for i, arr := range board {
		clone[i] = make([]int, m)
		copy(clone[i], arr)
	}

	count := func(i, j int) int {
		res := 0
		for _, op := range options {
			x, y := i + op[0], j + op[1]
			if x < 0 || x >= n || y <0 || y >= m {continue}
			res += clone[x][y]
		}
		return res
	}
    
    
    process := func(i, j, count int) {
        if clone[i][j] == 0 && count == 3 {
				board[i][j] = 1
        }
        if count == 2 || count == 3 {return}
        if clone[i][j] == 1 {
				board[i][j] = 0
        } 
    }

	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			current := count(i, j)
            process(i, j, current)
		}
	}
}
