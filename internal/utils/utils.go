package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomHexColor() string {
	bytes := make([]byte, 3)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
