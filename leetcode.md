### 777. Toeplitz Matrix

- level: 1-easy

```go
func isToeplitzMatrix(matrix [][]int) bool {
    n, m := len(matrix), len(matrix[0])
    for i, j:= 0, 0; i<n; i++ {
        if !check(i,j,n,m, matrix) {return false}
    }
    for i, j:= 0, 1; j<m; j++ {
        if !check(i,j,n,m, matrix) {return false}
    }
    return true
}

func check(i, j, n, m int, matrix [][]int) bool {
    var dic = make(map[int]int)
    for i < n && j < m {
        dic[matrix[i][j]]++
        if len(dic) > 1 {return false}
        i++ 
        j++
    }
    return true
}
```

### 767. Prime Number of Set Bits in Binary Representation

- level: 1-easy

```go
func countPrimeSetBits(L int, R int) int {
    var res int
    for i:=L; i<=R; i++ {
        var x = countOnes(i)
        if x == 2 || x == 3 || x == 5 || x == 7 || x == 11 || x == 13 || x == 17 || x == 19  {res++}
    }
    return res
}

func countOnes(n int) int {
    var count int
    for n > 0 {
        count += n & 1
        n >>= 1
    }
    return count
}
```

### 762. Find Anagram Mappings

- level: 1-easy

```go
func anagramMappings(A []int, B []int) []int {
    var m = make(map[int]int)
    for i, x := range B {
        m[x] = i
    }
    for i, x := range A {
        B[i] = m[x]
    }
    return B
}
```

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

### 748. Largest Number At Least Twice of Others

- level: 1-easy

```go
func dominantIndex(nums []int) int {
	var max, max2 = -12423421, -12423421
	var i = -1
	for k, num := range nums {
		if num > max {
			max, max2 = num, max
			i = k
		} else if num > max2 {
			max2 = num
		}
	}
	if max >= max2*2 {
		return i
	}
	return -1
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

### 696. Count Binary Substrings

- level: 1-easy

```go
func countBinarySubstrings(s string) int {
    var dp = []int{1}
    var n = len(s)
    for i:=1; i<n; i++ {
        if s[i-1] != s[i] {
            dp = append(dp, 1)
        } else {
            dp[len(dp)-1]++
        }
    }
    var res = 0
    for i:=1; i<len(dp); i++ {
        if dp[i-1] < dp[i] {
            res += dp[i-1]
        } else {
            res += dp[i]
        }
    }
    return res
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

### 686. Repeated String Match

- level: 1-easy

```go
func repeatedStringMatch(A string, B string) int {
    var increase = A
    var res int = 1
    for {
        if strings.Contains(increase, B) {
            return res
        } else if len(increase) > 2*len(B) {
            break
        } else {
            increase += A
            res++
        }
    }
    return -1
}
```

### 682. Baseball Game

- level: 1-easy

```go
func calPoints(ops []string) int {
    var stack []int
    for _, op := range ops {
        var n = len(stack)
        if op == "+" {
            stack = append(stack, stack[n-1], stack[n-2])
        } else if op == "C" {
            stack = stack[:n-1]
        } else if op == "D" {
            stack = append(stack, 2*stack[n-1])
        } else {
            opint, _ := strconv.Atoi(op)
            stack = append(stack, opint)
        }
    }
    var res = 0
    for _, x := range stack {
        res += x
    }
    return res
}
```

### 674. Longest Continuous Increasing Subsequence

- level: 1-easy

```go
func findLengthOfLCIS(nums []int) int {
    var res, current int
    for i, x := range nums {
        if i == 0 || x > nums[i-1]{
            current++
        } else {
            if current > res {
                res = current
            }
            current = 1
        }
    }
    if current > res {return current}
    return res
}
```

### 671. Second Minimum Node In a Binary Tree

- level: 1-easy

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

### 643. Maximum Average Subarray I

- level: 1-easy

```go
func findMaxAverage(nums []int, k int) float64 {
    var sum int
    var max int = -12312312
    for i:=0; i<len(nums); i++ {
        sum += nums[i]
        if i >= k {
            sum -= nums[i-k]
        }
        if i >= k-1 && sum > max {max = sum}
    }
    return float64(max) / float64(k)
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

### 633. Sum of Square Numbers

- level: 1-easy

```go
func judgeSquareSum(c int) bool {
    var cand []int{0}
    for i:=1; i*i<=c; i++ {
        cand = append(cand, i*i)
    }
    i, j := 0, len(cand)-1
    for i <= j {
        sum := cand[i] + cand[j]
        if sum == c {return true}
        if sum < c {
            i++
        } else {
            j--
        }
    }
    return false
}
```

### 628. Maximum Product of Three Numbers

- level: 1-easy

```go
func maximumProduct(nums []int) int {
    var n = len(nums)
    sort.Ints(nums)
    var cand1 = nums[n-1] * nums[n-2] * nums[n-3] 
    var cand2 = nums[n-1] * nums[0] * nums[1]
    if cand1 > cand2 {return cand1}
    return cand2
}
```

### 617. Merge Two Binary Trees

- level: 1-easy

```go
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {return nil}
    if t1 != nil && t2 != nil {
        return &TreeNode{t1.Val+t2.Val, mergeTrees(t1.Left, t2.Left), mergeTrees(t1.Right, t2.Right)}
    } 
    if t1 != nil {
        return &TreeNode{t1.Val, mergeTrees(t1.Left, nil), mergeTrees(t1.Right, nil)}
    } 
    return &TreeNode{t2.Val, mergeTrees(nil, t2.Left), mergeTrees(nil, t2.Right)}
}
```

### 606. Construct String from Binary Tree

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
func tree2str(t *TreeNode) string {
    if t == nil {return ""}
    var res = fmt.Sprintf("%v", t.Val)
    if t.Left != nil && t.Right != nil {
        res += fmt.Sprintf("(%v)(%v)", tree2str(t.Left), tree2str(t.Right))
    } else if t.Left != nil {
        res += fmt.Sprintf("(%v)", tree2str(t.Left))
    } else if t.Right != nil {
        res += fmt.Sprintf("()(%v)", tree2str(t.Right))
    }
    return res
}
```

