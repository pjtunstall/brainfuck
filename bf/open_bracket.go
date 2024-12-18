package bf

func openBracket(i string, ip int) (int, error) {
	match := 1
	for match > 0 {
		ip++
		if ip >= len(i) {
			return ip, OutOfRangeError{
				IP: ip,
				Msg: "instruction pointer out of range to the right.\nAre you missing a close bracket?",
			}
		}
		if i[ip] == 91 {
			match++
		}
		if i[ip] == 93 {
			match--
		}
	}
	return ip, nil
}