package main

import "fmt"

//Global variable declaration
var (
	m int
	n int
)

func main() {
	var x int = 1
	var y int
	fmt.Println(x)
	fmt.Println(y)

	var a, b, c = 5.25, 25.25, 14.15
	fmt.Println(a, b, c)

	city := "Berlin"
	Country := "Germany"
	fmt.Println(city)
	fmt.Println(Country)

	food, drink, price := "Pizza", "Pepsi", 125
	fmt.Println(food, drink, price)
	m, n = 1, 2
	fmt.Println(m, n)
}

