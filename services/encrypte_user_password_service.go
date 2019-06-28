package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

//EncryptePassword ...
func EncryptePassword(password string) string {

	secret := "msk*2019*?*xanik-gignox-?18"

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(password))

	// Get result and encode as hexadecimal strings
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
