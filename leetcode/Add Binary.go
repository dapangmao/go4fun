func addBinary(a string, b string) string {
    var prev int
    n, m := len(a), len(b)
    var res []byte
    for i:=0; i<n||i<m; i++ {
        var current = prev
        if i < n {
            if a[n-i-1] == '1' {current++}
        }
        if i < m {
            if b[m-i-1] == '1' {current++}
        }
        var c byte = '0'
        if current != 2 {
            c = '1'
        } 
        prev = 0
        if current > 1 {
            prev = 1
        } 
        res = append([]byte{c}, res...)
    }
    if prev == 1 {
        res = append([]byte{'1'}, res...)
    }
    return string(res)
}
