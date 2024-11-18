package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	secretKey := "ini_secret_key"
	timestamp := "2024-11-18T20:28:00+07:00"
	endpoint := "/wallet/v1/ex/1/unlink"
	strPayload := ``
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	strPayload = re.ReplaceAllString(strPayload, "")
	strPayload = strings.ToLower(strPayload) + timestamp + endpoint
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(strPayload))
	generatedSignature := hex.EncodeToString(h.Sum(nil))

	fmt.Println(generatedSignature)
}
