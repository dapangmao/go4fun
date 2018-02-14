var res []int
func largestValues(root *TreeNode) []int {
    dfs(root, 0)
    return res
}

func dfs(root *TreeNode, level int) {
    if root == nil {return}
    if len(res) <= level {
        res = append(res, root.Val)
    } else if root.Val > res[level] {
        res[level] = root.Val
    }
    dfs(root.Left, level+1)
    dfs(root.Right, level+1)
}
