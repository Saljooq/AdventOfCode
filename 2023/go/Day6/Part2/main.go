package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type ranges struct {
	startDestination int
	startSource      int
	sharedRange      int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// we scan the first line
	scanner.Scan()
	firstLine := scanner.Text()
	times := toIntArr(strings.Split(firstLine, ":")[1])
	scanner.Scan() // we move one next line
	secondLine := scanner.Text()
	distances := toIntArr(strings.Split(secondLine, ":")[1])

	marginOfError := 1

	for i, time := range times {
		dist := distances[i]
		currMargin := 0

		for j := 0; j <= time; j++ {
			speed := j
			timeLeft := time - j
			distanceCovered := speed * timeLeft
			if distanceCovered > dist {
				currMargin += 1
			}

		}
		log.Println("Curr Margin ", i, " : ", currMargin)
		marginOfError *= currMargin
	}

	log.Println("Margin of error: ", marginOfError)
}

func toIntArr(s string) []int {
	strNum := strings.Replace(s, " ", "", -1)
	n, _ := strconv.Atoi(strNum)

	return []int{n}
}
