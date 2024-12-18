package bf

func closeBracket(i string, ip int) (int, error) {
	match := 1
	for match > 0 {
		ip--
		if ip < 0 {
			return ip, OutOfRangeError{
				IP: ip,
				Msg: "instruction pointer out of range to the left.\nAre you missing an open bracket?",
			}
		}
		if i[ip] == 93 {
			match++
		}
		if i[ip] == 91 {
			match--
		}
	}
	return ip, nil
}
