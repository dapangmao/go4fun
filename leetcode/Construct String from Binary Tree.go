/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
    if t == nil {return ""}
    var res = fmt.Sprintf("%v", t.Val)
    if t.Left != nil && t.Right != nil {
        res += fmt.Sprintf("(%v)(%v)", tree2str(t.Left), tree2str(t.Right))
    } else if t.Left != nil {
        res += fmt.Sprintf("(%v)", tree2str(t.Left))
    } else if t.Right != nil {
        res += fmt.Sprintf("()(%v)", tree2str(t.Right))
    }
    return res
}


func tree2str(t *TreeNode) string {
    if t == nil {return ""}
    var res strings.Builder
    fmt.Fprintf(&res, "%d", t.Val)
    if t.Left != nil && t.Right != nil {
        fmt.Fprintf(&res, "(%v)(%v)", tree2str(t.Left), tree2str(t.Right))
    } else if t.Left != nil {
        fmt.Fprintf(&res, "(%v)", tree2str(t.Left))
    } else if t.Right != nil {
        fmt.Fprintf(&res, "()(%v)", tree2str(t.Right))
    }
    return res.String()
}
