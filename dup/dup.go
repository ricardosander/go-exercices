package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) > 0 {
		computeFiles(files, counts)
	} else {
		computeLines(counts)
	}

	fmt.Println("The following lines have dubs:")
	fmt.Println("dubs\tline")
	for line, num := range counts {
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, line)
		}
	}
}

func computeFiles(files []string, counts map[string]int) {
	for _, fileName := range files {
		computeFile2(fileName, counts)
	}
}

func computeFile(fileName string, counts map[string]int) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup: %v\n", err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		computeLine(counts, fileScanner.Text())
	}
}

func computeFile2(fileName string, counts map[string]int) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup: %v\n", err)
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		computeLine(counts, line)
	}
}

func computeLines(counts map[string]int) {
	inputScanner := bufio.NewScanner(os.Stdin)
	var input string

	fmt.Println("Give me input lines to be compute. Empty line to see the result:")
	for inputScanner.Scan() {
		input = inputScanner.Text()
		if input == "" {
			break
		}
		computeLine(counts, input)
	}
}

func computeLine(counts map[string]int, line string) {
	counts[line]++
}
