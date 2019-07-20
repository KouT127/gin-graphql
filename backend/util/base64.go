package util

import (
	"encoding/base64"
)

func Base64Encode(s string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(s))
	return encoded
}

func Base64Decode(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return s, err
	}
	return string(decoded), err
}
