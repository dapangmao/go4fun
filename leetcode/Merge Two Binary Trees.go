func dfs(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {return nil}
    var root TreeNode
    if t1 != nil && t2 != nil {
        root = TreeNode{t1.Val+t2.Val, dfs(t1.Left, t2.Left), dfs(t1.Right, t2.Right)}
    } else if t1 != nil {
        root = TreeNode{t1.Val, dfs(t1.Left, nil), dfs(t1.Right, nil)}
    } else {
        root = TreeNode{t2.Val, dfs(nil, t2.Left), dfs(nil, t2.Right)}
    }
    return &root
}


func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    return dfs(t1, t2)
}
