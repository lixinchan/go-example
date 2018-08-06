package fifthchapter

import "testing"

func TestPhoneTable(t *testing.T) {
	contacts := make(map[string]string)
	contacts["clx"] = "188"
	contacts["xiaoxin"] = "186"
	t.Log(PhoneTable(contacts, "clx"))
}
