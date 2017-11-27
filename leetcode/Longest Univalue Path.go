var res int

func longestUnivaluePath(root *TreeNode) int {
    dfs(root)
    return res
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l := dfs(root.Left)
    r := dfs(root.Right)
    left, right := 0, 0
    if root.Left != nil && root.Left.Val == root.Val {left = l + 1}
    if root.Right != nil && root.Right.Val == root.Val {right = r + 1}
    current := right + left
    if current > res {res = current}
    if right > left {return right}
    return left
}
