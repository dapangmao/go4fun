func intervalIntersection(A [][]int, B [][]int) [][]int {
        i,j := 0, 0
        var res [][]int
        for i < len(A) && j < len(B) {
            if A[i][1] < B[j][0] {
                i++
            } else if A[i][1] < B[j][1] {
                res = append(res, []int{max(A[i][0],B[j][0]),A[i][1]})
                i++
            } else if A[i][0] <= B[j][1] {
                res = append(res, []int{max(A[i][0],B[j][0]),B[j][1]})
                j++
            } else if A[i][0] > B[j][1] {
               j++
            }
        }
        return res
}

func max(i, j int) int {
    if i > j {return i}
    return j
}
