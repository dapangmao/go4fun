### 749. Shortest Completing Word

- level: 2-medium

```go
func shortestCompletingWord(licensePlate string, words []string) string {
    var ldic = make(map[rune]int)
    for _, x := range strings.ToLower(licensePlate) {
        if x >= 97 && x <= 122 { 
            ldic[x]++
        }
    }
    var res string
    for _, word := range words {
        var dic = make(map[rune]int)
        for _, x := range(word) {
            dic[x]++
        }
        for k, v := range ldic {
            value, ok := dic[k]
            if !ok || value < v {goto end}
        }
        if len(res) == 0 || len(word) < len(res) {res = word}
        end:
    }
    return res
}
```

### 747. Min Cost Climbing Stairs

- level: 1-easy

```go
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    var dp = make([]int, n)
    copy(dp, cost)
    for i:=2; i<n; i++ {
        if dp[i-2] < dp[i-1] {
            dp[i] += dp[i-2]
        } else {
            dp[i] += dp[i-1]
        }
    }
    if dp[n-1] < dp[n-2] {return dp[n-1]}
    return dp[n-2]
}
```

### 745. Find Smallest Letter Greater Than Target

- level: 1-easy

```go
func nextGreatestLetter(letters []byte, target byte) byte {
    var res byte = letters[0]
    for _, b := range letters {
        if b > target && (res <= target || b - target < res - target) {
                res = b
            } 
    }
    return res
}
```

### 739. Daily Temperatures

- level: 2-medium

```go
func dailyTemperatures(temperatures []int) []int {
	var n = len(temperatures)
	var res, stack []int
	for i:=0; i<n; i++ {
		res = append(res, 0)
	}
    for i:=n-1; i>=0; i-- {
        for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
            stack = stack[:len(stack)-1]
        }
        if len(stack) > 0 {
            res[i] = stack[len(stack)-1] - i
        }
        stack = append(stack, i)
    }
    return res
}
```

### 737. Sentence Similarity II

- level: 2-medium

```go
func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool {
    if len(words1) != len(words2) {return false}
    visited, dic := make(map[string]int), make(map[string][]string)
    count := 1
    for _, pair := range pairs {
        a, b := pair[0], pair[1]
        visited[a] = -count
        count++
        visited[b] = -count
        count++
        dic[a] = append(dic[a], b)
        dic[b] = append(dic[b], a)
    }
    var group int
    for k, v := range visited {
        if v > 0 {continue}
        group++
        queue := []string{k}
        for len(queue) > 0 {
            first := queue[0]
            queue = queue[1:]
            visited[first] = group
            for _, x := range dic[first] {
                if visited[x] < 0 {
                    queue = append(queue, x)
                    visited[x] = group
                }
            }
        }
    }
    for i:=0; i<len(words1); i++ {
        w1, w2 := words1[i], words2[i]
        if visited[w1] != visited[w2] {return false}
    }
    return true
}
```

### 734. Sentence Similarity

- level: 1-easy

```go
// empty data structure is OK for a go map
func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
    dic := make(map[string][]string)
    for _, pair := range pairs {
        a, b := pair[0], pair[1]
        dic[a] = append(dic[a], b)
        dic[b] = append(dic[b], a)
    }
    n, m := len(words1), len(words2)
    if n != m {return false}
    for i:=0; i<n; i++ {
        w1, w2 := words1[i], words2[i]
        if w1 == w2 {goto next}
        for _, x := range dic[w1] {
            if x == w2 {goto next}
        }
        return false
        next:
    }
    return true
}
```

### 733. Flood Fill

- level: 1-easy

```go
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    queue := [][]int{[]int{sr, sc}}
    n, m := len(image), len(image[0])
    visited := make([][]bool, n)
    for i:=0; i<n; i++ {
        visited[i] = make([]bool, m)
    }
    original := image[sr][sc]
    for len(queue) > 0 {
        r, c := queue[0][0], queue[0][1]
        queue = queue[1:]
        if visited[r][c] || image[r][c] != original {continue}
        image[r][c] = newColor
        visited[r][c] = true
        if r > 0 {queue = append(queue, []int{r-1, c})}
        if r + 1 < n {queue = append(queue, []int{r+1, c})}
        if c > 0 {queue = append(queue, []int{r, c-1})}
        if c + 1 < m {queue = append(queue, []int{r, c+1})}
    }
    return image  
}
```

