package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	if len(os.Args) < 2 {
		fmt.Println("No args received")
		os.Exit(1)
	}

	message, err := greetings.Hello(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
