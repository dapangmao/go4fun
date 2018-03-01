func trimBST(root *TreeNode, L int, R int) *TreeNode {
    if root == nil {return nil}
    l, r := trimBST(root.Left, L, R), trimBST(root.Right, L, R) 
    if root.Val > R {return l}
    if root.Val < L {return r}
    root.Left = l
    root.Right = r
    return root
}


