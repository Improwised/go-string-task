package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")
var digits = []rune("1234567890")

func main() {
	rand.Seed(time.Now().UnixNano())
	generated := generate(true)
	fmt.Println(generated)
	fmt.Println(generate(false))
	fmt.Println(testValidity(generated))
	fmt.Println(averageNumber(generated))
	fmt.Println(wholeStory(generated))
	fmt.Println(storyStats(generated))
}

func testValidity(str string) bool {
	matched, err := regexp.MatchString("^([0-9]+-[a-zA-Z]+)+(-?[0-9]+-[a-zA-Z]+)*$", str)
	if err != nil {
		panic(err)
	}
	return matched
}

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

func storyStats(str string) ([]string, []string, int, []string) {
	var shortestWords []string
	var longesttWords []string
	var avgLenWords = []string{}
	var shortestWordsLen = len(str)
	var longestWordsLen, avgWordLen, totalWordLen, intCounter = 0, 0, 0, 0

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

// generate will generate valid/invalid method based on argument
func generate(correct bool) string {
	var minRep = 2
	var minSubStrLen = 2
	var subStrs []string
	if correct {
		for i := 0; i < rand.Intn(10)+minRep; i++ {
			subStrs = append(subStrs, getSubStr(rand.Intn(5)+minSubStrLen, digits))
			subStrs = append(subStrs, getSubStr(rand.Intn(5)+minSubStrLen, letters))
		}
		return strings.Join(subStrs, "-")
	}
	return getSubStr(rand.Intn(12)+5, letters)
}

func getSubStr(n int, choices []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = choices[rand.Intn(len(choices))]
	}
	return string(b)
}
