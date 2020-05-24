package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func testhmac() {
	h := hmac.New(md5.New, []byte("thisisvectorofke"))
	//h := hmac.New(sha256.New, []byte("thisisvectorofke")) // 计算散列值也可以用其他算法
	r := h.Sum([]byte("hello jack"))
	result := hex.EncodeToString(r)
	fmt.Println(result)
}