### 728. Self Dividing Numbers

- level: 1-easy

```go
func selfDividingNumbers(left int, right int) []int {
    res := []int{}
    for x:=left; x<=right; x++ {
        y := x
        for y > 0 {
            current := y % 10
            if current == 0 || x % current != 0 {
                goto end
            }
            y /= 10
        }
        res = append(res, x)
        end:
    }
    return res
}
```

### 724. Find Pivot Index

- level: 1-easy

```go
func pivotIndex(nums []int) int {
    n := len(nums)
    if n == 0 {return -1}
    left, right := make([]int, n), make([]int, n)
    left[0], right[n-1] = nums[0], nums[n-1]
    for i:=1; i<n-1; i++ {
        left[i] += left[i-1] + nums[i]
    }
    for i:=n-2; i>=0; i-- {
        right[i] += right[i+1] + nums[i]
    }
    for i:=0; i<n; i++ {
        l := 0
        if i > 0 {
            l = left[i-1]
        }
        r := 0
        if i < n-1 {
            r = right[i+1]
        }
        if l == r {
            return i 
        }
    }
    return -1
}
```

### 722. Remove Comments

- level: 2-medium

```go
func removeComments(source []string) []string {
	res := []string{}
	inBlock := false
	for _, line := range source {
		var current string
		right := strings.Index(line, "*/")
		double := strings.Index(line, "//")
		left := strings.Index(line, "/*")
		if !inBlock && left >= 0 && right >= 0 {
            current = line[:left] + line[right+2:]
		} else if inBlock && right >= 0 {
			current = line[right+2:len(line)]
			inBlock = false
		} else if !inBlock && double >= 0 {
			current = line[:double]
		} else if !inBlock && left >= 0 {
			current = line[:left]
			inBlock = true
		} else if !inBlock {
			current = line
		}
        if len(current) > 0 {
		    res = append(res, current)
        }
	}
	return res
}
```

### 720. Longest Word in Dictionary

- level: 1-easy

```go
func longestWord(words []string) string {
    set := make(map[string]int)
    for _, x := range words {
        set[x] = len(x)
    }
    ans := ""
    for _, word := range words {
        if set[word] < set[ans] || (set[word] == set[ans] && word >= ans) {
            continue
        willAdd := true
        for k:=1; k<set[word]; k++ {
            if _, ok := set[word[:k]]; !ok {
                willAdd = false
                break
            } 
        }
        if willAdd {ans = word}
    }
    return ans
}
```

### 718. Maximum Length of Repeated Subarray

- level: 2-medium

```go
func findLength(A []int, B []int) int {
    n, m := len(A), len(B)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
    } 
    for i:=n-1; i>=0; i-- {
        for j:=m-1; j>=0; j-- {
            if A[i] == B[j] {
                dp[i][j] = dp[i+1][j+1] + 1
            }
        }
    }
    res := -1
    for _, x := range dp {
        for _, y := range x {
            if y > res {
                res = y
            }
        }
    }
    return res
}
```

### 717. 1-bit and 2-bit Characters

- level: 1-easy

```go
func isOneBitCharacter(bits []int) bool {
    i, n := 0, len(bits)
    for i < n - 1{
        i += bits[i] + 1
    }
    return i == n - 1
}
```

### 714. Best Time to Buy and Sell Stock with Transaction Fee

- level: 2-medium

```go
func maxProfit(prices []int, fee int) int {
    n := len(prices)
    cash, hold := 0, -prices[0]
    var currentCash, currentHold int
    for i:=1; i<n; i++ {
        currentCash = hold + prices[i] - fee
        if currentCash > cash {
            cash = currentCash
        }
        currentHold = cash - prices[i]
        if currentHold > hold {
            hold = currentHold
        }
    }
    return cash
}
```

### 712. Minimum ASCII Delete Sum for Two Strings

- level: 2-medium

