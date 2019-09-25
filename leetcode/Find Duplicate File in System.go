func findDuplicate(paths []string) [][]string {
    dic := make(map[string][]string)
    for _, path := range paths {
        var current = strings.Split(path, " ")
        var prefix = current[0]
        for _, file := range current[1:] {
            var i = strings.Index(file, "(")
            var value = prefix + "/" + file[:i]
            var key = file[i+1:len(file)-1]
            dic[key] = append(dic[key], value)  
        }
    }
    var res [][]string
    for _, v := range dic {
        if len(v) < 2 {continue}
        res = append(res, v)
    }
    return res
}
