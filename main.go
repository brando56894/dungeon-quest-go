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

// explore moves the player through the dungeon
func explore(p Player) {
	clearscreen.ClearScreen()

	switch rollDice(1, 6) {
	case 1:
		p = monsterAttack(p)
	case 2:
		foundGold := rollDice(1, 15)
		color.Yellow("You found %v pieces of gold!\n", foundGold)
		p.Gold += foundGold
	case 3:
		fmt.Println("You walk down a tunnel and feel something odd under your foot...")
		var damage int

		switch rollDice(1, 3) {
		case 1:
			fmt.Println("AHHH!!FIRE!!!! IT BURNS!!!")
			damage = rollDice(1, 14)
		case 2:
			fmt.Println("You were shot by an arrow trap!")
			damage = rollDice(1, 4)
		case 3:
			fmt.Println("You were hit by a spring-loaded spiked trap!")
			damage = rollDice(1, 8)
		}

		p.Health -= damage
		color.Red("You lost %v health\n", damage)
	case 4:
		fmt.Println("You reach a door and you turn the handle. It's locked though....")
		if p.HasKey {
			foundGold := rollDice(2, 10)
			fmt.Println("You use the key you found earlier...")
			color.Yellow("Inside the room you found a chest with %v gold in it!\n", foundGold)
			p.Gold += foundGold
		}
	case 5:
		fmt.Println("You find a dead body...looks like he died a while ago...")
		fmt.Println("You look through the pockets and backpack...")
		switch rollDice(1, 4) {
		case 1:
			foundGold := rollDice(1, 15)
			color.Yellow("You found %v gold in the backpack!\n", foundGold)
			p.Gold += foundGold
		case 2:
			if p.HasKey {
				foundPotions := rollDice(1, 4)
				p.Potions += foundPotions
				color.Cyan("You find %v potions!\n", foundPotions)
			} else {
				fmt.Println("You found a key in their pocket and wonder if it will be useful later.")
				p.HasKey = true
			}
		case 3:
			foundPotions := rollDice(1, 4)
			p.Potions += foundPotions
			color.Cyan("You find %v potions!\n", foundPotions)
		case 4:
			fmt.Println("But they're empty...looks like someone already got to him.")
		}
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
