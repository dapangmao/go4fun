var arr = []int{}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	val := root.Val
	if  len(arr) == 0 {
		arr = []int{val}
	} else if len(arr) == 1 {
		if arr[0] > val {
			arr = []int{arr[0], val}
		} else if arr[0] < val {
			arr = []int{arr[0], val}
		}
	} else {
		if arr[0] < val && val < arr[1] {
			arr = []int{val, arr[1]}
		} else if val > arr[1] {
			arr = []int{arr[1], val}
		}
	}
    dfs(root.Left)
	dfs(root.Right)
}

func findSecondMinimumValue(root *TreeNode) int {
    arr = []int{}
	dfs(root)
	if len(arr) < 2 {return -1}
	return arr[0]
}
