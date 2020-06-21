package main

import (
	"context"
	"fmt"
	"time"
)

func contextTest() (err error) {

	a := 0
	defer func(){
		fmt.Println("a is ", a)
	}()

	a = 1

	time.Sleep(3*time.Second)

	a = 2

	return err
}

func execWithContext(ctx context.Context) error {
	doneCh := make(chan error, 1)

	go func(){
		doneCh <- contextTest()
	}()

	select {
	case err := <-doneCh :
		return err
	case <-ctx.Done():
		<-doneCh
		return ctx.Err()
	}

	//time.Sleep(4*time.Second)
	//return nil
}





