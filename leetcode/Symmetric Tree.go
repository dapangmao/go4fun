func isSymmetric(root *TreeNode) bool {
    if root == nil {return true}
    return check(root.Left, root.Right)  
}

func check(t1 *TreeNode, t2 *TreeNode) bool {
    if t1 == nil && t2 == nil {return true}
    if t1 == nil || t2 == nil {return false}
    if t1.Val != t2.Val {return false}
    return check(t1.Left, t2.Right) && check(t1.Right, t2.Left)
}
