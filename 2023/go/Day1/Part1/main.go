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

	for {
		curr = 0
		_, err := fmt.Scanln(&input)
		if err != nil && errors.Is(err, io.EOF) {
			break
		}

		for _, a := range input {
			if a <= '9' && a >= '0' { // is alpha numeral
				curr = int(a - '0')
				curr *= 10
				break
			}
		}

		for i := range input {
			currAlph := input[len(input)-i-1]
			if currAlph <= '9' && currAlph >= '0' { // is alpha numeral
				curr += int(currAlph - '0')
				break
			}
		}
		fmt.Println(curr)

		total += curr
	}

	fmt.Println("Total: ", total)
}
