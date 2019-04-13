package helper

import (
	"crypto/rand"
	"io"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateVerificationCode() (string,error) {
    b := make([]byte, 6)
    n, err := io.ReadAtLeast(rand.Reader, b, 6)
    if n != 6 {
        return "",err
    }
    for i := 0; i < len(b); i++ {
        b[i] = table[int(b[i])%len(table)]
    }
    return string(b),nil
}

