package utils

import (
	"regexp"
	"strings"
	"time"

	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
)

var (
	owaspSymbols = `!"#$%&'()*+,-./:;<=>?@[\]^_{|}~` + "`"
	loginRegex   = regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`)
)

func ValidatePassword(password string) error {
	runesPassword := []rune(password)
	if len(runesPassword) < 8 {
		return domain.ErrPasswordTooShort
	}
	if len(runesPassword) > 128 {
		return domain.ErrPasswordTooLong
	}

	if !strings.ContainsAny(password, owaspSymbols) {
		return domain.ErrMissingRequiredCharacters
	}

	return nil
}

func ValidateUsername(userName string) error {
	runeUserName := []rune(userName)
	if !loginRegex.MatchString(userName) {
		return domain.ErrInvalidCharacters
	}

	if len(runeUserName) < 3 {
		return domain.ErrLoginTooShort
	}
	if len(runeUserName) > 20 {
		return domain.ErrLoginTooLong
	}
	if !HasNoConsecutiveSpecial(userName) {
		return domain.ErrInvalidCharacters
	}

	return nil
}

func ValidateIncomingDate(date time.Time) error {
	minDate := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
	if date.After(minDate) && date.Before(time.Now()) {
		return nil
	}
	return domain.ErrInvalidDate

}

func HasNoConsecutiveSpecial(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if IsSpecial(s[i]) && IsSpecial(s[i+1]) {
			return false
		}
	}
	return true
}

func IsSpecial(c byte) bool {
	return c == '.' || c == '_' || c == '-'
}
