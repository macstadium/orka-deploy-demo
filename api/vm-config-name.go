package api

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func GenerateVmConfigName() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		log.Fatalln(err)
	}

	return hex.EncodeToString(b)
}
