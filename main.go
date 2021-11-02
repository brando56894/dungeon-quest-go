package main

import (
	"fmt"
	"math/rand"
	"time"
)

const version = "1.1"

//generates a random integer between to integers and returns it
func randInt(min, max int) int {
	return 1 + rand.Intn(max-min)
}

//main part of the game
func explore(p player) {
	fmt.Printf("As you step foot into the dungeon, you mentally prepare yourself for the horrors ahead.\n\n")

	//continues the game until the player dies
	for p.health > 0 {
		fmt.Println("Player Health: ", p.health)
		//roll a 6 sided di
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
			fmt.Printf("\n\n***************************************************\n\n")
			fmt.Println("And nothing happened. You're safe.")
			fmt.Printf("\n\n***************************************************\n\n")
		}
	}
	if p.health <= 0 {
		dead()
	}
}

func main() {
	//seeding the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

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
	explore(adventurer)
}