```go
func minimumDeleteSum(s1 string, s2 string) int {
	n, m := len(s1), len(s2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][m] = dp[i+1][m] + int(s1[i])
	}
	for i := m - 1; i >= 0; i-- {
		dp[n][i] = dp[n][i+1] + int(s2[i])
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				opt1, opt2 := dp[i+1][j]+int(s1[i]), dp[i][j+1]+int(s2[j])
				if opt1 >= opt2 {
					dp[i][j] = opt2
				} else {
					dp[i][j] = opt1
				}
			}
		}
	}
	return dp[0][0]
}
```

### 697. Degree of an Array

- level: 1-easy

```go
func findShortestSubArray(nums []int) int {
    left, right, count := make(map[int]int), make(map[int]int), make(map[int]int)
    for i, x := range nums {
        if _, ok := left[x]; !ok {
            left[x] = i
        }
        right[x] = i
        l, ok := count[x] 
        if ok {
            count[x] = l + 1
        } else {
            count[x] = 1
        }
    }     
    ans := len(nums)
    degree := -1
    for _, v := range count {
        if v > degree {
            degree = v
        }
    }
    for x := range count {
        current := right[x] - left[x] + 1
        if count[x] == degree && current < ans {
            ans = current
        }
    }
    return ans
}
```

### 695. Max Area of Island

- level: 1-easy

```go
func maxAreaOfIsland(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    var res int
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if grid[i][j] != 1 {continue}
            queue := [][]int{[]int{i, j}}
            grid[i][j] = -1
            current := 0
            for len(queue) > 0 {
                var pop = queue[0]
                queue = queue[1:]
                var a, b = pop[0], pop[1]
                current++
                var options = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
                for _, option := range options {
                    var r = a + option[0]
                    var c = b + option[1]
                    if r >=0 && r < n && c >= 0 && c < m && grid[r][c] == 1 {
                        queue = append(queue, []int{r, c})
                        grid[r][c] = -1
                    }
                }
            }
            if current > res {res = current}
        }
    }
    return res
}
```

### 688. Knight Probability in Chessboard

- level: 2-medium

```go
func makeMatrix(N int) [][]float64 {
	res := make([][]float64, N)
	for i := range res {
		res[i] = make([]float64, N)
	}
	return res
}

func knightProbability(N int, K int, r int, c int) float64 {
	opts := [][]int{{2,1},{2,-1},{-2,1},{-2,-1},{1,2},{1,-2},{-1,2},{-1,-2}}
	dp := makeMatrix(N)
    dp[r][c] = 1
	for i:=0; i<K; i++ {
		current := makeMatrix(N)
		for j, x := range dp {
			for k, y := range x {
				for _, arr := range opts {
					row, col := j + arr[0], k + arr[1]
					if row >= 0 && row < N && col >=0 && col < N {
						current[row][col] += y / 8.0
					}
				}
			}
		}
        dp = current
	}
	var sum float64 
	for _, row := range dp {
		for _, val := range row {
			sum += val
		}
	}
	return sum
}
```

### 687. Longest Univalue Path

- level: 1-easy

```go
var res int

func longestUnivaluePath(root *TreeNode) int {
    dfs(root)
    return res
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l := dfs(root.Left)
    r := dfs(root.Right)
    left, right := 0, 0
    if root.Left != nil && root.Left.Val == root.Val {left = l + 1}
    if root.Right != nil && root.Right.Val == root.Val {right = r + 1}
    current := right + left
    if current > res {res = current}
    if right > left {return right}
    return left
}
```

### 661. Image Smoother

- level: 1-easy

```go
func imageSmoother(M [][]int) [][]int {
    n, m := len(M), len(M[0])
    res := make([][]int, n)
    for i := range M {
        res[i] = make([]int, m)
    }
    var current, count int
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            current = 0
            count = 0
            for _, l := range []int{-1, 0, 1} {
                for _, k := range []int{-1, 0, 1} {
                    r, c := i+l, j+k
                    if r >=0 && r < n && c >= 0 && c <m {
                        current += M[r][c]
                        count++
                    }
                }
            }
            res[i][j] = int(math.Floor(float64(current/count)))
        }
    }
    return res
}
```

