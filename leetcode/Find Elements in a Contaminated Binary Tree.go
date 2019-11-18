def dfs( root, val):
    return dfs(root.left, val*2+1) + [val] + dfs(root.right, val*2+2) if root else []


class FindElements:

    def __init__(self, root: TreeNode):
        self.data = dfs(root, 0)

    def find(self, target: int) -> bool:
        return target in self.data
