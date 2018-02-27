func customSortString(S string, T string) string {
    lookup := make(map[byte]int)
    for i, x := range []byte(S) {
        lookup[x] = i
    }
    var bs = []byte(T)
    sort.Slice(bs, func(i, j int) bool {
        ival, iok := lookup[bs[i]]
        jval, jok := lookup[bs[j]]
        if iok && jok {return ival < jval}
        if iok && !jok {return true}
        if !iok && jok {return false}
        return bs[i] < bs[j]
    })
    return string(bs)
}
