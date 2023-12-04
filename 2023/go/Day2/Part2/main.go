package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	redMax, greenMax, blueMax := 0, 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	curr := 0

	for {
		curr += 1
		redMax, greenMax, blueMax = 0, 0, 0

		var input string
		if scanner.Scan() {
			input = scanner.Text()
		} else {
			break
		}

		res := strings.Split(input, ":")
		// we can safely ignore the first part of the separate line
		data := strings.Split(res[1], ";")

		for _, draw := range data {
			drawArr := strings.Split(draw, ",")
			for _, numAndColor := range drawArr {
				clean := strings.TrimSpace(numAndColor)
				individual := strings.Split(clean, " ")
				num, color := individual[0], individual[1]
				numI, _ := strconv.Atoi(num)

				if color == "red" {
					if numI > redMax {
						redMax = numI
					}
				} else if color == "green" {
					if numI > greenMax {
						greenMax = numI
					}
				} else if color == "blue" {
					if numI > blueMax {
						blueMax = numI
					}
				} else {
					fmt.Println("Undefinded color found: ", color)
				}

			}
		}

		total += (blueMax * greenMax * redMax)
	}

	fmt.Println("total: ", total)
}
