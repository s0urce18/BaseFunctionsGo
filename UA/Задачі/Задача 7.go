package main

func f(firstName string, secondName string, age string, country string, city string) map[string]string {
	return map[string]string{
		"firstName":  firstName,
		"secondName": secondName,
		"age":        age,
		"country":    country,
		"city":       city,
	}
}

func main() {
	people := []map[string]string{}
	people = append(people, f("Jack", "Brown", "17", "USA", "NewYork"))
}
