func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
    g := func(a, b []int) int {
        x := a[0] - b[0]
        y := a[1] - b[1]
        return x*x + y*y
    } 
    ds := []int{g(p1, p3), g(p1, p4), g(p2, p3), g(p2, p4), g(p3, p4), g(p1, p2)}
    sort.Ints(ds)
    if ds[0] == 0 {return false}
    return ds[0] == ds[1] && ds[1] == ds[2] && ds[2] == ds[3] 
}
