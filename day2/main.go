package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// returns an int for the number of occurences of substring in a larger string
func countOccurrences(s, word string) int {
	return strings.Count(s, word)
}

// gets all indexes of a substring in a larger string
func getAllIndexes(s, sub string) []int {
	var indexes []int
	for i := 0; ; i++ {
		index := strings.Index(s, sub)
		if index == -1 {
			break
		}
		indexes = append(indexes, index+i)
		s = s[index+1:]
	}
	return indexes
}

// returns the first index of substring in a larger string
func getIndex(s, sub string) int {

	index := strings.Index(s, sub)

	return index
}

// checks a string for an array of numbered/int substring values and converts them to string ints
func parseStringInts(s string) string {

	strIntMap := map[string]string{
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
	strIntKeysArry := make([]string, 0, len(strIntMap))
	for k, _ := range strIntMap {
		strIntKeysArry = append(strIntKeysArry, k)
	}

	lineIndexes := make(map[string]int)
	for _, v := range strIntKeysArry {
		if strings.Contains(s, v) {
			lineIndexes[v] = getIndex(s, v)
		}
	}

	keys := make([]string, 0, len(lineIndexes))
	for key := range lineIndexes {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return lineIndexes[keys[i]] < lineIndexes[keys[j]]
	})

	fmt.Println("Value of string: ", s)
	fmt.Println("lineIndexes: ", keys)

	if len(keys) == 1 {
		fmt.Println(strIntMap[keys[0]] + strIntMap[keys[0]])
		return strIntMap[keys[0]] + strIntMap[keys[0]]

	} else {
		fmt.Println(strIntMap[keys[0]] + strIntMap[keys[len(keys)-1]])
		return strIntMap[keys[0]] + strIntMap[keys[len(keys)-1]]
	}
}

// go through string starting from 1st char all the way to the end until there is a match and break
// flip the string and do the same damn thing

func main() {
	total := 0
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := parseStringInts(line)

		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		total += n
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
