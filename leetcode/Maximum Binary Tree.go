/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {return nil}
    j, max := -1, -1
    for i, num := range nums {
        if num > max {
            max = num
            j = i
        }
    }
    return &TreeNode{
        max, 
        constructMaximumBinaryTree(nums[:j]),
        constructMaximumBinaryTree(nums[j+1:])} // weird behavior
}
