package user

import "unicode"

func ValidatePassword(password string) bool {
	length := false
	upper := false
	lower := false
	numeric := false

	if len(password) >= 8 {
		length = true
	}
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			numeric = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		}
	}

	return length && upper && lower && numeric
}
