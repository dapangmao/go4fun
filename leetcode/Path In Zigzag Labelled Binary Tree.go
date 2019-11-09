func pathInZigZagTree(label int) []int {
    var res []int
    for ; label > 0; label=label/2 {
        res = append([]int{label}, res...)
    }
    for i:=len(res)-2; i>0; i-=2 {
        res[i] = pow(2, i) + pow(2, i+1) - 1 - res[i]
    }
    return res
}

func pow(n, i int) int {
    res := 1
    for ; i>0; i-- {res *= n}
    return res
}
