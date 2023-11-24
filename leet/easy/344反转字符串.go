package main

import "fmt"

func main() {
	msg := []byte("hello")
	reverseString(msg)
	fmt.Println(string(msg))
}
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
