package strings

func OR(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	}
	return str
}
