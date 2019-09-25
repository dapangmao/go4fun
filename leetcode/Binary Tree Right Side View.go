/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
 // composite date type does not need *root.Val
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res = []int{root.Val}
	var current = []*TreeNode{root}
	for {
		var prev []*TreeNode
		for _, n := range current {
			if n.Left != nil {
				prev = append(prev, n.Left)
			}
			if n.Right != nil {
				prev = append(prev, n.Right)
			}
		}
		if len(prev) == 0 {break}
		res = append(res, prev[len(prev)-1].Val)
		current = prev
	}
	return res
}
