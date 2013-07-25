package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode/utf8"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\r\n")
		fmt.Println(reverseString(line));
	}
}

func reverseString(source string) string {
	ret := make([]rune, len(source))
	i := len(source)

	for _, c := range source {
		if c != utf8.RuneError {
			i--
			ret[i] = c
		}
	}

	return string(ret[i:])
}
