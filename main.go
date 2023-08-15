package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	} else {
		fmt.Println("Please pass a file name")
		os.Exit(1)
	}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Err ", err)
	}
	scanner := bufio.NewScanner(file)
	lines, words, characters := 0, 0, 0
	for scanner.Scan() {
		lines++

		line := scanner.Text()
		characters += len(line)

		splitLines := strings.Split(line, " ")
		words += len(splitLines)
	}
	fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)
}
