//Anything that relates to the player

package main

import (
	"bufio"
	"fmt"
	"os"
)

type player struct {
	health    int
	inventory []string
	xp        int
	gold      int
	hasKey    bool
}

//prints the player's stats
func (p player) PrintStats() {
	fmt.Printf("\n\n***************************************************\n\n")
	fmt.Printf("Current Health: %v\nInventory: %v\nXP: %v\n", p.health, p.inventory, p.xp)
	fmt.Printf("\n\n***************************************************\n\n")
}

//player attack method
func (p player) Attack(m monster) monster {
	damage := 10
	fmt.Printf("You atacked the %v and did %v damage!\n", m.species, damage)
	m.health -= damage
	return m
}

//you died
func dead() {
	fmt.Printf("\n\n***************************************************\n\n")
	fmt.Println("You died like so many before you have...")
	fmt.Println("Bring more potions next time...")
	fmt.Println("Would you like to play again?")

	//// TODO: test this
	input := bufio.NewReader(os.Stdin)
	answer, _ := input.ReadString('\n')

	if answer == "yes" {
		main()
	} else {
		fmt.Println("Thanks for playing!")
		os.Exit(0)
	}
}
