package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GetHmac(key string, value string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(value))
	return hex.EncodeToString(mac.Sum(nil))
}

func CheckHmac(key string, value string, hash string) bool {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(value))
	strHash := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(strHash), []byte(hash))
}
