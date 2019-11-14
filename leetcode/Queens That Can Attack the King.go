var dirs [][]int


func init() {
    for i:=-1; i<2; i++ {
        for j:=-1; j<2; j++ {
            if i == 0 && j == 0 {continue}
            dirs = append(dirs, []int{i, j})
        }
    }
}



func queensAttacktheKing(queens [][]int, king []int) [][]int {
    grid := make([][]int, 8)
    for i:=0; i<8; i++ {grid[i] = make([]int, 8)}
    for _, q := range queens {
        x, y := q[0], q[1]
        grid[x][y] = 1
    }
    x, y := king[0], king[1]
    var res [][]int
    for i:=1; i<8; i++ {
        var levels [][]int
        for _, d := range dirs {
            a := x + d[0] * i
            b := y + d[1] * i
            if 0 <= a && a< 8 && 0 <= b && b< 8 && grid[a][b] == 1 {
                res = append(res, []int{a, b})
            } else {
                levels = append(levels, d)
            }
        }
        dirs = levels
    }
    return res
}
