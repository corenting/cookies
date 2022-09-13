package base64

import (
	"testing"
)

func TestDecodeBase64StringWithPadding(t *testing.T) {
	testString := "aGVsbG8gd29ybGQ="

	decodedString, err := DecodeBase64String(testString)
	expected := "hello world"
	if err != nil || decodedString != expected {
		t.Errorf("DecodeBase64String incorrect, got: %s, want: %s.", decodedString, expected)
	}
}

func TestDecodeBase64StringWithoutPadding(t *testing.T) {
	testString := "aGVsbG8gLi4u"

	decodedString, err := DecodeBase64String(testString)
	expected := "hello ..."
	if err != nil || decodedString != expected {
		t.Errorf("DecodeBase64String incorrect, got: %s, want: %s.", decodedString, expected)
	}
}
