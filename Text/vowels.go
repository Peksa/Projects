package main

import (
	"fmt"
	"os"
	"bufio"
)

// TODO return counts of each

var vowels = [...]rune { 'a', 'e', 'i', 'o', 'u', 'y' }

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}
		fmt.Println(countVowels(line));
	}
}

func countVowels(source string) int {
	ret := 0
	for _, c := range source {
		for _, vowel := range vowels {
			if c == vowel {
				ret++
				break;
			}
		}
	}
	return ret
}
