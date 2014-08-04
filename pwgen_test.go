package pwgen

import "testing"

func TestNew(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = New(i)
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
}
