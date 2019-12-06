func validIPAddress(IP string) string {
	data := strings.Split(IP, ".")
	if len(data) == 4 {
		for _, d := range data {
			if !isIPv4(d) {return "Neither"}
		}
		return "IPv4"
	}
	data = strings.Split(IP, ":")
	if len(data) == 7 {
		for _, d := range data {
			if !isIPv6(d) {return "Neither"}
		}
		return "IPv6"
	}
	return "Neither"
}


func isIPv4(text string) bool {
	current, err := strconv.Atoi(text)
	if err != nil || current < 0 || current > 255 || strconv.Itoa(current) != text {
		return false
	}
	return true
}


func isIPv6(text string) bool {
	if len(text) != 4 || text[0] == '-' {return false}
	current, err := strconv.ParseInt("0x" + text, 0, 64)
	if err != nil || current < 0 {return false}
	return true
}
