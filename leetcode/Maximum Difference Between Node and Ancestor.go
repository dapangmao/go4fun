func maxAncestorDiff(root *TreeNode) int {
	res := -123213
	var dfs func(root *TreeNode, parents []int)
	dfs = func(root *TreeNode, parents []int) {
		if root == nil {return}
		current := getMaxAbs(root.Val, parents)
		if  current > res {res = current}
		next := append(parents, root.Val)
		dfs(root.Left, next)
		dfs(root.Right, next)
	}
	dfs(root, []int{})
	return res
}

func getMaxAbs(needle int, hays []int ) int {
	res := -1
	var current int
	for _, h := range hays {
		current = h - needle
		if current < 0 {current = -current}
		if current > res {res = current}
	}
	return res
}
