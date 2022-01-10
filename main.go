package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(testValidity("123-qwert-678-567"))
	fmt.Println(averageNumber("123-qwert-678-567"))
}

// testValidity returns true if given string is matching regex
func testValidity(str string) bool {
	matched, err := regexp.MatchString("^([0-9]+-[a-zA-Z]+)+(-?[0-9]+-[a-zA-Z]+)*$", str)
	if err != nil {
		panic(err)
	}
	return matched
}

// averageNumber will return average of all words in float64
func averageNumber(str string) float64 {
	var result float64
	var intCounter = 0
	if testValidity(str) {
		values := strings.Split(str, "-")
		for index, val := range values {
			if index%2 == 0 {
				intVal, err := strconv.Atoi(val)
				if err != nil {
					return 0
				}
				intCounter++
				result = result + float64(intVal)
			}
		}
		return result / float64(intCounter)
	}
	return 0
}
