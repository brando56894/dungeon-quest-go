//Functions for dice rolls

package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Monster struct {
	Health       int
	Species      string
	AttackDamage int
}

// case 2
func itsATrap(p Player) Player {
	fmt.Println("You walk down a tunnel and feel something odd under your foot...")
	var damage int

	switch randInt(1, 3) {
	case 1:
		fmt.Println("AHHH!!FIRE!!!! IT BURNS!!!")
		damage = randInt(7, 14)
	case 2:
		fmt.Println("You were shot by an arrow trap!")
		damage = randInt(1, 4)
	case 3:
		fmt.Println("You were hit by a spring-loaded spiked trap!")
		damage = randInt(3, 8)
	}
	p.Health -= damage
	color.Red("You lost %v health\n", damage)
	return p
}

// case 3
func lockedDoor(p Player) Player {
	fmt.Println("You reach a door and you turn the handle. It's locked though....")
	if p.HasKey {
		foundGold := randInt(10, 40)
		fmt.Println("You use the key you found earlier...")
		color.Yellow("Inside the room you found a chest with %v gold in it!\n", foundGold)
		p.Gold += foundGold
	}
	return p
}

// case 4
func deadGuy(p Player) Player {
	fmt.Println("You find a dead body...looks like he died a while ago...")
	fmt.Println("You look through the pockets and backpack...")
	switch randInt(1, 4) {
	case 1:
		foundGold := randInt(1, 15)
		color.Yellow("You found %v gold in the backpack!\n", foundGold)
		p.Gold += foundGold
	case 2:
		if p.HasKey {
			foundPotions := randInt(1, 4)
			p.Potions += foundPotions
			color.Cyan("You find %v potions!\n", foundPotions)
		} else {
			fmt.Println("You found a key in their pocket and wonder if it will be useful later.")
			p.HasKey = true
		}
	case 3:
		foundPotions := randInt(1, 4)
		p.Potions += foundPotions
		color.Cyan("You find %v potions!\n", foundPotions)
	case 4:
		fmt.Println("But they're empty...looks like someone already got to him.")
	}

	return p
}
