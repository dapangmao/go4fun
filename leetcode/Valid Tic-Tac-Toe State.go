func validTicTacToe(board []string) bool {
    nx, no := 0, 0
    winx, wino := 0, 0
    for _, row := range board {
        nx += strings.Count(row, "X")
        no += strings.Count(row, "O")
        if row == "XXX" {winx++}
        if row == "OOO" {wino++}
    }
    if no > nx {return false}
    if nx > no + 1 {return false}
    op4 := string([]byte{board[0][0], board[1][1], board[2][2]}) // it is easy to convert byte array to a string
    op5 := string([]byte{board[0][2], board[1][1], board[2][0]})
    op6 := string([]byte{board[0][0], board[1][0], board[2][0]})
    op7 := string([]byte{board[0][1], board[1][1], board[2][1]})
    op8 := string([]byte{board[0][2], board[1][2], board[2][2]}) 
    for _, op := range []string{op4, op5, op6, op7, op8} {
        if op == "XXX" {winx++}
        if op == "OOO" {wino++}
    }
    if winx > 0 && wino > 0 {return false}
    if winx > 0 && nx == no {return false}
    if wino > 0 && nx > no {return false}
    return true
}
