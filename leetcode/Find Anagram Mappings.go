func anagramMappings(A []int, B []int) []int {
    var m = make(map[int]int)
    for i, x := range B {
        m[x] = i
    }
    for i, x := range A {
        B[i] = m[x]
    }
    return B
}
