func groupThePeople(groupSizes []int) [][]int {
    var res [][]int
    freq := make(map[int][]int)
    for i, n := range groupSizes {
        freq[n] = append(freq[n], i)
    }
    for k, v := range freq {
        for i:=0; i<len(v)-k+1; i+=k {
            res = append(res, v[i:i+k])
        }
    }
    return res
}
