package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// game type to hold our number of red, green and blue cubes
type game struct {
	id, red, green, blue int
}

func newGame(i, r, g, b int) game {
	return game{i, r, g, b}
}

// parse input file for gamevalues
func getGame(s string) game {
	// total := 0
	game := newGame(0, 0, 0, 0)
	gameIDCaps := `Game\s(\d+)`
	blueCaps := `(\d+)\sblue`
	greenCaps := `(\d+)\sgreen`
	redCaps := `(\d+)\sred`

	// re search for game ids
	re := regexp.MustCompile(gameIDCaps)
	matches := re.FindAllStringSubmatch(s, -1)
	game.id, _ = strconv.Atoi(matches[0][1])

	// re search for blue digits and then total number up
	re = regexp.MustCompile(blueCaps)
	matches = re.FindAllStringSubmatch(s, -1)
	for _, v := range matches {
		for _, k := range v {
			i, _ := strconv.Atoi(k)
			game.blue += i
		}

	}

	// re search for green digits and then total number up
	re = regexp.MustCompile(greenCaps)
	matches = re.FindAllStringSubmatch(s, -1)
	for _, v := range matches {
		for _, k := range v {
			i, _ := strconv.Atoi(k)
			game.green += i
		}

	}

	// re search for red digits and then total number up
	re = regexp.MustCompile(redCaps)
	matches = re.FindAllStringSubmatch(s, -1)
	for _, v := range matches {
		for _, k := range v {
			i, _ := strconv.Atoi(k)
			game.red += i
		}

	}

	return game
}

func main() {
	idSum := 0
	//winningGame := newGame(0, 12, 13, 14)

	games := []game{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game := getGame(line)
		games = append(games, game)
	}

	for _, v := range games {
		if v.blue <= 14 &&
			v.red <= 12 &&
			v.green <= 13 {
			fmt.Println("Adding gameid %v", v.id)
			// fmt.Println()
			idSum += v.id
		}
	}
	fmt.Println(idSum)
	for _, v := range games {
		fmt.Println(v)
		fmt.Println()
	}

}

//game{i, r, g, b}
