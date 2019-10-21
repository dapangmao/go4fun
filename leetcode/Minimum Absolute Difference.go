func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
    diffs := make([]int, len(arr)-1)
    var current int
    minAbs := 3132131
    for i:=0; i<len(arr)-1; i++ {
        current = arr[i+1] - arr[i]
        diffs[i] = current
        if current < minAbs {minAbs = current}
    }
    var res [][]int
    for i, x := range diffs {
        if x != minAbs {continue}
        res = append(res, []int{arr[i], arr[i+1]})
    }
    return res
}
