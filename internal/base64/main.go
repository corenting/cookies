package base64

import (
	"encoding/base64"
)

// Decode a given base64-encoded string
func DecodeBase64String(encodedString string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
