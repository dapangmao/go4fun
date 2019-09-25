func rob(root *TreeNode) int {
    var dfs func(root *TreeNode) (int, int) 
    dfs = func(root *TreeNode) (int, int) {
        if root == nil {
            return 0, 0
        }
        rob_L, no_rob_L := dfs(root.Left)
        rob_R, no_rob_R := dfs(root.Right)
        var prev = rob_L + rob_R
        var res = prev
        var no_rob = no_rob_L + no_rob_R + root.Val
        if no_rob > res {res = no_rob}
        return res, prev 
    }
    res, _ := dfs(root)
    return res
}
