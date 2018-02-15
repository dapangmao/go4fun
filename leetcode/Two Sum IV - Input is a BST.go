var dic = make(map[int]bool)
var K int

func findTarget(root *TreeNode, k int) bool {
    K = k
    return dfs(root)
}

func dfs(root *TreeNode) bool {
    if root == nil {return false}
    if _, ok := dic[root.Val]; ok {return true}
    dic[K-root.Val] = true
    return dfs(root.Left) || dfs(root.Right)
}
