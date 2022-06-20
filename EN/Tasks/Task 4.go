package main

func main() {
	s := "123"
	u := ""
	for _, c := range s {
		u = string(c) + u
	}
}
