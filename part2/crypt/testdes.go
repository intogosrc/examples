package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

var (
	ede2Key = []byte("68f59249070c231f") // 16*8 = 128b
)

// 对称加密
func testdes() {
	plaintext := []byte("exampleplaintext")

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)

	block, err := des.NewTripleDESCipher(tripleDESKey) // 对称加密方式，也可以用 AES 等方式
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	} // 这里是把加密向量保存在了加密串的头部

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure. 官方在这里建议对加密数据进行签名

	fmt.Println(hex.EncodeToString(ciphertext))
}

// 解密
func desdecrypt() {
	cryptedStr := "6b44317fdd7901585cd6d867aa691c24cb6d9ecb254ecd2f" // testdes 加密后的16进制串
	ciphertext,err := hex.DecodeString(cryptedStr)

	if err!=nil {
		panic(err)
	}

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)

	block, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < des.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:des.BlockSize] // 先取出加密向量
	ciphertext = ciphertext[des.BlockSize:] // 加密向量后面的都是加密正文

	// CBC mode always works in whole blocks.
	// CBC 模式，分组加密模式，该模式加密后的长度是块大小的倍数，不够长会用 padding 算法补齐
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	fmt.Printf("%s\n", ciphertext)
}
