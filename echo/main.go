package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo3(os.Args[1:])
}

func echo(params []string) {
	var sep, s string
	for i := 0; i < len(params); i++ {
		s += sep + params[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(params []string) {
	for _, value := range params {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func echo3(params []string) {
	fmt.Println(strings.Join(params, " "))
}
