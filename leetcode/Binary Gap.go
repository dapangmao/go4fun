func binaryGap(N int) int {
    current, res := 0, 0
    hasFirstOne, hasTwoOnes := false, false
    for N > 0 {
        d := N % 2
        if d == 0 {
            current++
        } else {
            if current > res && hasFirstOne {res = current}
            current = 0
            if hasFirstOne {hasTwoOnes = true}
            hasFirstOne = true
        }
        N /= 2
    }
    if !hasTwoOnes {return 0}
    return res + 1
}
