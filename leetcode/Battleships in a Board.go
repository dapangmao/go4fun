func countBattleships(board [][]byte) int {
    var res int
    n, m := len(board), len(board[0])
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if board[i][j] != 'X' {continue}
            var queue = [][]int{[]int{i, j}}
            var current []int
            for len(queue) > 0 {
                current, queue = queue[0], queue[1:]
                l, k := current[0], current[1]
                if l > 0 && board[l-1][k] == 'X' {queue = append(queue, []int{l-1, k})}
                if l < n-1 && board[l+1][k] == 'X' {queue = append(queue, []int{l+1, k})}
                if k > 0 && board[l][k-1] == 'X' {queue = append(queue, []int{l, k-1})}
                if k < m-1 && board[l][k+1] == 'X' {queue = append(queue, []int{l, k+1})}
                board[l][k] = 'Z'
            }
            res++
        }
    }
    return res
}
