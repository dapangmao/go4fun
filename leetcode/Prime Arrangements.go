var primes = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

func numPrimeArrangements(n int) int {
    count := 0
    for _, p := range primes {
        if p <= n {count++}
    }
    return factorial(n - count) * factorial(count) 
}

func factorial(n int) int {
    res := 1
    for i:=n; i>=2; i-- {
        res *= i
    }
    return res
} 
