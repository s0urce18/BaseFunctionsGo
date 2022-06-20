package main

type Human interface {
	FullName() string
	HowOld() uint
	FullLocation() string
}

type Person struct {
	FirstName  string
	SecondName string
	Age        uint
	Country    string
	City       string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.SecondName
}

func (p Person) HowOld() uint {
	return p.Age
}

func (p Person) FullLocation() string {
	return p.City + ", " + p.Country
}

func main() {

}
