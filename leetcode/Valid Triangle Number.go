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
