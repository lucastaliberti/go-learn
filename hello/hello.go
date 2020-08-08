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
}
