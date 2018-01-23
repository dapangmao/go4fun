func simplifyPath(path string) string {
    var stack = strings.Split(path, "/")
    var res []string
    var current int
    for i:=len(stack)-1; i>=0; i-- {
        if stack[i] == "." || len(strings.TrimSpace(stack[i])) == 0 {continue}
        if stack[i] == ".." {
            current++
        } else if current > 0 {
            current--
        } else {
            res = append([]string{stack[i]}, res...)
        }
    }
    return "/" + strings.Join(res, "/")
}
