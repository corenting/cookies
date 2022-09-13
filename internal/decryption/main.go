package decryption

import (
	"strings"

	"github.com/corenting/cookies/internal/base64"
	"github.com/corenting/cookies/internal/gzip"
)

// Decrypt and return the given cookie.
func DecryptCookie(initialCookie string) (string, error) {
	compressed := false
	cookie := string(initialCookie)

	// If starts with . , should be a Flask compressed cookie
	if gzip.IsGzippedCookie(cookie) {
		compressed = true
		cookie = cookie[1:]
	}

	// Get only the first section
	cookie = strings.Split(cookie, ".")[0]

	// Decrypt the base 64 string
	cookie, decodeError := base64.DecodeBase64String(cookie)
	if decodeError != nil {
		return "", decodeError
	}

	if compressed {
		cookie = gzip.UngzipText(cookie)
	}

	return cookie, nil
}
