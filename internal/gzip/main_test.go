package gzip

import (
	"encoding/base64"
	"testing"
)

func TestUngzipText(t *testing.T) {
	// The text is stored as base64 to be included in the test easily
	compressedAndBase64Text := "H4sIAAAAAAAAAEtUKC4pysxLV8hPU0jOSCxKTC5JLSoGAOP+cfkWAAAA"
	data, err := base64.StdEncoding.DecodeString(compressedAndBase64Text)
	if err != nil {
		t.Errorf("Wrong test data")
	}

	decodedText := UngzipText(string(data[:]))
	expected := "a string of characters"
	if decodedText != expected {
		t.Errorf("Ungzip incorrect, got: %s, want: %s.", decodedText, expected)
	}
}
