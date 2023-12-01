package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

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
		d := getDigits(line)

		n, err := strconv.Atoi(d)
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
