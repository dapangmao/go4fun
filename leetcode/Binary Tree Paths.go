var res []string
func binaryTreePaths(root *TreeNode) []string {
    dfs(root, []string{})
    return res
}

func dfs(root *TreeNode, path []string) {
    if root == nil { return }
    if root.Left == nil && root.Right == nil {
        path = append(path, strconv.Itoa(root.Val))
        res = append(res, strings.Join(path, "->"))
        return
    }
    path = append(path, strconv.Itoa(root.Val))
    dfs(root.Left, path)
    dfs(root.Right, path)
}
