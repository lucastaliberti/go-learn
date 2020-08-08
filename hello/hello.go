package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Lucas"
	age := time.Now().Year() - 1990
	version := 1.1

	fmt.Printf("Hello World, %s, %d\n", name, age)
	fmt.Println("This is version", version)

	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")

	var command int
	// fmt.Scanf("%d", &command) // Explicit defines the input type
	fmt.Scan(&command)

	fmt.Println("Chosen command was", command)
}
