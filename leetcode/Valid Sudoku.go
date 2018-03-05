func isValidSudoku(board [][]byte) bool {
    
    check9 := func(i, j int) bool {
        set := make(map[byte]bool)
        for l:=-1; l<2; l++ {
            for k:=-1; k<2; k++ {
                if board[i+l][j+k] == '.' {continue}
                if set[board[i+l][j+k]] {return false}
                set[board[i+l][j+k]] = true
            }
        }
        return true
    }
    
    checkRow := func(i int) bool {
        set := make(map[byte]bool)
        for _, x := range board[i] {
            if x == '.' {continue}
            if set[x] {return false}
            set[x] = true
        }
        return true
    }
    
    checkColumn := func(j int) bool {
        set := make(map[byte]bool)
        for i:=0; i<9; i++ {
            if board[i][j] == '.' {continue}
            if set[board[i][j]] {return false}
            set[board[i][j]] = true
        }
        return true
    }
    
    v, h := 4, 4
    for _, i := range []int{-3, 0, 3} {
        for _, j := range []int{-3, 0, 3} {
            l, k := v + i, h +j 
            if !check9(l, k) {return false} 
        }
    }
    
    for i:=0; i<9; i++ {
        if !checkRow(i) {return false}
        if !checkColumn(i) {return false}
    }
    return true
}
