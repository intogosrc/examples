package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func testmd5() {
	md5.New()
	r := md5.Sum([]byte("hello"))
	result := hex.EncodeToString(r[:])
	fmt.Println(result)
}