package base64

import "encoding/base64"

// EncodeString encodes the given string into base64
func EncodeString(str string) string {
	encode := base64.StdEncoding.EncodeToString([]byte(str))
	return encode
}

// DecodeString decodes a base64 encoded string
func DecodeString(str string) string {
	decode, _ := base64.StdEncoding.DecodeString(str)
	return string(decode)
}
