func convert(s string, numRows int) string {
    if numRows == 0 {return ""}
    if numRows == 1 {return s}
    buffer := make([]string, numRows)
    down, j := true, 0
    for _, x := range s {
        buffer[j] += string(x)
        if down {
            j += 1
        } else {
            j -= 1
        }
        if j >= numRows {
            j = numRows - 2
            down = false
        } 
        if j < 0 {
            j = 1
            down = true
        }
    }
    return strings.Join(buffer, "")
}
