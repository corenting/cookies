package decryption

import (
	"errors"
	"strings"

	"github.com/corenting/cookies/internal/base64"
)

// Decrypt and return the given cookie.
func DecryptCookie(initialCookie string) (string, error) {
	if len(strings.TrimSpace(initialCookie)) == 0 {
		return "", errors.New("Empty cookie")
	}
	result, decodeError := base64.DecodeBase64([]byte(initialCookie))
	if decodeError != nil {
		return "", decodeError
	}

	return string(result), nil
}
