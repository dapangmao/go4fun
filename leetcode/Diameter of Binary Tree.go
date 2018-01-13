var res int

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {return res}
    dfs(root)
    return res
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l, r := dfs(root.Left), dfs(root.Right)
    if l + r > res {res = l + r}
    if l > r {return l + 1}
    return r + 1
}
