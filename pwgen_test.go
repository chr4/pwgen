package pwgen

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = New(i, "ab")
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
	if !verifyGeneratedPassword(password, "ab") {
		t.Errorf("Expected string to contain only {%s}, but got %s\n", alpha, password)
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
	if !verifyGeneratedPassword(password, num) {
		t.Errorf("Expected string to contain only {%s}, but got %s\n", num, password)
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
	if !verifyGeneratedPassword(password, alpha) {
		t.Errorf("Expected string to contain only {%s}, but got %s\n", alpha, password)
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
	if !verifyGeneratedPassword(password, symbols) {
		t.Errorf("Expected string to contain only {%s}, but got %s\n", symbols, password)
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
	if !verifyGeneratedPassword(password, alphaNum) {
		t.Errorf("Expected string to contain only {%s}, but got %s\n", alphaNum, password)
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
	if !verifyGeneratedPassword(password, alphaNumSymbols) {
		t.Errorf("Expected string to contain only {%s}, but got %s\n", alphaNumSymbols, password)
	}
}

func verifyGeneratedPassword(generatedString string, format string) bool {
	for _, char := range generatedString {
		if !strings.Contains(format, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}
