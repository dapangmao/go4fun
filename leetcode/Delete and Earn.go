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
