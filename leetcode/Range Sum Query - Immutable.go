type NumArray struct {
    left, right []int
}


func Constructor(nums []int) NumArray {
    n := len(nums)
    l, r := make([]int, n), make([]int, n)
    for i:=0; i<n; i++ {
        if i == 0 {
            l[i] = nums[i]
        } else {
            l[i] = nums[i] + l[i-1]
        }
    }
    for i:=n-1; i>=0; i-- {
        if i == n-1 {
            r[i] = nums[i]
        } else {
            r[i] += nums[i] + r[i+1]
        }
    }
    return NumArray{l, r}
}


func (this *NumArray) SumRange(i int, j int) int {
    var sum = this.left[len(this.left)-1]
    l, r := 0, 0
    if i-1 >=0 {l = this.left[i-1]}
    if j+1 < len(this.left) {r = this.right[j+1]}
    return sum - l - r
}
