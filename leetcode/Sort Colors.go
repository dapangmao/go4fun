func sortColors(nums []int)  {
    zero, one:= 0, 0
    for _, x := range nums {
        if x == 0 {
            zero++
        } else if x == 1 {
            one++
        }
    }
    for i:=0 ; i<len(nums); i++ {
        if zero > 0 {
            nums[i] = 0
            zero--
        } else if one > 0 {
            nums[i] = 1
            one--
        } else {
            nums[i] = 2
        }
    }
    
}
