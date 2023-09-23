package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	time.Sleep(1 * time.Second)
	fmt.Println("gogogo")
}

func test() {
	fmt.Println("panic")
	panic("")
}
