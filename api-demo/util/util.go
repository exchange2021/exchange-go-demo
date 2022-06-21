package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
)

// Get millisecond timestamp
func UnixMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// Calculate OpenApi signature
func GetSign(timestampMillis string, method string, requestPath string, queryString string, body string, apiSecret string) string {
	preMessage := getPreHash(timestampMillis, method, requestPath, queryString, body)
	return hmacSign([]byte(preMessage), []byte(apiSecret))
}

// get the string to be signed
func getPreHash(timestampMillis string, method string, requestPath string, queryString string, body string) string {
	var build strings.Builder
	build.WriteString(timestampMillis)
	build.WriteString(strings.ToUpper(method))
	build.WriteString(requestPath)
	if queryString != "" {
		build.WriteString("?")
		build.WriteString(queryString)
	}
	if body != "" {
		build.WriteString(body)
	}
	return build.String()
}

// hmac encrypted with sha256 algorithm
func hmacSign(message []byte, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return hex.EncodeToString(mac.Sum(nil))
}
