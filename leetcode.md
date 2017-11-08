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
        if set[word] < set[ans] || set[word] == set[ans] && word >= ans {
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


