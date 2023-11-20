package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	NumberOfBytes   = "c"
	NumberOfLines   = "l"
	NumberOfWords   = "w"
	NumberOfChars   = "m"
	DefaultFilename = "test.txt"
)

func main() {
	cm := flag.String("ccwc", "", "command to run")
	flag.Parse()

	var filename string

	args := flag.Args()
	if len(args) == 0 {
		filename = DefaultFilename
		handleCommand(*cm, filename)
		return
	}

	filename = flag.Args()[0]
	handleCommand(*cm, filename)
}

func handleCommand(command, filename string) {
	content := readFile(filename)

	switch command {

	case NumberOfBytes:
		fmt.Println(len(content))

	case NumberOfLines:
		count := countFileLines(content)
		fmt.Println(count)

	case NumberOfWords:
		//n := regexp.MustCompile("\\s+").FindAllString(string(content), -1)
		fmt.Println(len(strings.Fields(string(content))))

	case NumberOfChars:
		fmt.Println(utf8.RuneCount(content))

	default:
		fmt.Print(len(content), countFileLines(content), len(strings.Fields(string(content))))

	}
}

func readFile(filename string) []byte {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return content
}

func countFileLines(bytes []byte) int {
	counter := 0
	for _, v := range bytes {
		if v == '\n' {
			counter++
		}
	}
	return counter
}
