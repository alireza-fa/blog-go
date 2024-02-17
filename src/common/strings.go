package common

import "strings"

var (
	lowerCharSet string = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberSet    string = "1234567890"
)

func CheckPassword(password string) bool {
	if len(password) > 8 || len(password) > 64 {
		return false
	}

	if !strings.ContainsAny(password, lowerCharSet) {
		return false
	}

	if !strings.ContainsAny(password, upperCharSet) {
		return false
	}

	if !strings.ContainsAny(password, numberSet) {
		return false
	}

	return true
}
