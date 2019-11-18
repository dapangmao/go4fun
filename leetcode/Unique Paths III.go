func uniquePathsIII(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	var todo, sr, sc, tr, tc, ans int
	dr, dc := []int{0, -1, 0, 1}, []int{1, 0, -1, 0}
	for  r:= 0; r < R; r++ {
		for c := 0; c < C; c++ {
			switch grid[r][c] {
            case -1:
                goto next
			case 1:
				sr, sc = r, c
			case 2:
				tr, tc = r, c
			}
            todo++
            next:
		}
	}
	var dfs func(r, c, todo int)
	dfs = func(r, c, todo int) {
		todo--
		if todo < 0 {return}
		if r == tr && c == tc {
			if todo == 0 {ans++}
			return
		}
		grid[r][c] = 3
		for  k := 0; k < 4; k++ {
			var nr = r + dr[k]
			var nc = c + dc[k]
			if 0 <= nr && nr < R && 0 <= nc && nc < C {
				if grid[nr][nc] % 2 == 0 {dfs(nr, nc, todo)}
			}
		}
		grid[r][c] = 0
	}

	dfs(sr, sc, todo)
	return ans
}
