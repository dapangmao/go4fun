func isValidBST(root *TreeNode) bool {
    var dfs func(root *TreeNode, min, max int) bool
    dfs = func(root *TreeNode, min, max int) bool {
        if root == nil {return true}
        if root.Val <= min || root.Val >= max {return false}
        return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)  
    }
    return dfs(root, math.MinInt64, math.MaxInt64)
}
