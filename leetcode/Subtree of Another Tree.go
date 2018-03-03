func isSubtree(s *TreeNode, t *TreeNode) bool {
    if check(s, t) {return true}
    if s == nil {return false}
    return isSubtree(s.Left, t) || isSubtree(s.Right, t) 
}

func check(s *TreeNode, t *TreeNode) bool {
    if t == nil && s == nil{
        return true
    }
    if s == nil || t == nil {return false}
    if t.Val != s.Val {return false}
    return check(s.Left, t.Left) && check(s.Right, t.Right)
}
