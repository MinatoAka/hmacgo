package hmacgo

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// SignSHA256 menghasilkan signature HMAC-SHA256
func SignSHA256(payload, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

// SignSHA512 menghasilkan signature HMAC-SHA512
func SignSHA512(payload, secret string) string {
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}
