package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

// var blue = 14
// var red = 12
// var green = 13

// Game 1: 9 red, 5 blue, 6 green; 6 red, 13 blue; 2 blue, 7 green, 5 red
func getScore(s []byte) int {
	played := bytes.Split(s, []byte(`:`))

	red := 0
	blue := 0
	green := 0

	for _, pull := range bytes.Split(bytes.Trim(played[1], ` `), []byte(`;`)) {
		for _, bag := range bytes.Split(pull, []byte(`,`)) {

			bag = bytes.Trim(bag, ` `)
			v := bytes.Split(bag, []byte(` `))

			i, _ := strconv.Atoi(string(v[0]))

			switch string(v[1]) {
			case "blue":
				if blue < i {
					blue = i
				}
			case "green":
				if green < i {

					green = i
				}

			case "red":
				if red < i {
					red = i
				}
			}
		}

	}
	return red * blue * green
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		score += getScore(line)

	}
	fmt.Println(score)
}
