package basictype

// has prefix
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// has suffix
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// contains
func Contains(s, substr string) bool {
	for idx := 0; idx < len(s); idx++ {
		if HasPrefix(s[idx:], substr) {
			return true
		}
	}
	return false
}
