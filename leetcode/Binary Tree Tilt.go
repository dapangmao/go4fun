var res int 
func findTilt(root *TreeNode) int {
    sum(root)
    return res
}

func sum(root *TreeNode) int {
    if root == nil {return 0}
    if root.Left == nil && root.Left == nil {return root.Val}
    l, r :=  sum(root.Left), sum(root.Right)
    var diff = l - r
    if diff > 0 {
        res += diff
    } else {
        res -= diff
    }    
    return root.Val + r + l
}