### 605. Can Place Flowers

- level: 1-easy

```go
func canPlaceFlowers(flowerbed []int, n int) bool {
    var l = len(flowerbed)
    if n == 0 {return true}
    for i:=0; i<l; i++ {
        if flowerbed[i] == 1 {continue}
        if (i==0 || flowerbed[i-1] == 0 ) && (i == l-1 || flowerbed[i+1] == 0) {
            n--
            flowerbed[i] = 1
        }
        if n <= 0 {return true}
    }
    return false
}
```

### 599. Minimum Index Sum of Two Lists

- level: 1-easy

```go
func findRestaurant(list1 []string, list2 []string) []string {
    var dic = make(map[string]int)
    for i, x := range list1 {
        dic[x] = -i-1
    }
    for i, x := range list2 {
        if val, ok := dic[x]; ok {
            dic[x] = -val + i + 1
        }
    }
    res, count := []string{}, 13123213
    for k, v := range dic {
        if v < 0 {continue}
        if v == count { 
            res = append(res, k) 
        } else if v < count {
            res = []string{k}
            count = v
        }
    }
    return res
}
```

### 598. Range Addition II

- level: 1-easy

```go
func maxCount(m int, n int, ops [][]int) int {
    var x, y = m, n
    for _, op := range ops {
        if op[0] < x {x = op[0]}
        if op[1] < y {y = op[1]}
    }
    return x * y
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

### 581. Shortest Unsorted Continuous Subarray

- level: 1-easy

```go
func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    var should = make([]int, n)
    copy(should, nums)
    sort.Ints(nums)
    var i, j = 0, n-1
    for ; i<n; i++ {
        if nums[i] != should[i] {break}
    }
    for ; j>=0; j-- {
        if nums[j] != should[j] {break}
    }
    var res = j - i + 1
    if res > 0 {return res}
    return 0
}
```

### 575. Distribute Candies

- level: 1-easy

```go
func distributeCandies(candies []int) int {
    var num = len(candies) / 2
    var m = make(map[int]int)
    for _, x := range candies {
        m[x]++
    }
    if len(m) < num {return len(m)}
    return num
}
```

### 566. Reshape the Matrix

- level: 1-easy

```go
func matrixReshape(nums [][]int, r int, c int) [][]int {
    var n, m = len(nums), len(nums[0])
    if n * m != r * c {return nums}
    var res = make([][]int, r)
    for i := range res {res[i] = make([]int, c)}
    var k, l int
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            res[k][l] = nums[i][j]
            l++
            if l == c {
                k++
                l = 0
            }
        }
    }
    return res
}
```

### 563. Binary Tree Tilt

- level: 1-easy

```go
var res int 
func findTilt(root *TreeNode) int {
    sum(root)
    return res
}

