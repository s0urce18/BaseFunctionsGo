package main

import (
	"strings"
)

func f(str string, c string) string {
	s := strings.Split(str, c)
	for i := 1; i < len(s); i++ {
		s[i] = strings.ToUpper(string(s[i][0])) + s[i][1:]
	}
	return strings.Join(s, "")
}

func main() {
}
