package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/GeistInDerSH/clearscreen"
	"github.com/fatih/color"
	"github.com/justinian/dice"
)

const version = "1.2"

func main() {
	//seeding the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	clearscreen.ClearScreen()
	fmt.Println("Welcome to Dungeon Quest!")
	fmt.Printf("Version: %v\n\n", version)

	//creates the Player object from the struct in Player.go
	adventurer := Player{
		Health:  175,
		Weapons: []string{"dagger"},
		XP:      0,
		Gold:    20,
		Potions: 1,
	}

	//starts the actual game
	menu(adventurer)
}

func menu(p Player) {
	fmt.Printf("Your Health: %v\n\n", p.Health)
	fmt.Println(strings.Repeat("*", 15))
	fmt.Println("* (E)xplore")
	fmt.Println("* (P)layer Status")
	fmt.Println("* (S)tore")
	fmt.Println("* (Q)uit")
	fmt.Println(strings.Repeat("*", 15))
	fmt.Printf("\nWhat would you like to do?  ")
	input := bufio.NewReader(os.Stdin)
	answer, _ := input.ReadString('\n')

	switch strings.ToLower(strings.TrimRight(answer, "\n")) {
	case "e":
		explore(p)
		time.Sleep(4 * time.Second)
		clearscreen.ClearScreen()
	case "p":
		playerStatus(p)
	case "q":
		os.Exit(0)
	case "s":
		fmt.Println("Store is closed. Come back later.")
		time.Sleep(3 * time.Second)
		clearscreen.ClearScreen()
		menu(p)
	default:
		color.Red("Unknown option")
		time.Sleep(3 * time.Second)
		clearscreen.ClearScreen()
		menu(p)
	}
}

func playerStatus(p Player) {
	fmt.Println()
	color.Magenta("Health: %v\n", p.Health)
	color.Magenta("Experience Points: %v\n", p.XP)
	color.Magenta("Carried Gold: %v\n", p.Gold)
	color.Magenta("Weapons: %v\n", p.Weapons[0]) //only one weapon right now
	color.Magenta("Potions: %v\n", p.Potions)
	color.Magenta("Has Key: %v\n", p.HasKey)
	time.Sleep(3 * time.Second)
	clearscreen.ClearScreen()
	menu(p)
}

// rollDice rolls x dice with y number of sides
func rollDice(numberOfDice int, diceSides int) int {
	result, _, err := dice.Roll(fmt.Sprintf("%vd%v", numberOfDice, diceSides))
	//color.Yellow("result: %v, string: %v, err: %v\n", result, string, err)
	//time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("dice.Roll: ", err)
		os.Exit(1)
	}

	return result.Int()
}

// main part of the game
func explore(p Player) {
	clearscreen.ClearScreen()

	switch rollDice(1, 6) {
	//these are all in dice-rolls.go
	case 1:
		p = monsterAttack(p)
	case 2:
		foundGold := rollDice(1, 15)
		color.Yellow("You found %v pieces of gold!\n", foundGold)
		p.Gold += foundGold
	case 3:
		p = itsATrap(p)
	case 4:
		p = lockedDoor(p)
	case 5:
		p = deadGuy(p)
	default:
		fmt.Println("Nothing happened. You're safe.")
	}

	if p.Health <= 0 {
		dead(p)
	} else {
		time.Sleep(3 * time.Second)
		clearscreen.ClearScreen()
		menu(p)
	}
}
