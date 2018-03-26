var codes = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

func uniqueMorseRepresentations(words []string) int {
    set := make(map[string]bool)
    for _, word := range words {
        var current string
        for _, b := range []byte(word) {
            current += codes[int(b-'a')]
        }
        set[current] = true
    }
    return len(set)
}
