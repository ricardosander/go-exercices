package echo

import (
	"fmt"
	"os"
	"strings"
)

func Echo(params []string) {

	if len(params) < 1 {
		fmt.Println("No arguments given to echo.\n Exiting program.")
		os.Exit(0)
	}

	switch params[0] {
	case "1":
		echo1(params[1:])
	case "2":
		echo2(params[1:])
	case "3":
		echo3(params[1:])
	default:
		echo1(params)
	}
}

func echo1(params []string) {
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
