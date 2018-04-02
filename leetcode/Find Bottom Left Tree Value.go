/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    prev := []*TreeNode{root}
    for len(prev) > 0 {
        current := []*TreeNode{}
        for _, n := range prev {
            if n.Left != nil {
                current = append(current, n.Left)
            }
            if n.Right != nil {
                current = append(current, n.Right)
            }
        }
        if len(current) == 0 {return prev[0].Val}
        prev = current
    }
    return -1
}
