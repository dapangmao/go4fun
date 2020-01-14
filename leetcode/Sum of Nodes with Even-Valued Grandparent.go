func sumEvenGrandparent(root *TreeNode) int {
    var res int
    var dfs func(root, parent, grandparent *TreeNode)
    dfs = func(root, parent, grandparent *TreeNode) {
        if root == nil {return}
        if grandparent != nil && grandparent.Val % 2 == 0 {
            res += root.Val
        }
        dfs(root.Left, root, parent)
        dfs(root.Right, root, parent)
    }
    dfs(root, nil, nil)
    return res
}
