func insertIntoBST(root *TreeNode, val int) *TreeNode {
    
    var dfs func(root *TreeNode, val int) 
    dfs = func(root *TreeNode, val int) {
        if root == nil {return}
        if root.Val >= val {
            if root.Left == nil {
                root.Left = &TreeNode{val, nil, nil}
            } else {
                dfs(root.Left, val)
            }
        } else {
            if root.Right == nil {
                root.Right = &TreeNode{val, nil, nil}
            } else {
                dfs(root.Right, val)
            }
        }
    }
    dfs(root, val)
    return root
}
