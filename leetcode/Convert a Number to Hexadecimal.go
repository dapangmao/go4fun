func toHex(num int) string {
    if num < 0 {
        return fmt.Sprintf("%x", math.MaxInt32 * 2 + 2 + num)     
    }
    return fmt.Sprintf("%x", num)
}
