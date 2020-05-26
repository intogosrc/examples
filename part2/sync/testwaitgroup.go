package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func testwaitgroup() {
	var wg sync.WaitGroup
	urlfmt := "http://localhost/index.html?index=%d"

	for i:=0; i<3; i++ {

		wg.Add(1)

		go func(index int){
			defer wg.Done()

			url := fmt.Sprintf(urlfmt, index)
			resp,err := http.Get(url)

			if err!=nil {
				fmt.Println("Get error ", err.Error())
				return
			}

			defer resp.Body.Close()

			result,err :=ioutil.ReadAll(resp.Body)
			if err!=nil {
				fmt.Println("Get error ", err.Error())
				return
			}

			fmt.Println("Get result ", string(result))
		}(i)
	}

	wg.Wait()
}