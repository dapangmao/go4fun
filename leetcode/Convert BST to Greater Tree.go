func convertBST(root *TreeNode) *TreeNode {
    var s = 0
    dfs(root, &s)
    return root
}

func dfs (root *TreeNode, sum *int) {
    if root == nil {return}
    dfs(root.Right, sum)
    *sum += root.Val
    root.Val = *sum
    dfs(root.Left, sum)
}

// functional but 3 times slower
func convertBST(root *TreeNode) *TreeNode {
    var sum = 0
    var dfs func(root *TreeNode) 
    dfs = func(root *TreeNode) {
        if root == nil {return}
        dfs(root.Right)
        sum += root.Val
        root.Val = sum
        dfs(root.Left)
    }
    dfs(root)
    return root
}
