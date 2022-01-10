package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(testValidity("123-qwert-678-tgfhg"))
	fmt.Println(averageNumber("123-qwert-678-tgfhg"))
	fmt.Println(wholeStory("123-qwert-678-tgfhg"))
	fmt.Println(storyStats("123-qwert-678-tgfhg"))
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

// wholeStory returns all words of given string joined by space
func wholeStory(str string) string {
	var story []string
	if testValidity(str) {
		values := strings.Split(str, "-")
		for index, val := range values {
			if index%2 != 0 {
				story = append(story, val)
			}
		}
		return strings.Join(story, " ")
	}
	return ""
}

// storyStats will return shortest word, longest word, average length of all words and words which have average length
func storyStats(str string) ([]string, []string, int, []string) {
	var avgLenWords = []string{}
	var shortestWords []string
	var shortestWordsLen = len(str)
	var longestWordsLen = 0
	var longesttWords []string
	var avgWordLen = 0
	var totalWordLen = 0
	var intCounter = 0

	if testValidity(str) {
		values := strings.Split(str, "-")
		for index, val := range values {
			if index%2 != 0 {
				// Check is it shortest
				if len(val) < shortestWordsLen {
					shortestWordsLen = len(val)
					shortestWords = []string{val}
				} else if len(val) == shortestWordsLen {
					shortestWords = append(shortestWords, val)
				}
				// Check is it longest
				if len(val) > longestWordsLen {
					longestWordsLen = len(val)
					longesttWords = []string{val}
				} else if len(val) == longestWordsLen {
					longesttWords = append(longesttWords, val)
				}
				intCounter++
				totalWordLen = totalWordLen + len(val)
			}
		}
		avgWordLen = totalWordLen / intCounter
		for index, val := range values {
			if index%2 != 0 && len(val) == avgWordLen {
				avgLenWords = append(avgLenWords, val)
			}
		}
		return shortestWords, longesttWords, avgWordLen, avgLenWords
	}
	return []string{}, []string{}, 0, []string{}
}