func sum(root *TreeNode) int {
    if root == nil {return 0}
    if root.Left == nil && root.Left == nil {return root.Val}
    l, r :=  sum(root.Left), sum(root.Right)
    var diff = l - r
    if diff > 0 {
        res += diff
    } else {
        res -= diff
    }    
    return root.Val + r + l
}
```

### 557. Reverse Words in a String III

- level: 1-easy

```go
func reverseWords(s string) string {
	var i int
  var bytes = []byte(s)   // not {}
	for j, x := range s {
		if x == ' ' {
			rw(bytes[:], i, j-1) // have to use slice
			i = j + 1
		}
	}
	rw(bytes[:], i, len(s)-1) // have to use slice
	return string(bytes)
}

func rw(b []byte, i, j int)  {
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}
```

### 554. Brick Wall

- level: 2-medium

```go
func leastBricks(wall [][]int) int {
    var n = len(wall)
    var max int
    var dic = prep(wall)
    for _, val := range dic {
        if val > max {max = val}
    }
    return n - max
}

func prep(wall [][]int) map[int]int {
    var res = make(map[int]int)
    for _, w := range wall {
        var current int
        for i:=0; i<len(w)-1; i++ {
            current += w[i]
            res[current]++
        }
    }
    return res
}
```

### 543. Diameter of Binary Tree

- level: 1-easy

```go
var res int

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {return res}
    dfs(root)
    return res
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l, r := dfs(root.Left), dfs(root.Right)
    if l + r > res {res = l + r}
    if l > r {return l + 1}
    return r + 1
}
```

### 532. K-diff Pairs in an Array

- level: 1-easy

```go
func findPairs(nums []int, k int) int {
    if k < 0 {return 0}
    var dic = make(map[int]bool)
    var res int
    for _, x := range nums {
        if k != 0 {
            if _, ok := dic[x]; ok {continue}
            if _, ok := dic[x+k]; ok {res++}
            if _, ok := dic[x-k]; ok {res++}
            dic[x] = false
        } else {
            if val, ok := dic[x]; ok {
                if !val {res++}
                dic[x] = true
            } else {
                dic[x] = false
            }
        }
    }
    return res
}
```

### 530. Minimum Absolute Difference in BST

- level: 1-easy

```go
var res = 13123213
var prev = new(TreeNode)


func getMinimumDifference(root *TreeNode) int {
    dfs(root)
    return res
}

