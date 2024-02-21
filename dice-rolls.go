//Functions for dice rolls

package main

import (
	"fmt"
)

type Monster struct {
	health       int
	species      string
	attackDamage int
}

// dice roll #1
func attack(p player) player {
	var monster Monster

	//creating a monster
	monsterNumber := randInt(0, 4)
	switch monsterNumber {
	case 1:
		monster = Monster{
			health:       randInt(9, 16),
			species:      "Draugr",
			attackDamage: randInt(2, 7),
		}
	case 2:
		monster = Monster{
			health:       randInt(11, 19),
			species:      "Demon",
			attackDamage: randInt(4, 9),
		}
	case 3:
		monster = Monster{
			health:       randInt(13, 22),
			species:      "Eyeless Creeper",
			attackDamage: randInt(5, 10),
		}
	}

	//attack sequence
	for monster.health > 0 {
		monster = p.Attack(monster) //player attacks monster
		p = monster.Attack(p)       //monster attacks player
		fmt.Println("")
		if p.health <= 0 {
			fmt.Println("...and everything goes black...")
			//ends the game
			dead()
		} else if monster.health <= 0 {
			fmt.Printf("You killed the %v!\n", monster.species)
			break
		}
	}
	return p
}

// dice roll #2
func findGold(p player) player {
	foundGold := randInt(1, 15)
	fmt.Printf("You found %v pieces of gold!\n", foundGold)
	p.gold += foundGold
	return p
}

// dice roll #3
func itsATrap(p player) player {
	fmt.Println("You walk down a tunnel and feel something odd under your foot...")
	number := randInt(0, 3)

	switch number {
	case 0:
		fmt.Println("Eww! It's a dead rat!")
	case 1:
		fmt.Println("AHHH!!FIRE!!!!\nIT BURNS!!!")
		damage := randInt(7, 14)
		p.health -= damage
		fmt.Printf("You lost %v health\n", damage)
	case 2:
		fmt.Println("You were shot by an arrow trap!")
		damage := randInt(1, 4)
		p.health -= damage
		fmt.Printf("You lost %v health\n", damage)
	case 3:
		fmt.Println("You were hit by a spring-loaded spiked trap!")
		damage := randInt(3, 8)
		p.health -= damage
		fmt.Printf("You lost %v health\n", damage)
	}
	return p
}

// dice roll #4
func lockedDoor(p player) player {
	fmt.Println("You reach a door and you turn the handle. It's locked though....")
	if p.hasKey {
		foundGold := randInt(10, 40)
		fmt.Println("You use the key you found earlier...")
		fmt.Printf("Inside the room you found a chest with %v gold in it!\n", foundGold)
		p.gold += foundGold
	}
	return p
}

// dice roll #5
func deadGuy(p player) player {
	fmt.Println("You find a dead body...looks like he died a while ago...")
	fmt.Println("You look through the pockets and backpack...")
	number := randInt(0, 3)
	if number == 1 {
		foundGold := randInt(1, 30)
		fmt.Printf("You found %v gold in the backpack!\n", foundGold)
		p.gold += foundGold
	} else if number == 2 {
		fmt.Println("You found a key in their pocket and wonder if it will be useful later.")
		p.hasKey = true
	} else {
		fmt.Println("But they're empty...looks like someone already got to him.")
	}
	return p
}
