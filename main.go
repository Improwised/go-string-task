package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(testValidity("123-qwert-678-567"))
}

// testValidity returns true if given string is matching regex
func testValidity(str string) bool {
	matched, err := regexp.MatchString("^([0-9]+-[a-zA-Z]+)+(-?[0-9]+-[a-zA-Z]+)*$", str)
	if err != nil {
		panic(err)
	}
	return matched
}
