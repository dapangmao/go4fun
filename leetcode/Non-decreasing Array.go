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
