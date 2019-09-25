func canReorderDoubled(A []int) bool {
    dict := make(map[int]int)
    var ints []int
    for _, x := range A {
        if _, ok := dict[x]; !ok {ints = append(ints, x)}
        dict[x]++
    }
    sort.Slice(ints, func(i, j int) bool {
        a, b := ints[i], ints[j]
        if a < 0 {a = -a}
        if b < 0 {b = -b}
        return a < b
    })
    for _, x := range ints {
        if dict[x] == 0 {continue}
        if _, ok := dict[2*x]; !ok || dict[2*x] - dict[x] < 0 {
            return false
        }
        dict[2*x] -= dict[x]
        dict[x] = 0

    }
    for _, v := range dict {
        if v != 0 {return false}
    }
    return true
}
