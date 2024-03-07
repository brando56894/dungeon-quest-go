//Anything that relates to the Player

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/GeistInDerSH/clearscreen"
	"github.com/fatih/color"
)

type Player struct {
	Health         int
	Weapons        []string
	Potions        int
	XP             int
	Gold           int
	HasKey         bool
	MonstersKilled int
}

// playerAttack Player attack method
func (p Player) playerAttack(m Monster) Monster {
	clearscreen.ClearScreen()
	damage := rollDice(2, 8)
	color.Green("You attacked the %v and did %v damage!\n", m.Species, damage)
	m.Health -= damage
	return m
}

// playerStatus prints the player's vital stats and inventory
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

// drink a health potion
func usePotion(p Player) Player {
	if p.Potions >= 1 {
		p.Potions -= 1
		p.Health += rollDice(3, 10)
		color.Cyan("You drank a health potion. Your health is now %v\n", p.Health)
	} else {
		fmt.Println("You don't have any health potions.")
	}
	return p
}

// you died
func dead(p Player) {
	color.Red("\nYou died like so many before you have...")
	fmt.Println("Bring more potions next time...")
	color.White("You killed %v monsters before you died. \n\n", p.MonstersKilled)
	fmt.Println("Would you like to play again?")

	input := bufio.NewReader(os.Stdin)
	answer, _ := input.ReadString('\n')

	if answer == "yes\n" {
		main()
	} else {
		fmt.Printf("\nThanks for playing!\n")
		os.Exit(0)
	}
}
