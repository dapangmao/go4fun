func canJump(nums []int) bool {
    prev := 1
    for _, x := range nums[:len(nums)-1] {
        prev -= 1
        if x == 0 && prev <= 0 {return false}
        if x > prev {prev = x}
    } 
    return true
}
