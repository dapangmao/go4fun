var n []int

func sortedArrayToBST(nums []int) *TreeNode {
    n = nums
    return dfs(0, len(nums)-1)
}

func dfs(lo, hi int) *TreeNode {
    if lo > hi {
        return nil
    }
    var i = (lo + hi) / 2
    return &TreeNode{
        n[i],
        dfs(lo, i-1),
        dfs(i+1, hi),
    }
}
