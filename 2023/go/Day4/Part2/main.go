package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	copies := map[int]int{}

	for i := 1; i <= 197; i++ {
		copies[i] = 1
	}

	total := 0

	for scanner.Scan() {

		input := scanner.Text()
		mainSplit := strings.Split(input, ":")

		cardStrArr := strings.Split(strings.TrimSpace(mainSplit[0]), " ")
		cardNum, _ := strconv.Atoi(cardStrArr[len(cardStrArr)-1])

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

		fmt.Println("Card num: ", cardNum, "  wins: ", currWin, "  copies: ", copies[cardNum])

		for i := cardNum + 1; i <= cardNum+currWin; i++ {
			copies[i] += copies[cardNum]
		}

	}

	for _, v := range copies {
		total += v
	}

	fmt.Println("Total: ", total)
}
