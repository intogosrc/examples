package main

// go 官方按照标准提供 公钥加密，私钥解密 ；私钥签名，公钥验证

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// 公钥加密
func RSAEncrypt() string {
	pk_bytes,err := ioutil.ReadFile("./rsa_public_key.pem")

	if err!=nil {
		panic(err)
	}

	block,_ := pem.Decode(pk_bytes)

	if block == nil {
		panic("parse public key failed")
	}

	pub_inf,err := x509.ParsePKIXPublicKey(block.Bytes)

	if err!=nil {
		panic(err)
	}

	pub := pub_inf.(*rsa.PublicKey)

	en,err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte("hello world"))
	if err!=nil {
		panic(err)
	}

	return hex.EncodeToString(en)
}

// 私钥解密
func RSADecrypt(en_str string) string {

	en_byte,err := hex.DecodeString(en_str)

	if err!=nil {
		panic(err)
	}

	pk_bytes,err := ioutil.ReadFile("./rsa_private_key.pem")

	if err!=nil {
		panic(err)
	}

	block,_ := pem.Decode(pk_bytes)

	if block == nil {
		panic("parse private key failed")
	}

	prik,err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err!=nil {
		panic(err)
	}

	de,err := rsa.DecryptPKCS1v15(rand.Reader, prik, en_byte)
	if err!=nil {
		panic(err)
	}

	return string(de)
}

// 私钥签名
func RSASign(str string) string {

	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng := rand.Reader

	message := []byte(str)

	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).
	hashed := sha256.Sum256(message)

	pk_bytes,err := ioutil.ReadFile("./rsa_private_key.pem")

	if err!=nil {
		panic(err)
	}

	block,_ := pem.Decode(pk_bytes)

	if block == nil {
		panic("parse private key failed")
	}

	prik,err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err!=nil {
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rng, prik, crypto.SHA256, hashed[:])

	if err!=nil {
		panic(err)
	}

	return hex.EncodeToString(signature)
}

//公钥验签
func RSAVerify(msg_str,sign_str string) bool {
	sign,err := hex.DecodeString(sign_str)

	if err!=nil {
		panic(err)
	}

	message := []byte(msg_str)

	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).
	hashed := sha256.Sum256(message)

	pk_bytes,err := ioutil.ReadFile("./rsa_public_key.pem")

	if err!=nil {
		panic(err)
	}

	block,_ := pem.Decode(pk_bytes)

	if block == nil {
		panic("parse public key failed")
	}

	pub_inf,err := x509.ParsePKIXPublicKey(block.Bytes)

	if err!=nil {
		panic(err)
	}

	pub := pub_inf.(*rsa.PublicKey)

	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], sign)

	if err!=nil {
		fmt.Println("verify failed ", err.Error())
		return false
	}

	return true
}