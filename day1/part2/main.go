package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// checks a string for regex of numbered/int substring values and returns them
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

	first := ""
	last := ""
	strIntKeysArry := make([]string, 0, len(strIntMap))
	for k, _ := range strIntMap {
		strIntKeysArry = append(strIntKeysArry, k)
	}

	// Join the strings with the pipe character, which represents "OR" in regex
	pattern := strings.Join(strIntKeysArry, "|")

	// Compile the regex
	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		log.Fatal(err)
	}

	// Find first match going forward in string
	for i, _ := range s {
		if r.MatchString(string(s[0:i])) {
			first = r.FindString(string(s[0:i]))
			break
		}

	}

	// Find first match going backward in string
	for i, _ := range s {
		if r.MatchString(string(s[len(s)-i:])) {
			last = r.FindString(string(s[len(s)-i:]))
			break
		}
	}

	// some strings just have one number so in those cases its for both first/last
	if first == "" {
		first = last
	} else if last == "" {
		last = first
	}

	return strIntMap[first] + strIntMap[last]
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

// total should be 55614
