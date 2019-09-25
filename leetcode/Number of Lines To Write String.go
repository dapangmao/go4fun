func numberOfLines(widths []int, S string) []int {
    if len(S) == 0 {return []int{0, 0}}
    rowNo, currentRow := 1, 0
    for _, bt := range []byte(S) {
        var current = widths[int(bt - 'a')]
        if currentRow + current <= 100 {
            currentRow += current
        } else {
            rowNo++
            currentRow = current
        }
    }
    return []int{rowNo, currentRow}
}
