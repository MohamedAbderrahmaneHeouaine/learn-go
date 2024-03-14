package main

import (
	"fmt"
	"time"
)

func goroutine() {
	result := make(chan string) // unbuffered chanel
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch)
	for msg := range msgch {
		fmt.Println(msg)
	}

	//go func() {
	//	result := fetchResource(1)
	//	fmt.Println(result)
	//}()
	go func() {
		res := <-result
		fmt.Println(res)
	}()
	result <- "foo" // is now full -> it will block -> block here

	//this code below will never execute!!!
	//res := <-result
	//fmt.Println(res)
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("some result %d", n)
}
