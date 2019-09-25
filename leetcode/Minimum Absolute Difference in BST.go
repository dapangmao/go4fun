var res = 13123213
var prev = new(TreeNode)


func getMinimumDifference(root *TreeNode) int {
    dfs(root)
    return res
}

func dfs(root *TreeNode) {
    if root == nil {return}
    dfs(root.Left)
    if prev != nil && res > root.Val - prev.Val {
        res = root.Val - prev.Val
    }
    prev = root
    dfs(root.Right)
}
