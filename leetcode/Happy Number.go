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
