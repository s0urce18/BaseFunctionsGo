package main

import "fmt"

func main() {
	for n := 1; n <= 100; n++ {
		if n%2 == 0 {
			fmt.Println(n)
		}
	}
}
