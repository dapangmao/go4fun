func validUtf8(data []int) bool {
    var count = 0
    for _, c := range data {
        if count == 0 {
            var start = 1<<2 + 1<<1
            if c >> 5 == start {
                count = 1
            } else if c >> 4 == start + 1<<3 {
                count = 2
            } else if c >> 3 == start + 1<<3 + 1<<4 {
                count = 3  
            } else if c >> 7 != 0 {
                return false
            }
        } else {
            if c >> 6 != 2 {return false}
            count--
        }
    }
    return count == 0
}
