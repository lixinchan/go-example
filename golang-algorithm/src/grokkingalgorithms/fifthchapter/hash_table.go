package fifthchapter

func PhoneTable(contacts map[string]string, key string) string {
	if len(key) == 0 {
		return "nil"
	}

	if phone, ok := contacts[key]; ok {
		return phone
	}
	return "nil"
}
