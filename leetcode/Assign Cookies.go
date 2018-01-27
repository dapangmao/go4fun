func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    var j, res int
    for _, x := range g {
        for j < len(s) && s[j] < x {
            j++
        } 
        if j == len(s) {break}
        if s[j] >= x {
            j++
            res++
        }
    }
    return res
}
