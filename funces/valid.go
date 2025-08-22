package funces

func Isvalid(str string) bool {
	if len(str) > 24 {
		return false
	}
	if str == "" {
		return false
	}
	for _, char := range str {
		if char < 32 || char > 126 {
			return false
		}
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			return true
		}
	}
	return false
}

func Isvalidname(name string) bool {
	mux.Lock()
	for _, v := range names {
		if v == name {
			mux.Unlock()
			return false
		}
	}
	mux.Unlock()
	return true
}

func Isvalidmassage(msg string) bool {
	for _, v := range msg {
		if v < 32 || v > 126 {
			return false
		}
	}
	return true
}