func dfs(root *TreeNode) {
    if root == nil {return}
    dfs(root.Left)
    if prev != nil && res > root.Val - prev.Val {
        res = root.Val - prev.Val
    }
    prev = root
    dfs(root.Right)
}
```

### 525. Contiguous Array

- level: 2-medium

```go
func findMaxLength(nums []int) int {
    var n = len(nums)
    if n < 0 {return 0}
    var dic = make(map[int]int)
    dic[0] = -1
    res, count := 0, 0
    for i:=0; i<n; i++ {
        if nums[i] == 0 {
            count--
        } else {
            count++
        }
        if val, ok := dic[count]; ok {
            var current = i - val
            if current > res {res = current}
        } else {
            dic[count] = i
        }
    }
    return res
}
```

### 520. Detect Capital

- level: 1-easy

```go
func detectCapitalUse(word string) bool {
    var allLower, allUpper, firstUpper = 0, 0, 0
    var n = len(word)
    for i, c := range word {
        if c >= 'A' && c <= 'Z' {
            allUpper++
            if i == 0 {firstUpper = 1}
        } else {
            allLower++
        }
    } 
    if allLower == n || allUpper == n || (firstUpper == 1 && allLower == n-1) {return true}
    return false
}
```

### 506. Relative Ranks

- level: 1-easy

```go
func findRelativeRanks(nums []int) []string {
	max := -1
	for _, num := range nums {
		if num > max {max = num}
	}
	var dp = make([]int, max+1)
	for i, num := range nums {
		dp[num] = i+1
	}
	var res = make([]string, len(nums))
	var top = []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	var j = 4
	for i:=max; i>=0; i-- {
		if dp[i] == 0 {continue}
		if len(top) > 0 {
			res[dp[i]-1] = top[0]
			top = top[1:]
		} else {
            res[dp[i]-1] = fmt.Sprint(j)
			j++
		}
	}
	return res
}
```

### 500. Keyboard Row

- level: 1-easy

```go
func findWords(words []string) []string {
    var dict = make(map[rune]int)
    var count = 0
    for _, c := range "qwertyuiop asdfghjkl zxcvbnm" {
        if count == ' ' {
            count++
        } else {
            dict[c] = count
        }
    }
    var res []string
    for _, w := range words {
        var current = strings.ToLower(w)
        var first = dict[current[0]]
        for _, c := range current {
            if first != dict[c] {goto end}
        }
        res = append(res, w)
        end:
    }
    return res
}
```

### 492. Construct the Rectangle

- level: 1-easy

```go
func constructRectangle(area int) []int {
    var sqrt = int(math.Sqrt(float64(area)))
    for i:=sqrt; i>1; i-- {
        var div = area / i
        if i * div == area  {return []int{div, i}}
    }
    return []int{area, 1}
}
```

### 482. License Key Formatting

- level: 1-easy

```go
func licenseKeyFormatting(S string, K int) string {
    var res []byte
    var count = K
    for i:=len(S)-1; i>=0; i-- {
        var c = S[i]
        if c == '-' {continue}
        res = append([]byte{c}, res...)
        count--
        if i > 0 && count == 0 {
            res = append([]byte{'-'}, res...)
            count = K
        }
    }
    count = 0
    for _, c := range res {
        if c != '-' {break}
        count++
    }
    return strings.ToUpper(string(res[count:]))
}
```

### 463. Island Perimeter

- level: 1-easy

```go
func islandPerimeter(grid [][]int) int {
    var n, m = len(grid), len(grid[0])
    var res int
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if grid[i][j] != 1 {continue}
            if i+1 >= n || grid[i+1][j] == 0 {res++}
            if i-1 < 0 || grid[i-1][j] == 0 {res++}
            if j+1 >= m || grid[i][j+1] == 0 {res++}
            if j-1 < 0 || grid[i][j-1] == 0 {res++}
        }
    }
    return res
}
```

### 461. Hamming Distance

- level: 1-easy

```go
func hammingDistance(x int, y int) int {
    var a = x ^ y
    var res int
    for a > 0 {
        if a % 2 == 1 {res++}
        a >>= 1
    }
    return res
}
```

### 448. Find All Numbers Disappeared in an Array

- level: 1-easy

```go
func findDisappearedNumbers(nums []int) []int {
    for _, x := range nums {
        if x < 0 {x = -x}
        if nums[x-1] > 0 {
            nums[x-1] = -nums[x-1]    
        }
    }
    var res []int
    for i, x := range nums {
        if x > 0 {res = append(res, i+1)}
    }
    return res
}
```

### 434. Number of Segments in a String

- level: 1-easy

```go
func countSegments(s string) int {
    var res = 0
    var prev = ' '
    for _, x := range s + " " {
        if x == ' ' && prev != ' ' {
            res++
        }
        prev = x
    }
    return res
}
```

### 414. Third Maximum Number

- level: 1-easy

```go
func thirdMax(nums []int) int {
    const m = -9223372036854775808
    var max1, max2, max3 = m, m, m
    for _, num := range nums {
        if num == max1 || num == max2 || num == max3 {continue}
        if num > max1 {
           max3 = max2
           max2 = max1
           max1 = num
        } else if num > max2 {
           max3 = max2
           max2 = num
        } else if num > max3 {
           max3 = num
        }
    }
    if max3 == m {return max1}
    return max3
}
```

### 409. Longest Palindrome

- level: 1-easy

```go
func longestPalindrome(s string) int {
    var set = make(map[rune]bool)
    var res = 0
    for _, x := range s {
        if _, ok := set[x]; ok {
            delete(set, x)
            res += 2
        } else {
            set[x] = true
        }
    }
    if len(set) > 0 {return res + 1}
    return res
}
```

### 389. Find the Difference

- level: 1-easy

```go
func findTheDifference(s string, t string) byte {
    var m = make([]int, 26)
    for _, x := range []byte(s) {
        m[x - 'a']++
    }
    for _, x := range []byte(t) {
        m[x - 'a']--
        if m[x - 'a'] == -1 {return x}
    }
    return 0
}
```

### 387. First Unique Character in a String

- level: 1-easy

```go
func firstUniqChar(s string) int {
    m := make(map[rune]int)
    for _, x := range s {m[x]++}
    for i, x := range s {
        if m[x] == 1 {return i}
    }
    return -1
}
```

### 367. Valid Perfect Square

- level: 1-easy

```go
func isPerfectSquare(num int) bool {
    var sqrt = int(math.Sqrt(float64(num)))
    return sqrt * sqrt == num
}
```

### 345. Reverse Vowels of a String

- level: 1-easy

```go
func reverseVowels(s string) string {
    var bs = []byte(s)
    i, j := 0, len(bs)-1
    for i < j {
        for i < j && !isVowel(bs[i]) {i++}
        for i < j && !isVowel(bs[j]) {j--}
        bs[i], bs[j] = bs[j], bs[i]
        i++
        j--
    }
    return string(bs)
}

