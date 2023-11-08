package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	var name string = "Alex"
	lastname := "Kuznetsov"

	city, hobby := "Vilnius", "Hiking"

	const STATUS = "married"

	fmt.Println(name, lastname)
	fmt.Println(city, hobby)
	fmt.Println(STATUS)
}