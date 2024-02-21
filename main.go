package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/GeistInDerSH/clearscreen"
)

const version = "1.1.1"

func main() {
	//seeding the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	clearscreen.ClearScreen()
	fmt.Println("Welcome to Dungeon Quest!")
	fmt.Printf("Version: %v\n\n", version)

	//creates the player object from the struct in player.go
	adventurer := player{
		health:    100,
		inventory: []string{"dagger"},
		xp:        0,
		gold:      0,
	}

	//starts the actual game
	menu(adventurer)
}

func menu(p player) {
	fmt.Println(strings.Repeat("*", 15))
	fmt.Println("* Explore")
	fmt.Println("* Store")
	fmt.Println("* Exit")
	fmt.Println(strings.Repeat("*", 15))
	fmt.Printf("\nWhat would you like to do?  ")
	input := bufio.NewReader(os.Stdin)
	answer, _ := input.ReadString('\n')

	switch strings.ToLower(strings.TrimRight(answer, "\n")) {
	case "explore":
		explore(p)
		time.Sleep(5 * time.Second)
		clearscreen.ClearScreen()
	case "exit":
		os.Exit(0)
	case "store":
		fmt.Println("Store is closed. Come back later.")
		time.Sleep(3 * time.Second)
		clearscreen.ClearScreen()
		menu(p)
	default:
		fmt.Println("Unknown option")
		time.Sleep(3 * time.Second)
		clearscreen.ClearScreen()
		menu(p)
	}
}

// generates a random integer between two integers and returns it
func randInt(min, max int) int {
	return 1 + rand.Intn(max-min)
}

// main part of the game
func explore(p player) {
	clearscreen.ClearScreen()
	//roll a 6 sided dice
	diceRoll := randInt(1, 6)
	fmt.Printf("You rolled a %v\n", diceRoll)

	switch diceRoll {
	//these are all in dice-rolls.go
	case 1:
		p = attack(p)
	case 2:
		p = findGold(p)
	case 3:
		p = itsATrap(p)
	case 4:
		p = lockedDoor(p)
	case 5:
		p = deadGuy(p)
	case 6:
		fmt.Println("And nothing happened. You're safe.")
	}

	if p.health <= 0 {
		dead()
	} else {
		time.Sleep(3 * time.Second)
		clearscreen.ClearScreen()
		menu(p)
	}
}
