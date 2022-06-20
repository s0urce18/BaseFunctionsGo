package main

import (
	"fmt"
	"strconv"
)

func main() {
	for n := 1; n <= 100; n++ {
		if len(strconv.Itoa(n)) == 1 {
			if n == 1 {
				fmt.Println(n)
			}
		} else if len(strconv.Itoa(n)) == 2 {
			if strconv.Itoa(n)[0] == '1' || strconv.Itoa(n)[1] == '1' {
				fmt.Println(n)
			}
		} else if len(strconv.Itoa(n)) == 3 {
			if strconv.Itoa(n)[0] == '1' || strconv.Itoa(n)[1] == '1' || strconv.Itoa(n)[2] == '1' {
				fmt.Println(n)
			}
		}
	}
}
