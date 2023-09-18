package main

import (
	"container/ring"
	"fmt"
)

func main() {
	l := ring.New(1)
	fmt.Println(l.Value)
}
