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
