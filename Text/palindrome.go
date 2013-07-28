package main

import (
	"fmt"
	"os"
	"bufio"
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
		fmt.Println(isPalindrome(line));
	}
}

func isPalindrome(source string) bool {
	runes := []rune(source)
	half := len(runes)/2
	for i := 0; i < half; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}
