func partitionLabels(S string) []int {
    dic := make(map[rune]int)
    for i, x := range S {
        dic[x] = i
    }
    var arrs [][]int
    for i, x := range S {
        if val, ok := dic[x]; ok {
            if len(arrs) == 0 || i > arrs[len(arrs)-1][1] {
                arrs = append(arrs, []int{i, val})
            } else if val > arrs[len(arrs)-1][1] {
                arrs[len(arrs)-1][1] = val
            }
            delete(dic, x)
        }
    }
    var res []int
    for _, arr := range arrs {
        res = append(res, arr[1]-arr[0]+1)
    }
    return res
}