func isVowel(a byte) bool {
    if a == 'a' || a == 'e' || a == 'u' || a == 'o' || a == 'i' {return true}
    if a == 'A' || a == 'E' || a == 'U' || a == 'O' || a == 'I' {return true}
    return false
}
```

### 292. Nim Game

- level: 1-easy

```go
func canWinNim(n int) bool {
    return n % 4 != 0 
}
```

### 290. Word Pattern

- level: 1-easy

```go
func wordPattern(pattern string, str string) bool {
    var words = strings.Split(str, " ")
    if len(words) != len(pattern) {return false}
    var m1, m2 = make(map[rune]string), make(map[string]rune)
    for i, x := range pattern {
        word := words[i]
        if w, ok := m1[x]; ok{
            if w != word {return false}
        } else {
            m1[x] = word
        }
        if r, ok := m2[word]; ok {
            if r != x {return false}
        } else {
            m2[word] = x
        }
    }
    return true
}
```

### 268. Missing Number

- level: 1-easy

```go
func missingNumber(nums []int) int {
    var sum, n = 0, len(nums)
    for _, x := range nums {sum += x}
    return (n+1) *n/2 - sum
}
```

### 242. Valid Anagram

- level: 1-easy

```go
func isAnagram(s string, t string) bool {
    var m = make(map[rune]int)
    for _, x := range s {
        m[x]++
    }
    for _, x := range t {
        m[x]--
    }
    for _, v := range m {
        if v != 0 {return false}
    }
    return true
}
```

### 226. Invert Binary Tree

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
func invertTree(root *TreeNode) *TreeNode {
    dfs(root)
    return root
}

func dfs(root *TreeNode) {
    if root == nil {return}
    root.Left, root.Right = root.Right, root.Left
    dfs(root.Left)
    dfs(root.Right)
}
```

### 219. Contains Duplicate II

- level: 1-easy

```go
func containsNearbyDuplicate(nums []int, k int) bool {
    var m = make(map[int]int)
    for i, x := range nums {
        if val, ok := m[x]; ok {
            if i-val <= k {return true}
        } 
        m[x] = i
    }
    return false
}
```

