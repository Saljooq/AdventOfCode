package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	pointLookup := map[int]int{}

	pointLookup[1] = 1
	for i := 2; i <= 11; i++ {
		pointLookup[i] = pointLookup[i-1] * 2
	}

	total := 0

	for scanner.Scan() {

		input := scanner.Text()
		mainSplit := strings.Split(input, ":")
		numberStr := strings.Split(mainSplit[1], "|")
		ourNumStrArr := strings.Split(strings.TrimSpace(numberStr[0]), " ")
		winningNumStr := strings.Split(strings.TrimSpace(numberStr[1]), " ")

		currWin := 0

		for _, winner := range winningNumStr {
			for _, our := range ourNumStrArr {
				if our == "" || winner == "" {
					continue
				}
				if our == winner {
					currWin += 1
				}
			}
		}

		total += pointLookup[currWin]

	}

	fmt.Println("Total: ", total)
}
