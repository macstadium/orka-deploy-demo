package api

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func GenerateVmConfigName() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		log.Fatalln(err)
	}

	return fmt.Sprintf("sb-%s", hex.EncodeToString(b))
}
