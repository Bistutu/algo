package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("foo", "add"))
}

// 秘诀：映射
func isIsomorphic(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}
	sm := make(map[byte]byte, ls)
	tm := make(map[byte]byte, lt)

	for k := range s {
		ss, ts := s[k], t[k]
		if sm[ss] > 0 && sm[ss] != ts || tm[ts] > 0 && tm[ts] != ss {
			return false
		}
		sm[ss] = ts
		tm[ts] = ss
	}
	return true
}
