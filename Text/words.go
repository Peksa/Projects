package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

// TODO read from file and generate summary

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\r\n")
		fmt.Println(countWords(line));
	}
}

func countWords(source string) int {
	if source == "" {
		return 0
	}
	ret := 1
	for _, c := range source {
		if c == ' ' {
			ret++
		}
	}
	return ret
}
