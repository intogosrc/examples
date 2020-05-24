package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func shatest() {
	r := sha256.Sum256([]byte("hello jack"))
	result := hex.EncodeToString(r[:])
	fmt.Println(result)
}