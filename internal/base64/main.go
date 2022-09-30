package base64

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"errors"
	"io"
)

// Code from https://github.com/octopart/go-itsdangerous
// See licence: https://github.com/octopart/go-itsdangerous/blob/master/LICENSE
// Ignore some errors as cookies may not fit the specifications for zlib/base64
func DecodeBase64(encodedBytes []byte) ([]byte, error) {
	// if leading '.' itsdangerous has compressed this with zlib
	decompress := false
	if encodedBytes[0] == '.' {
		decompress = true
		encodedBytes = encodedBytes[1:]
	}

	for i := 0; i < len(encodedBytes)%4; i++ {
		encodedBytes = append(encodedBytes, '=')
	}

	resultBytes := make([]byte, base64.URLEncoding.DecodedLen(len(encodedBytes)))
	n, _ := base64.URLEncoding.Decode(resultBytes, encodedBytes)
	resultBytes = resultBytes[:n]

	if len(resultBytes) == 0 {
		return nil, errors.New("Invalid data for cookie")
	}

	// if leading '.' decompress now
	if decompress {
		bytesReader := bytes.NewReader(resultBytes)
		zReader, _ := zlib.NewReader(bytesReader)
		decompressedBytes, _ := io.ReadAll(zReader)
		return decompressedBytes, nil
	}

	return resultBytes, nil
}
