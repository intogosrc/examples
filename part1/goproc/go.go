package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func gotest1() {
	users := []string{
		"Leo",
		"Jack",
		"Tommy",
	}

	results := new(sync.Map) // 系统 map 类型为多go程不安全的 NTS，sync.Map 多go程安全

	url := "http://localhost/demos/test.php?name=%s"

	var wg sync.WaitGroup

	for _, user := range users {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			reqUrl := fmt.Sprintf(url, u)
			resp, err := http.Get(reqUrl)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			result, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			results.Store(u, string(result))

		}(user)
	}

	wg.Wait()

	results.Range(func(k, v interface{}) bool {
		key := k.(string)
		value := v.(string)

		fmt.Println(key + " = " + value)

		return true
	})

}
