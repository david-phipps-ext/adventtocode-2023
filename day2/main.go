package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// gets digit characters from a string
func getDigits(s string) string {
	strDigits := ""
	for _, c := range s {
		if unicode.IsDigit(c) {
			strDigits += string(c)
		}
	}
	// Need first and last digit
	if len(strDigits) > 2 {
		strDigits = string(strDigits[0]) + string(strDigits[len(strDigits)-1])
	}
	// If there is only one digit we double it
	if len(strDigits) == 1 {
		strDigits = string(strDigits[0]) + string(strDigits[0])
	}
	return strDigits
}

// returns an int for the number of occurences of substring in a larger string
func countOccurrences(s, word string) int {
	return strings.Count(s, word)
}

// returns the indexes of multiple identical substrings in a larger string
func getIndexes(s, sub string) []int {
	var indexes []int
	lastIndex := 0
	for {
		index := strings.Index(s[lastIndex:], sub)
		if index == -1 {
			break
		}
		indexes = append(indexes, index+lastIndex)
		lastIndex = index + lastIndex + len(sub)
	}
	return indexes
}

// checks a string for an array of numbered/int substring values and converts them to string ints
func parseStringInts(s string) string {
	parsedString := ""
	mapVals := make(map[string][]int)

	strToIntMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	// need to range over slice of possible strings found and return the index for their location
	// to the map along with the int value. This will allow us to render the value after and get the
	// first and last number
	for k, _ := range strToIntMap {
		if strings.Contains(s, k) {
			mapVals[strToIntMap[k]] = getIndexes(s, k)
		}
	}

	// we should be able to parse mapVals now and only return the beginning and end values or do a sort
	for k, v := range mapVals {
		first := ""
		last := ""
		end := 1
		for _, i := range v {
			if i == 0 {
				first = k
			}
			if i >= end {
				last = k
				end = i
			}
		}
		parsedString = first + last
	}
	return parsedString
}

func main() {
	total := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Checking line: ", line)
		fmt.Println(parseStringInts(line))
		// s := parseStringInts(line)
		// check and replace number words in d with digit chars
		// d := replaceStringNamesWithInts(line)
		// fmt.Printf("Old line value is %v, new value is %v", line, d)
		// fmt.Println()

		// d := getDigits(s)
		// n, err := strconv.Atoi(d)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// total += n
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
