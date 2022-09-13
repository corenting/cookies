package gzip

import (
	"compress/gzip"
	"io"
	"log"
	"strings"
)

// Check if the cookie looks compressed with gzip
func IsGzippedCookie(cookie string) bool {
	return cookie[0] == '.'
}

// Ungzip the provided text
func UngzipText(initialText string) string {
	reader := strings.NewReader(initialText)
	gzipReader, gzipErr := gzip.NewReader(reader)
	if gzipErr != nil {
		log.Fatal("Error when ungzipping text")
	}

	decodedText, decodeError := io.ReadAll(gzipReader)
	if decodeError != nil {
		log.Fatal("Error when ungzipping text")
	}

	return string(decodedText)
}