### 657. Judge Route Circle

- level: 1-easy

```go
func judgeCircle(moves string) bool {
    var v, h = 0, 0
    for _, x := range moves {
        switch x {
            case 'U': v++
            case 'D': v--
            case 'L': h--
            default: h++
        }
    }
    return v == 0 && h == 0
}
```

### 645. Set Mismatch

- level: 1-easy

```go
func findErrorNums(nums []int) []int {
    var res = []int{-1, -1}
    var dic = make(map[int]int)
    var sum, should = 0, 0
    for i, x := range nums {
        dic[x]++
        should += i+1
        sum += x
        if dic[x] > 1 {res[0] = x}
    }
    res[1] = res[0] + should - sum
    return res
}
```

### 637. Average of Levels in Binary Tree

- level: 1-easy

```go
var dic = make(map[int][]int)
var max int

func averageOfLevels(root *TreeNode) []float64 {
    var res []float64
    dfs(root, 0)
    for k:=0; k<=max; k++ {
        v := dic[k]
        res = append(res, float64(v[1]) / float64(v[0]))
	}
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {return}
	if val, ok := dic[level]; ok {
		dic[level] = []int{val[0] + 1, val[1] + root.Val}
	} else {
		dic[level] = []int{1, root.Val}
	}
    if level > max {max = level}
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
```

### 636. Exclusive Time of Functions

- level: 2-medium

```go
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
```

### 594. Longest Harmonious Subsequence

- level: 1-easy

```go
func findLHS(nums []int) int {
    m := make(map[int]int)
    for _, x := range(nums) {
        if _, ok := m[x]; ok {
            m[x]++
        } else {
            m[x] = 1
        }
    }
    var res int
    for k := range m {
        if val, ok := m[k-1]; ok {
            option := m[k] + val
            if option > res {
                res = option
            }
        }
    }
    return res
}
```

### 387. First Unique Character in a String

- level: 1-easy

```go
func firstUniqChar(s string) int {
    m := make(map[rune]int)
    for _, x := range s {
        if val, ok := m[x]; ok {
            m[x] = val + 1
        } else {
            m[x] = 1
        }
    }
    for i, x := range s {
        if m[x] == 1 {return i}
    }
    return -1
}
```

### 187. Repeated DNA Sequences

- level: 2-medium

```go
func findRepeatedDnaSequences(s string) []string {
    var dic = make(map[string]int)
    for i:=0; i<=len(s)-10; i++ {
        dic[s[i:i+10]]++
    }
    var res = []string{}
    for k, v := range dic {
        if v > 1 {
            res = append(res, k)
        }
    }
    return res
}
```

### 165. Compare Version Numbers

- level: 2-medium

```go
func compareVersion(version1 string, version2 string) int {
	v1split := strings.Split(version1, ".")
	v2split := strings.Split(version2, ".")
	n, m := len(v1split), len(v2split)
	max := n
	if m > n {max = m}
	for i:=0; i<max; i++ {
		v1current, v2current := 0, 0
		if i < n {
			v1current, _ = strconv.Atoi(v1split[i])
		}
		if i < m {
			v2current, _ = strconv.Atoi(v2split[i])
		}
		if v1current > v2current {
			return 1
		}
		if v1current < v2current {
			return -1
		}
	}
	return 0
}
```

### 155. Min Stack

- level: 1-easy

```go
type MinStack struct {
    stack []int
    minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    if len(this.minStack) == 0 || x <= this.minStack[len(this.minStack)-1]{
        this.minStack = append(this.minStack, x) 
    }
}

func (this *MinStack) Pop()  {
    var pop = this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    if pop == this.GetMin() {
       this.minStack = this.minStack[:len(this.minStack)-1] 
    }
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}
```

### 150. Evaluate Reverse Polish Notation

- level: 2-medium

```go
func evalRPN(tokens []string) int {
    stack := []int{}
    for _, token := range tokens {
        t, err := strconv.Atoi(token)
        if len(stack) == 0 || err == nil {
            stack = append(stack, t)
            continue
        } 
        n := len(stack)
        e1, e2 := stack[n-2], stack[n-1]
        stack = stack[:n-2]
        var current int
        if token == "+" {
            current = e1 + e2
        } else if token == "-" {
            current = e1 - e2
        } else if token == "*" {
            current = e1 * e2 
        } else {
            current = e1 / e2
        }
        stack = append(stack, current)
    }
    return stack[0]
}
```

