package service

import (
	"crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
)
func EncryptePassword(password_ string) string {
	
	secret := "msk*2019*os*xanik-gignox-?18"
	data := password_

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))
	
	// Write Data to it
	h.Write([]byte(data))
	
	// Get result and encode as hexadecimal strings
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
