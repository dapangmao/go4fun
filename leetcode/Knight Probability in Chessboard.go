func makeMatrix(N int) [][]float64 {
	res := make([][]float64, N)
	for i := range res {
		res[i] = make([]float64, N)
	}
	return res
}

func knightProbability(N int, K int, r int, c int) float64 {
	opts := [][]int{{2,1},{2,-1},{-2,1},{-2,-1},{1,2},{1,-2},{-1,2},{-1,-2}}
	dp := makeMatrix(N)
    dp[r][c] = 1
	for i:=0; i<K; i++ {
		current := makeMatrix(N)
		for j, x := range dp {
			for k, y := range x {
				for _, arr := range opts {
					row, col := j + arr[0], k + arr[1]
					if row >= 0 && row < N && col >=0 && col < N {
						current[row][col] += y / 8.0
					}
				}
			}
		}
        dp = current
	}
	var sum float64 
	for _, row := range dp {
		for _, val := range row {
			sum += val
		}
	}
	return sum
}
