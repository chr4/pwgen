package pwgen

import "testing"
// import "fmt"

func TestNew(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = New(i, "ab")
		// fmt.Println(password)
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
}

func TestNum(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = Num(i)
		// fmt.Println(password)
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
}

func TestAlphaNum(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = AlphaNum(i)
		// fmt.Println(password)
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
}

func TestAlphaNumSymbols(t *testing.T) {
	var password string
	for i := 1; i < 21; i++ {
		password = AlphaNumSymbols(i)
		// fmt.Println(password)
		if len(password) != i {
			t.Errorf("Expected length %d, got %d\n", i, len(password))
		}
	}
}
