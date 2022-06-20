package main

import "fmt"

func main() {
	for n := 1; n <= 100; n++ {
		if n%5 == 0 {
			fmt.Println(n)
		} else if n%7 == 0 {
			fmt.Println(n)
		}
	}
}
