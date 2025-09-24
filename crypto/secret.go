package crypto

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSecretKey() string {
	key := make([]byte, 48)
	rand.Read(key)
	return hex.EncodeToString(key)
}
