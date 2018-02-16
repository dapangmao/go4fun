func complexNumberMultiply(a string, b string) string {
    partA, partB := strings.Split(a, "+"), strings.Split(b, "+")
    
    t := func(s string) int {
        res, _ := strconv.Atoi(s)
        return res
    }
    g := func(s string) int {
        res, _ := strconv.Atoi(s[:len(s)-1])
        return res
    }
  
    nonI := t(partA[0])*t(partB[0]) - g(partA[1])*g(partB[1])
    I := g(partA[1])*t(partB[0]) + t(partA[0])*g(partB[1])
    return fmt.Sprintf("%v+%vi", nonI, I)
}
