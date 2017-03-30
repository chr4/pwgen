package pwgen

import (
	"crypto/rand"
)

// Characters the password can contain
var num = "0123456789"
var alphaNum = num + "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var alphaNumSymbols = alphaNum + "_-?!.,@#$%^&*()=[]{}<>"

// New generates a random string of the given length out of the characters in char
func New(length int, chars string) string {
	var bytes = make([]byte, length)
	var op = byte(len(chars))

	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = chars[b%op]
	}
	return string(bytes)
}

// Num generates a random string of the given length out of numeric characters
func Num(length int) string {
	return New(length, num)
}

// AlphaNum generates a random string of the given length out of alphanumeric characters
func AlphaNum(length int) string {
	return New(length, alphaNum)
}

// AlphaNumSymbols generates a random string of the given length out of alphanumeric characters and
// symbols
func AlphaNumSymbols(length int) string {
	return New(length, alphaNumSymbols)
}
