package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {

	var input string

	var curr int = 0
	var total int = 0

	numDict := []string{"one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine"}

	for {
		curr = 0
		_, err := fmt.Scanln(&input)
		if err != nil && errors.Is(err, io.EOF) {
			break
		}

		for i := range input {
			shouldBreak, newCurr := isAlphaNumeral(i, input, true, numDict)
			if shouldBreak {
				curr = newCurr * 10
				break
			}
		}

		for i := range input {
			shouldBreak, newCurr := isAlphaNumeral(len(input)-1-i, input, false, numDict)
			if shouldBreak {
				curr += newCurr
				break
			}
		}
		fmt.Println(curr)

		total += curr
	}

	fmt.Println("Total: ", total)
}

func isAlphaNumeral(ind int, str string, posDir bool, numDict []string) (bool, int) {
	a := rune(str[ind])
	if a <= '9' && a >= '0' {
		curr := int(a - '0')
		return true, curr
	} else if posDir {
		for i, word := range numDict {
			if len(word)+ind < len(str) {
				if word == str[ind:ind+len(word)] {
					return true, i + 1
				}
			}
		}
	} else { //if negative dir
		for i, word := range numDict {
			if ind+1-len(word) >= 0 {
				if word == str[ind+1-len(word):ind+1] {
					return true, i + 1
				}
			}
		}
	}
	return false, -1
}
