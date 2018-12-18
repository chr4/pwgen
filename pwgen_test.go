package pwgen

import "testing"

func TestNew(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = New(i, "ab")
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
}

func TestNum(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = Num(i)
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
}

func TestAlpha(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = Alpha(i)
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
}

func TestSymbols(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = Symbols(i)
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
}

func TestAlphaNum(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = AlphaNum(i)
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
}

func TestAlphaNumSymbols(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = AlphaNumSymbols(i)
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
}
