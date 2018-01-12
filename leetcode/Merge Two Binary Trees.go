func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {return nil}
    if t1 != nil && t2 != nil {
        return &TreeNode{t1.Val+t2.Val, mergeTrees(t1.Left, t2.Left), mergeTrees(t1.Right, t2.Right)}
    } 
    if t1 != nil {
        return &TreeNode{t1.Val, mergeTrees(t1.Left, nil), mergeTrees(t1.Right, nil)}
    } 
    return &TreeNode{t2.Val, mergeTrees(nil, t2.Left), mergeTrees(nil, t2.Right)}
}
