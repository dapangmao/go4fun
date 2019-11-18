func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
    res := [][]int{{r0, c0}}
    total := R*C
    for k:=1; k<2*(R+C); k=k+2 {
        for _, option := range [][]int{{0, 1, k}, {1, 0, k}, {0, -1, k+1}, {-1, 0, k+1}} {
            dr, dc, dk := option[0], option[1], option[2]
            for i:=0; i<dk; i++ {
                r0, c0 = r0+dr, c0+dc
                if 0 <= r0 && r0 < R && 0 <= c0 && c0 < C {
                    res = append(res, []int{r0, c0})
                    if len(res) == total {return res}
                }
            }
        }
    }
    return res
}