### 205. Isomorphic Strings

- level: 1-easy

```go
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {return false}
    return isOnewayIsomorphic(s, t) && isOnewayIsomorphic(t, s)
}

func isOnewayIsomorphic(s, t string) bool {
    var dic = make(map[byte]byte)
    for i, x := range []byte(s){
        if val, ok := dic[x]; ok {
            if val != t[i] {return false}
        } else {
            dic[x] = t[i]
        }
    }
    return true
}
```

### 204. Count Primes

- level: 1-easy

```go
func countPrimes(n int) int {
    var dp = make([]bool, n)
    if n <= 2 {return 0}
    var res int
    for i:=2; i<n; i++ {
        if dp[i] == true {continue}
        for j:=2; i*j<n; j++ {
            dp[i*j] = true
        }
        res++
    }
    return res
}
```

### 189. Rotate Array

- level: 1-easy

```go
func rotate(nums []int, k int)  {
    var n = len(nums)
    k %= n
    reverse(nums[:], 0, n-1)
    reverse(nums[:], 0, k-1)
    reverse(nums[:], k, n-1)
}

func reverse(nums []int, i, j int) {
    for i < j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
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

### 125. Valid Palindrome

- level: 1-easy

```go
func isPalindrome(s string) bool {
	var bytes []int
	for _, x := range []byte(s) {
		var cand1 = x - 'a'
		var cand2 = x - 'A'
		var cand3 = x - '0'
		if cand1 <= 25 {
			bytes = append(bytes, int(cand1))
		} else if cand2 <= 25 {
			bytes = append(bytes, int(cand2))
		} else if cand3 <= 9 {
			bytes = append(bytes, int(cand3))
		}
	}
	i, j := 0, len(bytes)-1
	for i < j {
		if bytes[i] != bytes[j] {return false}
		i++
		j--
	}
	return true
}
```

### 122. Best Time to Buy and Sell Stock II

- level: 1-easy

```go
func maxProfit(prices []int) int {
	var res = 0
	for i:=1; i<len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
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

### 119. Pascal's Triangle II

- level: 1-easy

```go
func getRow(rowIndex int) []int {
    var res = []int{1}
    for i:=1; i<=rowIndex; i++ {
        res = append(res, 1)
        var prev = 1
        for j:=1; j<len(res)-1; j++ {
            current := res[j]
            res[j] += prev
            prev = current
        }
    }
    return res
}
```

### 118. Pascal's Triangle

- level: 1-easy

```go
func generate(numRows int) [][]int {
    var res [][]int
    var prev = []int{1}
    if numRows > 0 {res = append(res, prev)}
    for i:=1; i<numRows; i++ {
        var n = len(prev)
        var current = make([]int, n+1)
        current[0], current[n] = 1, 1
        for j:=1; j<n; j++ {
            current[j] = prev[j-1] + prev[j]
        }
        res = append(res, current)
        prev = current
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

### 111. Minimum Depth of Binary Tree

- level: 1-easy

```go
var res = 424242

func minDepth(root *TreeNode) int {
    dfs(root, 0)
    return res
}

func dfs(root *TreeNode, path int) {
    if root == nil {
        if path < res {res = path}
        return 
    }
    dfs(root.Left, path+1)
    dfs(root.Right, path+1)
}
```

### 104. Maximum Depth of Binary Tree

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
 
func maxDepth(root *TreeNode) int {
    if root == nil {return 0}
    l, r := maxDepth(root.Left), maxDepth(root.Right)
    if l > r {return 1 + l}
    return 1 + r
}
```

### 101. Symmetric Tree

- level: 1-easy

```go
func isSymmetric(root *TreeNode) bool {
    if root == nil {return true}
    return check(root.Left, root.Right)  
}

func check(t1 *TreeNode, t2 *TreeNode) bool {
    if t1 == nil && t2 == nil {return true}
    if t1 == nil || t2 == nil {return false}
    if t1.Val != t2.Val {return false}
    return check(t1.Left, t2.Right) && check(t1.Right, t2.Left)
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

### 88. Merge Sorted Array

- level: 1-easy

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    n--
    m--
    var j = n + m + 1
    for m >= 0 && n >= 0 {
        if nums1[m] >= nums2[n] {
            nums1[j] = nums1[m]
            m--
        } else {
            nums1[j] = nums2[n]
            n--
        }
        j--
    }
    for n >= 0 {
        nums1[j] = nums2[n]
        n--
        j--
    }
}
```

### 71. Simplify Path

- level: 2-medium

```go
func simplifyPath(path string) string {
    var stack = strings.Split(path, "/")
    var res []string
    var current int
    for i:=len(stack)-1; i>=0; i-- {
        if stack[i] == "." || len(strings.TrimSpace(stack[i])) == 0 {continue}
        if stack[i] == ".." {
            current++
        } else if current > 0 {
            current--
        } else {
            res = append([]string{stack[i]}, res...)
        }
    }
    return "/" + strings.Join(res, "/")
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

### 56. Merge Intervals

- level: 2-medium

```go
/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    var res []Interval
    for i, x := range intervals {
        n := len(res) 
        if i == 0 || x.Start > res[n-1].End {
            res = append(res, x)
        } else if x.End > res[n-1].End {
            res[n-1].End = x.End
        }        
    }
    return res
}
```

### 49. Group Anagrams

- level: 2-medium

```go
func groupAnagrams(strs []string) [][]string {
    var dic = make(map[string][]string)
    for _, str := range strs {
        var key = makeKey(str)
        dic[key] = append(dic[key], str)
    }
    var res [][]string
    for _, val := range dic {
        res = append(res, val)
    }
    return res
}

