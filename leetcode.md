### 112. Path Sum

- level: 1-easy

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return sum - root.Val == 0
    }
    return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}
```

### 66. Plus One

- level: 1-easy

```go
func plusOne(digits []int) []int {
    n := len(digits)
    for i:=n-1; i>=0; i-- {
        if digits[i] != 9 {
            digits[i] += 1
            return digits
        }
        digits[i] = 0
    }
    return append([]int{1}, digits...)
}
```


