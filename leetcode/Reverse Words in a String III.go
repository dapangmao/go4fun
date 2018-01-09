func reverseWords(s string) string {
	var i int
  var bytes = []byte(s)   // not {}
	for j, x := range s {
		if x == ' ' {
			rw(bytes[:], i, j-1) // have to use slice
			i = j + 1
		}
	}
	rw(bytes[:], i, len(s)-1) // have to use slice
	return string(bytes)
}

func rw(b []byte, i, j int)  {
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}
