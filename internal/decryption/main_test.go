package decryption

import (
	"fmt"
	"testing"
)

func TestDecryptCookieWithFlaskCookie(t *testing.T) {
	testCases := []struct {
		cookieContent   string
		expectedContent string
	}{
		{
			"eyJhbmFseXRpY3NfdGFncyI6W10sIl9mcmVzaCI6ZmFsc2V9.YyCNdg.4HOiVlVCi4f4imDshhUyhkyrSSk",
			"{\"analytics_tags\":[],\"_fresh\":false}",
		},
		{
			"eyJmb28iOiJiYXIifQ===",
			"{\"foo\":\"bar\"}",
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			decoded, err := DecryptCookie(testCase.cookieContent)
			if err != nil {
				t.Errorf("Error on DecryptCookie: %v", err)
			}
			if decoded != testCase.expectedContent {
				t.Errorf("DecryptCookie incorrect, got: %s, want: %s.", decoded, testCase.expectedContent)
			}
		})
	}
}

func TestDecryptCookieInvalidInput(t *testing.T) {
	testCases := []struct {
		cookieContent string
	}{
		{
			"",
		},
		{
			".",
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			decoded, err := DecryptCookie(testCase.cookieContent)
			if err == nil || decoded != "" {
				t.Errorf("DecryptCookie should raise error")
			}
		})
	}
}
