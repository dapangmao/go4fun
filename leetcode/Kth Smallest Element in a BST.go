func kthSmallest(root *TreeNode, k int) int {
    var dfs func(root *TreeNode) []int
    dfs = func(root *TreeNode) []int {
        if root == nil {return []int{}}
        var current = append(append(dfs(root.Left), root.Val), dfs(root.Right)...)
        return current
    }
    res := dfs(root)
    return res[k-1]
}
