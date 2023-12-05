package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

var blue = 14
var red = 12
var green = 13

// Game 1: 9 red, 5 blue, 6 green; 6 red, 13 blue; 2 blue, 7 green, 5 red
func getScore(s []byte) int {
	played := bytes.Split(s, []byte(`:`))
	id := bytes.Split(played[0], []byte(` `))[1]

	idInt, _ := strconv.Atoi(string(id))

	for _, pull := range bytes.Split(bytes.Trim(played[1], ` `), []byte(`;`)) {
		for _, bag := range bytes.Split(pull, []byte(`,`)) {

			bag = bytes.Trim(bag, ` `)
			v := bytes.Split(bag, []byte(` `))

			i, _ := strconv.Atoi(string(v[0]))

			switch string(v[1]) {
			case "blue":
				if i > blue {
					return 0
				}
			case "green":
				if i > green {
					return 0
				}
			case "red":
				if i > red {
					return 0
				}
			}
		}
	}

	return idInt

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
