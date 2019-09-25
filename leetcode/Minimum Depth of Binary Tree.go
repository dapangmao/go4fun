var res = 424242

func minDepth(root *TreeNode) int {
    dfs(root, 0)
    return res
}

func dfs(root *TreeNode, path int) {
    if root == nil {
        if path < res {res = path}
        return 
    }
    dfs(root.Left, path+1)
    dfs(root.Right, path+1)
}
