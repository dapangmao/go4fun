func generateParenthesis(n int) []string {
    var res []string
    var dfs func(left, right int, path string)
    dfs = func(left, right int, path string) {
        if left == 0 && right == 0 {
            res = append(res, path)
            return
        }
        if left > 0 {
            dfs(left-1, right, path + "(")
        }    
        if right > left {
            dfs(left, right-1, path + ")")
        }
    }
    dfs(n, n, "")
    return res
}
