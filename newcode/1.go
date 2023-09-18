package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	m := make(map[int]struct{})
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	t := s.Text()
	strs := strings.Split(t, ",")
	for i := 0; i < len(strs); i++ {
		a, _ := strconv.Atoi(strs[i])
		m[a] = struct{}{}
	}
	fmt.Println(len(m))
}
