func numUniqueEmails(emails []string) int {
    var set = make(map[[2]string]bool)
    clean := func(s string) string {
        idx := strings.Index(s, "+")
        if idx == -1 {
            idx = len(s)
        }
        return strings.Replace(s[:idx], ".", "", -1)
    }
    for _, email := range emails {
        var key [2]string
        arr := strings.Split(email, "@")
        key[0], key[1] = clean(arr[0]), arr[1]
        set[key] = true
        
    }
    return len(set) 
}
