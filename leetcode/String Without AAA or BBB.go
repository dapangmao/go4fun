func strWithout3a3b(A int, B int) string {
    var res []string
    b, s := A, B
    bc, sc := "a", "b"
    if B > A {
        b, s = B, A
        bc, sc = "b", "a"
    }
    for i:=0; i<s; i++ {
        res = append(res, bc)
        res = append(res, sc)
        b--
    }
    if b > 0 {
        res = append(res, bc)
        b--
    }
    for i, x := range res {
        if b == 0 {break}
        if x == bc {
            res[i] = bc + bc
            b--
        }
    }
    return strings.Join(res, "")
}
