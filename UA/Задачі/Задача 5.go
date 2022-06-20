package main

func main() {
	s := "abcdef"
	u := ""
	for _, c := range s {
		u = u + string(c+2)
	}
}
