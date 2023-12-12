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

	for i := 0; i < len(currRes); i += 2 {
		c := currRes[i]
		if c < lowest {
			lowest = c
		}
	}
	fmt.Println(lowest)

}

func evaluate(r []ranges, curr []int) []int {

	res := []int{}

	for i := 0; i < len(curr); i += 2 {

		currStart, currEnd := curr[i], curr[i+1]+curr[i]

		for currStart < currEnd {

			found := false

			for _, currRange := range r {

				currRangeEnd := currRange.startSource + currRange.sharedRange

				if currStart >= currRange.startSource && currStart < currRangeEnd {
					found = true
					newRes := currStart + currRange.startDestination - currRange.startSource
					res = append(res, newRes)

					// if the end is within the range then we will add the whole range
					if currEnd <= currRangeEnd {
						res = append(res, currEnd-currStart)
						currStart = currEnd
					} else {
						// otherwise we will only take the interval within the range
						res = append(res, currRangeEnd-currStart)
						currStart = currRangeEnd
					}

				}
			}

			if !found {
				// we find the smallest startSource that is bigger than currStart

				smallest := currEnd

				for _, currRange := range r {
					if currRange.startSource > currStart && currRange.startSource < smallest {
						smallest = currRange.startSource
					}
				}

				res = append(res, currStart)
				res = append(res, smallest-currStart)
				currStart = smallest

			}
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
