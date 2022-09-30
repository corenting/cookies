package base64

import (
	"testing"
)

func TestDecodeBase64StringWithPadding(t *testing.T) {
	testString := "aGVsbG8gd29ybGQ="

	decodedData, err := DecodeBase64([]byte(testString))
	expected := "hello world"

	if err != nil {
		t.Errorf("DecodeBase64String error: %s", err)
	}
	if string(decodedData) != expected {
		t.Errorf("DecodeBase64String incorrect, got: %s, want: %s.", string(decodedData), expected)
	}
}

func TestDecodeBase64StringWithoutPadding(t *testing.T) {
	testString := "aGVsbG8gLi4u"

	decodedData, err := DecodeBase64([]byte(testString))
	expected := "hello ..."

	if err != nil {
		t.Errorf("DecodeBase64String error: %s", err)
	}
	if string(decodedData) != expected {
		t.Errorf("DecodeBase64String incorrect, got: %s, want: %s.", string(decodedData), expected)
	}
}
