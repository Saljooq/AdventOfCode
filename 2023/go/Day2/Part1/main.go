package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	redMax, greenMax, blueMax := 12, 13, 14

	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	curr := 0
	var failed bool

	for {
		curr += 1

		var input string
		if scanner.Scan() {
			input = scanner.Text()
		} else {
			break
		}

		res := strings.Split(input, ":")
		// we can safely ignore the first part of the separate line
		data := strings.Split(res[1], ";")
		failed = false

		for _, draw := range data {
			drawArr := strings.Split(draw, ",")
			for _, numAndColor := range drawArr {
				clean := strings.TrimSpace(numAndColor)
				individual := strings.Split(clean, " ")
				num, color := individual[0], individual[1]
				numI, _ := strconv.Atoi(num)

				if color == "red" {
					if numI > redMax {
						failed = true
						break
					}

				} else if color == "green" {
					if numI > greenMax {
						failed = true
						break
					}

				} else if color == "blue" {
					if numI > blueMax {
						failed = true
						break
					}

				} else {
					fmt.Println("Undefinded color found: ", color)
				}

			}
			if failed {
				break
			}

		}

		if !failed {
			total += curr
		}
	}

	fmt.Println("total: ", total)
}
