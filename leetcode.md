### 799. Minimum Distance Between BST Nodes

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
var min = 23424234
func minDiffInBST(root *TreeNode) int {
    inorder(root)
    return min
}

func inorder(root *TreeNode) []int {
    if root == nil {return []int{}} 
    var res = inorder(root.Left)
    if len(res) > 0 && root.Val - res[len(res)-1] < min {
        min = root.Val - res[len(res)-1]
    }
    res = append(res, root.Val)
    res = append(res, inorder(root.Right)...)
    return res
}
```

### 797. Rabbits in Forest

- level: 2-medium

```go
func numRabbits(answers []int) int {
    dic := make(map[int]int)
    for _, a := range answers {
        dic[a]++
    }
    var res int
    for k, v := range dic {
        var current = v / (k+1)
        if v % (k+1) != 0 {current += 1}
        res += current * (k+1)
    }
    return res
}
```

### 782. Jewels and Stones

- level: 1-easy

```go
func numJewelsInStones(J string, S string) int {
    var dic = make(map[rune]int)
    for _, x := range J {
        dic[x] = 0
    }
    var res int
    for _, x := range S {
        if _, ok := dic[x]; ok {
            res++
        }
    }
    return res
}
```

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

### 768. Partition Labels

- level: 2-medium

```go
func partitionLabels(S string) []int {
    dic := make(map[rune]int)
    for i, x := range S {
        dic[x] = i
    }
    var arrs [][]int
    for i, x := range S {
        if val, ok := dic[x]; ok {
            if len(arrs) == 0 || i > arrs[len(arrs)-1][1] {
                arrs = append(arrs, []int{i, val})
            } else if val > arrs[len(arrs)-1][1] {
                arrs[len(arrs)-1][1] = val
            }
            delete(dic, x)
        }
    }
    var res []int
    for _, arr := range arrs {
        res = append(res, arr[1]-arr[0]+1)
    }
    return res
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

### 740. Delete and Earn

- level: 2-medium

```go
`
Given an array nums of integers, you can perform operations on the array.

In each operation, you pick any nums[i] and delete it to earn nums[i] points. After, you must delete every element equal to nums[i] - 1 or nums[i] + 1.

You start with 0 points. Return the maximum number of points you can earn by applying such operations.

Example 1:
Input: nums = [3, 4, 2]
Output: 6
Explanation: 
Delete 4 to earn 4 points, consequently 3 is also deleted.
Then, delete 2 to earn 2 points. 6 total points are earned.
Example 2:
Input: nums = [2, 2, 3, 3, 3, 4]
Output: 9
Explanation: 
Delete 3 to earn 3 points, deleting both 2's and the 4.
Then, delete 3 again to earn 3 points, and 3 again to earn 3 points.
9 total points are earned.
`

func deleteAndEarn(nums []int) int {
    if len(nums) == 0 {return 0}
    if len(nums) == 1 {return nums[0]}
    dic := make(map[int]int)
    max := -1
    for _, num := range nums {
        dic[num] += num
        if num > max {max = num}
    }
    dp := make([]int, max+1)
    for k, v := range dic {
        dp[k] = v
    }
    for i:=2; i<=max; i++ {
        if dp[i-1] > dp[i] + dp[i-2] {
            dp[i] = dp[i-1]
        } else {
            dp[i] += dp[i-2]
        }
    }
    return dp[max]
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

### 729. My Calendar I

- level: 2-medium

```go
type MyCalendar struct {
    starts []int
    ends []int
    n int
    sdic map[int]int
}

func Constructor() MyCalendar {
    return MyCalendar{[]int{}, []int{}, 0, map[int]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
    if _, ok := this.sdic[start]; ok {return false}
    var i = sort.SearchInts(this.starts, start)
    if  (i < this.n && end > this.starts[i]) || (i > 0  && start < this.ends[i-1]) {return false}
    this.starts = append(this.starts[:i], append([]int{start}, this.starts[i:]...)...)
    this.ends = append(this.ends[:i], append([]int{end}, this.ends[i:]...)...)
    this.n++
    this.sdic[start] = 0
    return true
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

### 693. Binary Number with Alternating Bits

- level: 1-easy

```go
func hasAlternatingBits(n int) bool {
    prev := n & 1
    n >>= 1
    for n > 0 {
        var current = n & 1
        if current == prev {return false}
        prev = current
        n >>= 1
    }
    return true
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

### 677. Map Sum Pairs

- level: 2-medium

```go
type MapSum struct {
    dic map[string]int
    lastInsert map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
    return MapSum{map[string]int{}, map[string]int{}}
}

func (this *MapSum) Insert(key string, val int)  {
    lastInsertVal, ok := this.lastInsert[key]
    for i:=1; i<=len(key); i++ {
        var current = key[:i]
        this.dic[current] += val
        if ok {this.dic[current] -= lastInsertVal}
    }
    this.lastInsert[key] = val
}


func (this *MapSum) Sum(prefix string) int {
    return this.dic[prefix]
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

### 665. Non-decreasing Array

- level: 1-easy

```go
func checkPossibility(nums []int) bool {
    used := false   
    var copies = make([]int, len(nums))
    copy(copies, nums)
    for i:=1; i< len(nums); i++ { 
        if nums[i] >= nums[i-1] { continue } // continue once > turns to <
        if used {
            return false
        } else {
            used = true
            nums[i-1] = nums[i]
            copies[i] = copies[i-1]
        }
    }
    unmet := 0
     for i:=1; i< len(nums); i++ { 
        if nums[i] < nums[i-1] {  
            unmet++
        }
        if copies[i] < copies[i-1] {  
            unmet++
        }
         if unmet > 1 {return false}
     }
    return true
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

### 654. Maximum Binary Tree

- level: 2-medium

```go
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
```

### 653. Two Sum IV - Input is a BST

- level: 1-easy

```go
var dic = make(map[int]bool)
var K int

func findTarget(root *TreeNode, k int) bool {
    K = k
    return dfs(root)
}

func dfs(root *TreeNode) bool {
    if root == nil {return false}
    if _, ok := dic[root.Val]; ok {return true}
    dic[K-root.Val] = true
    return dfs(root.Left) || dfs(root.Right)
}
```

### 650. 2 Keys Keyboard

- level: 2-medium

```go
func minSteps(n int) int {
    dp := make([]int, n+1)
    for i:=2; i<n+1; i++ {
        dp[i] = i
    }
    for i:=4; i<n+1; i++ {
        for j:=2; j<i/2+1; j++ {
            if i % j != 0 {continue}
            var current = dp[j] + i / j
            if current < dp[i] {dp[i] = current}
        }
    }
    return dp[n]
}
```

### 648. Replace Words

- level: 2-medium

```go
func replaceWords(dict []string, sentence string) string {
    sort.Slice(dict, func(i, j int) bool {
        return len(dict[i]) < len(dict[j])
    })
    var res []string
    var strs = strings.Split(sentence, " ")
    for _, str := range strs {
        var current = str
        for _, prefix := range dict {
            if strings.HasPrefix(str, prefix) {
                current = prefix
                break
            }
        }
        res = append(res, current)
    }
    return strings.Join(res, " ")
}
```

### 647. Palindromic Substrings

- level: 2-medium

```go
func countSubstrings(s string) int {
    count := func(left, right int) int {
        ans := 0
        for left >= 0 && right < len(s) && s[left] == s[right] {
            ans += 1
            left -= 1
            right += 1
        }
        return ans
    }
    var res int
    for i:=0; i<len(s); i++ {
        res += count(i, i) + count(i, i+1)
    }
    return res
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

### 611. Valid Triangle Number

- level: 2-medium

```go
func triangleNumber(nums []int) int {
    n := len(nums)
    if n < 2 {return 0}
    sort.Ints(nums)
    var res int
    for i:=0; i<n-2; i++ {
        for j:=i+1; j<n-1; j++ {
            for k:=j+1; k<n; k++ {
                if nums[i] + nums[j] > nums[k] {res++}
            }
        }   
    }
    return res
}
```

### 609. Find Duplicate File in System

- level: 2-medium

```go
func findDuplicate(paths []string) [][]string {
    dic := make(map[string][]string)
    for _, path := range paths {
        var current = strings.Split(path, " ")
        var prefix = current[0]
        for _, file := range current[1:] {
            var i = strings.Index(file, "(")
            var value = prefix + "/" + file[:i]
            var key = file[i+1:len(file)-1]
            dic[key] = append(dic[key], value)  
        }
    }
    var res [][]string
    for _, v := range dic {
        if len(v) < 2 {continue}
        res = append(res, v)
    }
    return res
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

### 565. Array Nesting

- level: 2-medium

```go
func arrayNesting(nums []int) int {
	var res int
	for i:=0; i<len(nums); i++ {
		var current = dfs(i, nums)
		if current > res {res = current}
	}
	return res
}

func dfs(i int, path []int) int {
	if path[i] < 0 {return 0}
	var current = path[i]
	path[i] = -1
	return 1 + dfs(current, path)
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

### 561. Array Partition I

- level: 1-easy

```go
func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    var res int
    for i, x := range nums {
        if i%2==0 {res += x}
    }
    return res
}
```

### 560. Subarray Sum Equals K

- level: 2-medium

```go
func subarraySum(nums []int, k int) int {
    dic := make(map[int]int)
    dic[0] = 1
    var current, res int
    for _, x := range nums {
        current += x
        if val, ok := dic[current-k]; ok {res += val}
        dic[current]++

    }
    return res
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

### 553. Optimal Division

- level: 2-medium

```go
func optimalDivision(nums []int) string {
    n := len(nums)
    if n == 0 {return ""}
    var res = strconv.Itoa(nums[0])
    if n == 1 {return res}
    if n == 2 {return fmt.Sprintf("%v/%v", res, nums[1])}
    res += "/("
	for i:=1; i<n-1; i++ {
		res += strconv.Itoa(nums[i]) + "/"
	}
    return fmt.Sprintf("%v%v)", res, nums[n-1])  
}
```

### 551. Student Attendance Record I

- level: 1-easy

```go
func checkRecord(s string) bool {
    n := len(s)
    hasA := false
    for i:=0; i<n; i++ {
        if s[i] == 'A' {
            if hasA {
                return false
            } else {
                hasA = true
            }
        }
        if i > 0 && i < n-1 && s[i] == 'L' && s[i-1] == 'L' && s[i+1] == 'L' {
            return false
        }
    }
    return true
}
```

### 547. Friend Circles

- level: 2-medium

```go
func findCircleNum(M [][]int) int {
    n, m := len(M), len(M[0])
    var visited = make(map[int]bool)

    graph := make(map[int][]int)
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if j > i {continue}  // don't give up those orpahns 
            if M[i][j] == 1 {
                graph[i] = append(graph[i], j)  
                graph[j] = append(graph[j], i) 
            }
        }
    }
    
    var dfs func(i int, graph map[int][]int)
    dfs = func(i int, graph map[int][]int) {
        if _, ok := visited[i]; ok {return}
        visited[i] = true
        for _, j := range graph[i] {
            dfs(j, graph)
        }
    }
    
    var res int
    for k, _ := range graph {
        if _, ok := visited[k]; ok {continue}
        dfs(k, graph)
        res++
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

### 537. Complex Number Multiplication

- level: 2-medium

```go
func complexNumberMultiply(a string, b string) string {
    partA, partB := strings.Split(a, "+"), strings.Split(b, "+")
    
    t := func(s string) int {
        res, _ := strconv.Atoi(s)
        return res
    }
    g := func(s string) int {
        res, _ := strconv.Atoi(s[:len(s)-1])
        return res
    }
  
    nonI := t(partA[0])*t(partB[0]) - g(partA[1])*g(partB[1])
    I := g(partA[1])*t(partB[0]) + t(partA[0])*g(partB[1])
    return fmt.Sprintf("%v+%vi", nonI, I)
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

### 515. Find Largest Value in Each Tree Row

- level: 2-medium

```go
var res []int
func largestValues(root *TreeNode) []int {
    dfs(root, 0)
    return res
}

func dfs(root *TreeNode, level int) {
    if root == nil {return}
    if len(res) <= level {
        res = append(res, root.Val)
    } else if root.Val > res[level] {
        res[level] = root.Val
    }
    dfs(root.Left, level+1)
    dfs(root.Right, level+1)
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

### 498. Diagonal Traverse

- level: 2-medium

```go
func findDiagonalOrder(matrix [][]int) []int {
    n, m := len(matrix), len(matrix[0])
    var dirs [][]int
    for j:=0; j<m; j++ {
        dirs = append(dirs, []int{0, j})
    }
    for i:=1; i<n; i++ {
        dirs = append(dirs, []int{i, m-1})
    }
    var res []int
    for k, dir := range dirs {
        i, j := dir[0], dir[1]
        var current []int
        for i < n && j >=0  {
            current = append(current, matrix[i][j])
            i++
            j--
        }
        if k%2 == 0 {reverse(current[:]))}
        res = append(res, current...)
    }
    return res
}
```

### 494. Target Sum

- level: 2-medium

```go
func findTargetSumWays(nums []int, S int) int {
    var dic = make(map[int]int)
    if len(nums) == 0 {return 0}
    dic[nums[0]]++
    dic[-nums[0]]++
    for i:=1; i<len(nums); i++ {
        var current = make(map[int]int)
        for k, v := range dic {
            current[k+nums[i]] += v
            current[k-nums[i]] += v
        }
        dic = current
    }
    if val, ok := dic[S]; ok {return val}
    return 0
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

### 485. Max Consecutive Ones

- level: 1-easy

```go
func findMaxConsecutiveOnes(nums []int) int {
    var max = -1
    var current int
    for _, x := range nums {
        if x == 0 {
            if current > max {max = current}
            current = 0
        } else {
            current++
        }
    }
    if current > max {return current}
    return max
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

### 476. Number Complement

- level: 1-easy

```go
func findComplement(num int) int {
    var res int
    i := 1
    for ;num > 0; num >>= 1 {
        if num & 1 == 0 {
            res += i
        }
        i *= 2
    }
    return res
}
```

### 475. Heaters

- level: 1-easy

```go
func findRadius(houses []int, heaters []int) int {
  sort.Ints(heaters)
	var radius = 0
	for _, x := range houses {
		var j = sort.SearchInts(heaters, x)
		if j < len(heaters) && x == heaters[j] {continue}
		var a, b = 1<<32, 1<<32
		if j-1 >= 0 {a = x - heaters[j-1] }
		if j < len(heaters) {b = heaters[j] - x}
		var current = a
		if b < a {current = b}
		if current > radius {radius = current}
	}
	return radius
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

### 462. Minimum Moves to Equal Array Elements II

- level: 2-medium

```go
func minMoves2(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    var median int
    if n % 2 == 0 {
        median = (nums[n/2-1] + nums[n/2]) / 2
    } else {
        median = nums[n/2]
    }
    var res int
    for _, num := range nums {
        var current = median - num
        if current < 0 {current = -current}
        res += current
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

### 459. Repeated Substring Pattern

- level: 1-easy

```go
func repeatedSubstringPattern(s string) bool {
    n := len(s)
    for i:=1; i <= n/2; i++ {
        if n%i != 0 {continue}
        var fragment = []byte(s[:i])
        var dup []byte
        for j:=0; j<n/i; j++ {
            dup = append(dup, fragment...)
        }
        if string(dup) == s {return true}
    } 
    return false
}
```

### 455. Assign Cookies

- level: 1-easy

```go
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    var j, res int
    for _, x := range g {
        for j < len(s) && s[j] < x {
            j++
        } 
        if j == len(s) {break}
        if s[j] >= x {
            j++
            res++
        }
    }
    return res
}
```

### 453. Minimum Moves to Equal Array Elements

- level: 1-easy

```go
func minMoves(nums []int) int {
    var min = 2 << 61
    var res int
    for _, num := range nums {
        if num < min {min = num}
    }
    for _, num := range nums {
        res += num - min
    }
    return res
}
```

### 451. Sort Characters By Frequency

- level: 2-medium

```go
func frequencySort(s string) string {
    var dic = make(map[byte]int)
    var bs = []byte(s)
    for _, x := range bs {
        dic[x]++
    }
    sort.Slice(bs, func(i, j int) bool {
        return float64(dic[bs[j]]) + float64(bs[j]) / 1000.0 < float64(dic[bs[i]]) + float64(bs[i]) / 1000.0 
    })
    return string(bs)
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

### 442. Find All Duplicates in an Array

- level: 2-medium

```go
func findDuplicates(nums []int) []int {
    var res []int
    for _, num := range nums {
        if num < 0 {num = -num}
        if nums[num-1] < 0 {
            res = append(res, num)
        } else {
            nums[num-1] = -nums[num-1]
        }
    }
    return res
}
```

### 438. Find All Anagrams in a String

- level: 1-easy

```go
func findAnagrams(s string, p string) []int {
    hash := make([]int, 26)
    var res []int
    n := len(p)
    if n > len(s) {return res}
    for _, x := range []byte(p) {
        hash[int(x-'a')]++
    }
    for _, x := range []byte(s[:n-1]) {
        hash[int(x-'a')]--
    }
    for i:=0; i<len(s)-n+1; i++{
        var inchar = s[i+n-1]
        hash[int(inchar-'a')]--
        if checkHash(hash) {res = append(res, i)}
        var outchar = s[i]
        hash[int(outchar-'a')]++
    }
    return res
}

func checkHash(hash []int) bool {
    for _, x := range hash {
        if x != 0 {return false}
    }
    return true
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

### 419. Battleships in a Board

- level: 2-medium

```go
func countBattleships(board [][]byte) int {
    var res int
    n, m := len(board), len(board[0])
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if board[i][j] != 'X' {continue}
            var queue = [][]int{{i, j}}  // save some codes
            var current []int
            for len(queue) > 0 {
                current, queue = queue[0], queue[1:]
                l, k := current[0], current[1]
                if l > 0 && board[l-1][k] == 'X' {queue = append(queue, []int{l-1, k})}
                if l < n-1 && board[l+1][k] == 'X' {queue = append(queue, []int{l+1, k})}
                if k > 0 && board[l][k-1] == 'X' {queue = append(queue, []int{l, k-1})}
                if k < m-1 && board[l][k+1] == 'X' {queue = append(queue, []int{l, k+1})}
                board[l][k] = 'Z'
            }
            res++
        }
    }
    return res
}
```

### 415. Add Strings

- level: 1-easy

```go
func addStrings(num1 string, num2 string) string {
    i, j := len(num1)-1, len(num2)-1
    var sum []int
    var prev int
    for i >= 0 || j >= 0 {
        var current = prev
        if i >= 0 {
            current += int(num1[i] - '0')
            i--
        }
        if j >= 0 {
            current += int(num2[j] - '0')
            j--
        }
        sum = append(sum, current%10)
        prev = current / 10
    }
    if prev > 0 {sum = append(sum, prev)}
    var res string
    for i:=len(sum)-1; i>=0; i-- {
        res += strconv.Itoa(sum[i])
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

### 404. Sum of Left Leaves

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

var res int

func sumOfLeftLeaves(root *TreeNode) int {
    dfs(root, false)
    return res
}

func dfs(root *TreeNode, isLeft bool) {
    if root == nil {return}
    if isLeft {res += root.Val}
    dfs(root.Left, true)
    dfs(root.Right, false)
}
```

### 401. Binary Watch

- level: 1-easy

```go
func readBinaryWatch(num int) []string {
    var res []string
    for i:=0; i<12; i++ {
        for j:=0; j<60; j++ {
            if num == countOne(i) + countOne(j) {
                res = append(res, fmt.Sprintf("%d:%02d", i, j))
            }
        }
    }
    return res
}

func countOne(n int) int {
    var res int 
    for n > 0 {
        if n & 1 == 1 {res++}
        n >>= 1
    }
    return res
}
```

### 392. Is Subsequence

- level: 2-medium

```go
func isSubsequence(s string, t string) bool {
    n, m := len(s), len(t)
    if n == 0 {return true}
    var j int
    for i:=0; i<m; i++ {
        if s[j] == t[i] {
            j++
        }
        if j == n {return true}
    }
    return false
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

### 383. Ransom Note

- level: 1-easy

```go
func canConstruct(ransomNote string, magazine string) bool {
    var dic = make(map[rune]int)
    for _, x := range magazine {
        dic[x]++
    }
    for _, x := range ransomNote {
        dic[x]--
        if dic[x] < 0 {return false}
    }
    return true
}
```

### 371. Sum of Two Integers

- level: 1-easy

```go
func getSum(a, b int) int {
    if b == 0 {return a}
    return getSum(a ^ b, (a & b) << 1)
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

### 347. Top K Frequent Elements

- level: 2-medium

```go
func topKFrequent(nums []int, k int) []int {
    dic := make(map[int]int)
    max := -1
    for _, num := range nums {
        dic[num]++
        if dic[num] > max {max = dic[num]}
    }
    var bucket = make([][]int, max+1)
    for k, v := range dic {
        bucket[v] = append(bucket[v], k)
    }
    var res []int
    for i:=max; i>=0; i-- {
        if len(bucket[i]) == 0 {continue}
        res = append(res, bucket[i]...)
        if len(res) == k {return res}        
    }
    return res
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

### 342. Power of Four

- level: 1-easy

```go
func isPowerOfFour(num int) bool {
    var x = 1
    for x <= num {
        if x == num {return true}
        x *= 4
    }
    return false
}
```

### 326. Power of Three

- level: 1-easy

```go
func isPowerOfThree(n int) bool {
    if n == 0 {return false}
    if n == 1 {return true}
    if n % 3 != 0 {return false}
    return isPowerOfThree(n/3)
}
```

### 318. Maximum Product of Word Lengths

- level: 2-medium

```go
func maxProduct(words []string) int {
    hash := make(map[string]int)
    var max int
    for _, word := range words {
        for k, v := range hash {
            if v >= len(word) {continue}
            var current = numIfDiff(k, word)
            if current > v {
                hash[k] = v
                if current*len(k) > max {max = current*len(k)}
            }
        }
        hash[word] = 0
    }
    return max
}

func numIfDiff(a, b string) int {
    var bucket [26]int
    for _, x := range []byte(a) {
        bucket[int(x-'a')]++
    }
    for _, x := range []byte(b) {
        if bucket[int(x-'a')] > 0 {return -1}
    }
    return len(b)
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

### 287. Find the Duplicate Number

- level: 2-medium

```go
func findDuplicate(nums []int) int {
    sort.Ints(nums)
    for i:=1; i<len(nums); i++ {
        if nums[i] == nums[i-1] {return nums[i]}
    }
    return -1
}
```

### 283. Move Zeroes

- level: 1-easy

```go
func moveZeroes(nums []int)  {
    var j int 
    for _, x := range nums {
        if x != 0 {
            nums[j] = x 
            j++
        }
    }
    for i:=j; i<len(nums); i++ {
        nums[i] = 0
    }
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

### 264. Ugly Number II

- level: 2-medium

```go
func nthUglyNumber(n int) int {
    var ugly = make([]int, n)
    ugly[0] = 1
    var index2, index3, index5 int
    factor2, factor3, factor5 := 2, 3, 5
    for i:=1; i<n; i++ {
        var min = factor2
        if factor3 < min {min = factor3}
        if factor5 < min {min = factor5}
        ugly[i] = min
        if factor2 == min {
            index2++
            factor2 = 2*ugly[index2]
        }
        if factor3 == min {
            index3++
            factor3 = 3*ugly[index3]
        }
        if factor5 == min {
            index5++
            factor5 = 5*ugly[index5]
        }
    }
    return ugly[n-1]
}
```

### 263. Ugly Number

- level: 1-easy

```go
func isUgly(num int) bool {
    for num > 1 && num % 2 == 0 {
        num /= 2
    }
    for num > 1 && num % 3 == 0 {
        num /= 3
    }
    for num > 1 && num % 5 == 0 {
        num /= 5
    }
    return num == 1
}
```

### 258. Add Digits

- level: 1-easy

```go
func addDigits(num int) int {
    for num > 9 {
        var current int
        for num > 0 { 
            current += num % 10
            num /= 10
        }
        num = current
    }
    return num
}
```

### 257. Binary Tree Paths

- level: 1-easy

```go
var res []string
func binaryTreePaths(root *TreeNode) []string {
    dfs(root, []string{})
    return res
}

func dfs(root *TreeNode, path []string) {
    if root == nil { return }
    if root.Left == nil && root.Right == nil {
        path = append(path, strconv.Itoa(root.Val))
        res = append(res, strings.Join(path, "->"))
        return
    }
    path = append(path, strconv.Itoa(root.Val))
    dfs(root.Left, path)
    dfs(root.Right, path)
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

### 238. Product of Array Except Self

- level: 2-medium

```go
func productExceptSelf(nums []int) []int {
    total, n, j := 1, len(nums), 0
    var zeroCount int
    for i, num := range nums {
        if num == 0 {
            zeroCount++
            j = i
        } else {
            total *= num
        }
    }
    res := make([]int, n)
    if zeroCount > 1 {
        return res
    }
    if zeroCount == 1 {
        res[j] = total
        return res
    }
    for i, x := range nums {
        res[i] = total / x
    }
    return res
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

### 216. Combination Sum III

- level: 2-medium

```go
var K, N int
var res [][]int

func combinationSum3(k int, n int) [][]int {
    N = n
    K = k
    dfs(0, 0, []int{})
    return res
}

func dfs(i, sum int, path []int) {
    if len(path) == K {
        if sum == N {res = append(res, path)}
        return
    }
    for j:=i+1; j<=9; j++ {
        if j + sum <= N {
            var current = append(path, j)
            dfs(j, sum+j, current)
        }
    }
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

### 202. Happy Number

- level: 1-easy

```go
func isHappy(n int) bool {
    var dic = make(map[int]bool)    
    for {
        n = part(n)
        if n == 1 {return true}
        if _, ok := dic[n]; ok {return false}
        dic[n] = true
    }
    return false
}

func part(n int) int {
    var res int
    for ; n > 0; n /= 10 {
        res += (n%10) * (n%10)
    }
    return res
}
```

### 200. Number of Islands

- level: 2-medium

```go
func numIslands(grid [][]byte) int {
    var res int
    n, m := len(grid), len(grid[0])
    if n == 0 || m == 0 {return 0}
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if grid[i][j] != '1' {continue}
            var queue = [][]int{[]int{i, j}}
            for len(queue) > 0 {
                r, c := queue[0][0], queue[0][1]
                queue = queue[1:]
                if r+1 < n && grid[r+1][c] == '1' {
                    queue = append(queue, []int{r+1, c})
                }
                if r-1 >= 0 && grid[r-1][c] == '1' {
                    queue = append(queue, []int{r-1, c})
                }
                if c-1 >= 0 && grid[r][c-1] == '1' {
                    queue = append(queue, []int{r, c-1})
                }
                if c+1 < m && grid[r][c+1] == '1' {
                    queue = append(queue, []int{r, c+1})
                }
                grid[r][c] = '2'
            }
            res++
        }
    }
    return res
}
```

### 199. Binary Tree Right Side View

- level: 2-medium

```go
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
```

### 198. House Robber

- level: 1-easy

```go
func rob(nums []int) int {
    n := len(nums)
    if n == 0 {return 0}
    if n == 1 {return nums[0]}
    if nums[1] < nums[0] {nums[1] = nums[0]} 
    for i:=2; i<n; i++ {
        if nums[i-1] > nums[i-2] + nums[i] {
            nums[i] = nums[i-1]
        } else {
            nums[i] += nums[i-2]
        }
    }
    return nums[n-1]
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

### 179. Largest Number

- level: 2-medium

```go
func largestNumber(nums []int) string {
    var strs []string
    for _, num := range nums {
        strs = append(strs, strconv.Itoa(num))
    }
    sort.Slice(strs, func(i, j int) bool {
        return strs[i] + strs[j] > strs[j] + strs[i]
    })
    if strs[0] == "0" {return "0"}
    return strings.Join(strs, "")
}
```

### 172. Factorial Trailing Zeroes

- level: 1-easy

```go
func trailingZeroes(n int) int {
    if n <= 1 {return 0}
    var res int
    for i:=n; i>0 && i%5==0; i=i/5 {
        res++
    }
    return trailingZeroes(n-1) + res
}
```

### 171. Excel Sheet Column Number

- level: 1-easy

```go
func titleToNumber(s string) int {
    n := len(s)
    var res int
    j := 1
    for i:= n-1; i>=0; i-- {
        var current = int(s[i] - 'A') + 1
        current *= j
        res += current
        j *= 26
    }
    return res
}
```

### 166. Fraction to Recurring Decimal

- level: 2-medium

```go
func fractionToDecimal(numerator int, denominator int) string {
	var res string
	if numerator * denominator < 0 {res += "-"}
	if numerator < 0 {numerator = - numerator}
	if denominator < 0 {denominator = -denominator}
	res += strconv.Itoa(numerator / denominator)
	var current = numerator % denominator
	if current == 0 {return res}
	res += "."
	dic := make(map[int]int)
	for {
		if val, ok := dic[current]; ok {
			return fmt.Sprintf("%s(%s)", res[:val], res[val:])
		}
		dic[current] = len(res)
		current *= 10
		res += strconv.Itoa(current / denominator )
		current %= denominator;
		if current == 0 {return res}
	}
	return ""
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

### 129. Sum Root to Leaf Numbers

- level: 2-medium

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func sumNumbers(root *TreeNode) int {
    dfs(root, 0)
    return res
}

func dfs(root *TreeNode, path int) {
    if root == nil {
        return
    }
    var current = path*10 + root.Val
    if root.Left == nil && root.Right == nil {
        res += current
        return
    }
    if root.Left != nil {dfs(root.Left, current)}
    if root.Right != nil {dfs(root.Right, current)}
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

### 120. Triangle

- level: 2-medium

```go
func minimumTotal(triangle [][]int) int {
    var n = len(triangle)
    for i:=1; i<n; i++ {
        for j := range triangle[i] {
            var current = 3243242343243
            if j-1>=0 && triangle[i-1][j-1] < current {current = triangle[i-1][j-1]}
            if j < i && triangle[i-1][j] < current {current = triangle[i-1][j]}
            triangle[i][j] += current
        }
    }
    var res = 2423423432
    for _, x := range triangle[n-1] {
        if x < res {res = x}
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

### 108. Convert Sorted Array to Binary Search Tree

- level: 1-easy

```go
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

### 70. Climbing Stairs

- level: 1-easy

```go
func climbStairs(n int) int {
    if n == 1 {return 1}
    prevprev, prev, current := 1, 2, 2
    for i:=2; i<n; i++ {
        current = prev + prevprev
        prevprev = prev
        prev = current
    }
    return current
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

### 48. Rotate Image

- level: 2-medium

```go
func rotate(matrix [][]int)  {
    var n = len(matrix)
    for i:=0; i<n/2; i++ {
        for j:=i; j<n-i-1; j++ {
            matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = 
                matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
        }
    }
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

### 35. Search Insert Position

- level: 1-easy

```go
func searchInsert(nums []int, target int) int {
    return sort.SearchInts(nums, target)
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


