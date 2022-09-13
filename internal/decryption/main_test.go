package decryption

import (
	"testing"
)

func TestDecryptCookieWithFlaskCookie(t *testing.T) {
	encoded := "eyJhbmFseXRpY3NfdGFncyI6W10sIl9mcmVzaCI6ZmFsc2V9.YyCNdg.4HOiVlVCi4f4imDshhUyhkyrSSk"

	decoded, err := DecryptCookie(encoded)
	expected := "{\"analytics_tags\":[],\"_fresh\":false}"
	if err != nil || decoded != expected {
		t.Errorf("DecodeBase64String incorrect, got: %s, want: %s.", decoded, expected)
	}
}
