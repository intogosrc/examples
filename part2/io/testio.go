package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func testreadfile() {
	data,err := ioutil.ReadFile("./iotestfile")
	if err!=nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func testreaddir() {
	files,err := ioutil.ReadDir(".")
	if err!=nil {
		panic(err)
	}

	for k,v := range files {
		fmt.Println(k, v.Name())
	}

}

func testreadall() {
	resp,err := http.Get("http://localhost/index.html")
	if err!=nil {
		panic(err)
	}

	defer resp.Body.Close()

	data,err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		panic(err)
	}

	fmt.Println(string(data))
}