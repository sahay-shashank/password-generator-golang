package password

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var ()

func Password(length int) string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%"
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	var result strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := r.Intn(len(charSet))
		char := charSet[randomIndex]
		result.WriteByte(char)
		if unicode.IsUpper(rune(char)) {
			hasUpper = true
		} else if unicode.IsLower(rune(char)) {
			hasLower = true
		} else if unicode.IsDigit(rune(char)) {
			hasDigit = true
		} else {
			hasSpecial = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		return Password(length)
	}
	return result.String()
}
