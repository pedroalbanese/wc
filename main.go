package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please pass at least one file name")
		os.Exit(1)
	}

	totalLines, totalWords, totalBytes := 0, 0, 0

	for _, fileName := range os.Args[1:] {
		fileData, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error reading %s: %s\n", fileName, err)
			continue
		}

		content := string(fileData)
		lines := strings.Split(content, "\n")
		words := 0
		bytes := len(fileData)

		if len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
			lines = lines[:len(lines)-1]
		}

		for _, line := range lines {
			words += len(strings.Fields(line))
		}

		fmt.Printf("%10d%10d%10d %s\n", len(lines), words, bytes, fileName)

		totalLines += len(lines)
		totalWords += words
		totalBytes += bytes
	}

	if len(os.Args) > 2 {
		fmt.Printf("%10d%10d%10d total\n", totalLines, totalWords, totalBytes)
	}
}
