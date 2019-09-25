func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    // use two appends to insert an element
    return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...) 
}
