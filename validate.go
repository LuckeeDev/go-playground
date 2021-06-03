package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func validateSignature(token string, data []byte, expectedSignature string) (result bool, err error) {
	secret := []byte(token)
		
	hash := hmac.New(sha256.New, secret)
	hash.Write(data)
	
	hex := "sha256=" + hex.EncodeToString(hash.Sum(nil))

	return hmac.Equal([]byte(hex), []byte(expectedSignature)), nil
}
