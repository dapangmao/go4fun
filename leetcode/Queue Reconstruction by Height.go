func reconstructQueue(people [][]int) [][]int {
    sort.Slice(people, func(i, j int) bool {
        if people[i][0] != people[j][0] {
            return people[i][0] > people[j][0]
        }
        return people[i][1] < people[j][1]
    }) 
    var res [][]int
    for _, p := range people {
        var i = p[1]
        // res = append(res[:i], append([][]int{p}, res[i:]...)...)
        res = append(res, []int{0, 0})
        copy(res[i+1:], res[i:])
        res[i] = p
    }
    return res
}

