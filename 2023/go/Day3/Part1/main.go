package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	arr := [][]rune{}
	line_num := 0

	start_y_coors := []int{}
	start_x_coors := []int{}
	end_x_coors := []int{}

	for scanner.Scan() {

		num_found := false

		input := scanner.Text()

		// we store all the chars in an array
		arr = append(arr, []rune{})
		for _, ch := range input {
			arr[line_num] = append(arr[line_num], ch)
		}

		// next we collect all the numbers in a list along with thir x and y coordinates
		for x_coor, ch := range input {
			if num_found && (ch < '0' || ch > '9') {
				num_found = false
				end_x_coors = append(end_x_coors, x_coor-1)
			} else if !num_found && (ch >= '0' && ch <= '9') {
				num_found = true
				start_y_coors = append(start_y_coors, line_num)
				start_x_coors = append(start_x_coors, x_coor)

			}
		}

		if num_found {
			end_x_coors = append(end_x_coors, len(input)-1)
		}

		line_num += 1
	}

	total := 0

	for i, _ := range start_y_coors {
		y, x0, x1 := start_y_coors[i], start_x_coors[i], end_x_coors[i]

		if isSurroundedBySymbol(y, x0, x1, arr) {
			var res int
			if x1+1 >= len(arr[y]) {
				res, _ = strconv.Atoi(string(arr[y])[x0:])
			} else {
				res, _ = strconv.Atoi(string(arr[y])[x0 : x1+1])
			}
			total += res
		}

	}

	fmt.Println("Total: ", total)
}

func isSurroundedBySymbol(y, x0, x1 int, store [][]rune) bool {

	y0, y1 := y, y
	if y > 0 {
		y0 = y - 1
	}
	if y < len(store)-1 {
		y1 = y + 1
	}
	if x0 > 0 {
		x0 -= 1
	}
	if x1 < (len(store[0]) - 1) {
		x1 += 1
	}

	for yInd := y0; yInd <= y1; yInd += 1 {
		for xInd := x0; xInd <= x1; xInd += 1 {
			curr := store[yInd][xInd]
			if (curr < '0' || curr > '9') && curr != '.' {
				return true
			}
		}
	}

	return false
}
