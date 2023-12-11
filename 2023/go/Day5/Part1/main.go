package main

import (
	"bufio"
	"fmt"
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
	seeds := toIntArr(strings.Split(firstLine, ":")[1])
	scanner.Scan() // we move one next line
	currRes := seeds

	r := []ranges{}

	for scanner.Scan() {

		input := scanner.Text()

		if strings.Contains(input, ":") {
			r = []ranges{}
			continue
		} else if input == "" {
			//evaluate
			currRes = evaluate(r, currRes)
			continue
		}

		arr := toIntArr(input)
		r = append(r, ranges{startDestination: arr[0],
			startSource: arr[1],
			sharedRange: arr[2],
		})

	}

	lowest := currRes[0]

	for _, c := range currRes {
		if c < lowest {
			lowest = c
		}
	}
	fmt.Println(lowest)

}

func evaluate(r []ranges, curr []int) []int {

	res := make([]int, len(curr))

	for i, c := range curr {

		found := false

		for _, currRange := range r {

			if c >= currRange.startSource && c < (currRange.startSource+currRange.sharedRange) {
				found = true
				res[i] = c + currRange.startDestination - currRange.startSource
			}
		}

		if !found {
			res[i] = c
		}
	}

	return res

}

func toIntArr(s string) []int {
	strArr := strings.Split(strings.TrimSpace(s), " ")
	intArr := []int{}
	for _, s := range strArr {
		i, _ := strconv.Atoi(s)
		intArr = append(intArr, i)
	}

	return intArr
}
