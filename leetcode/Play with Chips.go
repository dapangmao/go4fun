func minCostToMoveChips(chips []int) int {
    var odd int
    for _, c := range chips {
        if c % 2 == 1 {odd++}
    }
    even := len(chips) - odd
    if even < odd {return even}
    return odd
}
