package validators

import (
	"unicode"

	"github.com/gopher-fleece/runtime"
)

// Custom validation function to check if a string starts with a letter
func ValidateStartsWithLetter(fl runtime.ValidationFieldLevel) bool {
	field := fl.Field().String()
	if field == "" {
		return false
	}
	firstChar := rune(field[0])
	return unicode.IsLetter(firstChar)
}
