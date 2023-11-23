package main

import (
	"fmt"
	dup "go-exercices/dub"
	"go-exercices/echo"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("No command selected. Exiting program.")
		os.Exit(0)
	}

	command := args[0]
	switch command {
	case "echo":
		echo.Echo(args[1:])
	case "dub":
		dup.Dup(args[1:])
	default:
		fmt.Printf("Invalid command: %s\n", command)
	}
}