func makeKey(s string) string {
    var bucket = make([]int, 26)
    for _, x := range s {
        bucket[x - 'a']++
    }
    var res string
    for i, x := range bucket {
        if x <= 0 {continue}
        res += fmt.Sprintf("%d-%d/", i, x)
    }
    return res
}
```

### 43. Multiply Strings

- level: 2-medium

```go
func multiply(num1 string, num2 string) string {
    var len1, len2 = len(num1), len(num2)
    var product = make([]int, len1+len2)
    for i := len1 - 1; i >= 0; i-- {
        for  j := len2 - 1; j >= 0; j-- {
            var index = len1 + len2 - i - j - 2
            product[index] += int(num1[i] - '0') * int(num2[j] - '0')
            product[index + 1] += product[index] / 10;
            product[index] %= 10;
        }
    }
    var stringBuilder []byte
    for i := len(product) - 1; i > 0; i-- {
        if len(stringBuilder) == 0 && product[i] == 0 {
            continue
        }
        stringBuilder = append(stringBuilder, byte(product[i]))
    }
    stringBuilder = append(stringBuilder, byte(product[0]))
    var res string
    for _, x := range stringBuilder {res += fmt.Sprint(x)}
    return res
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

### 20. Valid Parentheses

- level: 1-easy

```go
func isValid(s string) bool {
    var pairs = map[rune]rune{')': '(', ']': '[', '}': '{' }
    var stack []rune
    var popout rune
    for _, x := range s {
        if x == '(' || x == '[' || x == '{' {
            stack = append(stack, x)
        } else {
            var n = len(stack)
            if n == 0 {return false}
            popout, stack = stack[n-1], stack[:n-1]
            if popout != pairs[x] {return false}
        }
    }
    return len(stack) == 0
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

### 13. Roman to Integer

- level: 1-easy

```go
func romanToInt(s string) int {
    var res = 0
    if strings.Contains(s, "IV")  {res -= 2}
    if strings.Contains(s, "IX")  {res -= 2}
    if strings.Contains(s, "XL") {res -= 20}
    if strings.Contains(s, "XC") {res -= 20}
    if strings.Contains(s, "CD") {res -= 200}
    if strings.Contains(s, "CM") {res -= 200}
    for _, c := range s {
        switch c {
            case 'M': 
            res += 1000
            case 'D':
            res += 500
            case 'C':
            res += 100
            case 'L':
            res += 50
            case 'X':
            res += 10
            case 'V':
            res += 5
            case 'I':
            res += 1
        }
    }
    return res
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


