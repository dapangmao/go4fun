func isPowerOfFour(num int) bool {
    var x = 1
    for x <= num {
        if x == num {return true}
        x *= 4
    }
    return false
}
