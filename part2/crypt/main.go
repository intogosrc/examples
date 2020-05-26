package main

import "fmt"

func main() {
	//testdes()
	//desdecrypt()

	//enstr := RSAEncrypt()
	//fmt.Println(enstr)
	//
	//destr := RSADecrypt(enstr)
	//fmt.Println(destr)

	s := "hello world"
	sign := RSASign(s)
	isok := RSAVerify(s, sign)

	fmt.Println("verify result is ", isok)


}


