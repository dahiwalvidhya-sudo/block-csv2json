package encode

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBytes(hexStr string) ([]byte, error) {
	return hex.DecodeString(hexStr)
}

func BytesToHex(b []byte) string {
	return hex.EncodeToString(b)
}

func Base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
