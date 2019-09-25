func nextGreatestLetter(letters []byte, target byte) byte {
    var res byte = letters[0]
    for _, b := range letters {
        if b > target && (res <= target || b - target < res - target) {
                res = b
            } 
    }
    return res
}
