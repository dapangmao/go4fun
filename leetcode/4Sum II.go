func fourSumCount(A []int, B []int, C []int, D []int) int {
    var res int
    var dic = make(map[int]int)
    for _, a := range A {
        for _, b := range B {
            dic[-a-b]++
        }
    }
    for _, c := range C {
        for _, d := range D {
            if val, ok := dic[c+d]; ok {res += val}
        }
    }
    return res
}
