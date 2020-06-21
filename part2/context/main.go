package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	err := execWithContext(ctx)

	if err!=nil {
		fmt.Println("err is ", err.Error())
	}else{
		fmt.Println("no err")
	}

	time.Sleep(5*time.Second)
}