### 146. LRU Cache

- level: 3-hard

```go
type LRUCache struct {
    capacity int
    cache map[int]int
    order []int
    pos map[int]int
}

func Constructor(capacity int) LRUCache {
    var c = make(map[int]int)
    var o = []int{}
    var p = make(map[int]int)
    return LRUCache{capacity: capacity, cache: c, order: o, pos: p}
}

func (this *LRUCache) Get(key int) int {
    if val, ok := this.cache[key]; ok {
        var p = this.pos[key]
        this.order = append(this.order[:p], this.order[p+1:]...)
        this.order = append(this.order, key)
        this.pos[key] = len(this.cache) - 1
        return val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int)  {
    var current = this.Get(key)
    if current != -1 {
        this.cache[key] = value
        return
    }
    if len(this.cache) == this.capacity {
        var current = this.order[0]
        delete(this.cache, current)
        this.order = this.order[1:]
    }
    this.order = append(this.order, key)
    this.cache[key] = value
    this.pos[key] = len(this.cache) - 1
}
```

### 121. Best Time to Buy and Sell Stock

- level: 1-easy

```go
func maxProfit(prices []int) int {
    res := 0
    lowest := 324324
    for _, x := range prices {
        if x > lowest && x - lowest > res {
            res = x - lowest
        }
        if x < lowest {
            lowest = x
        }
    }
    return res
}
```

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

### 100. Same Tree

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
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil{
        return true
    }
    if p == nil || q == nil || p.Val != q.Val {
        return false
    }
    return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
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

### 58. Length of Last Word

- level: 1-easy

```go
func lengthOfLastWord(s string) int {
    n := len(s)
    for i:=len(s)-1; i>=0; i-- {
        if s[i] != ' ' {
            break
        }
        n -= 1 
    }
    for i:=n-1; i>=0; i-- {
        if s[i] == ' ' {
            return n-1-i   
        }
    }
    return n
}
```

### 38. Count and Say

- level: 1-easy

```go
func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		res = nextState(res)
	}
	return res
}

func nextState(s string) string {
	res, prev := "", ""
	count := 0
	for _, x := range s {
		current := string(x)
		if prev != current  {
			if count != 0 {
				res += strconv.Itoa(count) + prev
			}
			prev = current
			count = 1
		} else {
			count += 1
		}
	}
	return res + strconv.Itoa(count) + prev
}
```

### 21. Merge Two Sorted Lists

- level: 1-easy

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := new(ListNode) 
    dummy := head
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val{
            dummy.Next = l1
            l1 = l1.Next
        } else {
            dummy.Next = l2
            l2 = l2.Next
        }
        dummy = dummy.Next
    }
    if l1 != nil {
        dummy.Next = l1
    } else {
        dummy.Next = l2
    }
    return head.Next
}
```

### 14. Longest Common Prefix

- level: 1-easy

```go
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    first := strs[0]
    for i:=0; i<len(first); i++ {
        for _, str := range strs[1:] {
            if i >= len(str) || str[i] != first[i] {
                return first[:i]
            }
        }
    }
    return first
}
```

### 6. ZigZag Conversion

- level: 2-medium

```go
func convert(s string, numRows int) string {
    if numRows == 0 {return ""}
    if numRows == 1 {return s}
    buffer := make([]string, numRows)
    down, j := true, 0
    for _, x := range s {
        buffer[j] += string(x)
        if down {
            j += 1
        } else {
            j -= 1
        }
        if j >= numRows {
            j = numRows - 2
            down = false
        } 
        if j < 0 {
            j = 1
            down = true
        }
    }
    return strings.Join(buffer, "")
}
```

### 1. Two Sum

- level: 1-easy

```go
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, x := range nums {
        if v, ok := m[x]; ok {
            return []int{v, i}
        } else {
            m[target-x] = i
        }
    }
    return make([]int, 2)
}
```


