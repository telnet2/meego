package strings

func Version() string {
	return "v1"
}

func Reverse(s string) string {
	bts := make([]byte, len(s))
	for i := len(s); i > 0; i-- {
		bts[i-1] = s[len(s)-i]
	}
	return string(bts)
}
