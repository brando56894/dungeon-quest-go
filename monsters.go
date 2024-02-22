package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createMonster() Monster {
	var monster Monster

	//creating a monster
	//fmt.Println("monsterNumber random int")
	monsterNumber := randInt(2, 6)
	switch monsterNumber {
	case 1:
		//fmt.Println("draugr health and attack")
		monster = Monster{
			Health:       randInt(16, 50),
			Species:      "Draugr",
			AttackDamage: randInt(6, 20),
		}
	case 2:
		//fmt.Println("demon health and attack")
		monster = Monster{
			Health:       randInt(22, 70),
			Species:      "Demon",
			AttackDamage: randInt(8, 25),
		}
	case 3:
		//fmt.Println("eyeless creeper health and attack")
		monster = Monster{
			Health:       randInt(26, 90),
			Species:      "Eyeless Creeper",
			AttackDamage: randInt(10, 30),
		}
	default:
		//fmt.Println("default skeleton health and attack")
		monster = Monster{
			Health:       randInt(16, 50),
			Species:      "Skeleton",
			AttackDamage: randInt(6, 20),
		}
	}

	return monster
}

// monster attack method
func (m Monster) monsterAttacksYou(p Player) Player {
	fmt.Printf("The %v attacked you and did %v damage!\n\n", m.Species, m.AttackDamage)
	p.Health -= m.AttackDamage
	return p
}

func monsterAttack(p Player) Player {
	monster := createMonster()
	fmt.Printf("A %v is blocking your path...\n\n", monster.Species)
	for monster.Health > 0 {

		fmt.Println("Your Health: ", p.Health)
		fmt.Println("What are you going to do?")
		fmt.Println("* (A)ttack")
		fmt.Println("* (R)un")
		fmt.Printf("* (U)se potion\n\n")
		input := bufio.NewReader(os.Stdin)
		answer, _ := input.ReadString('\n')

		switch strings.ToLower(strings.TrimRight(answer, "\n")) {
		case "a":
			monster = p.playerAttack(monster)
			p = monster.monsterAttacksYou(p)

			//fmt.Printf("p.Health: %v\nmonster.Health: %v\n\n", p.Health, monster.Health)

			if p.Health <= 0 {
				fmt.Println("\n...and everything goes black...")
				dead(p) //ends the game
			} else if monster.Health <= 0 {
				fmt.Printf("You killed the %v!\n", monster.Species)
				p.MonstersKilled += 1
				break
			}
		case "u":
			usePotion(p)
		case "r":
			if randInt(1, 26)%2 == 0 {
				fmt.Println("You were able to get away.")
				menu(p)
			} else {
				fmt.Println("You weren't able to get away!")
				monster.monsterAttacksYou(p)
			}
		default:
			fmt.Println("Not a valid option")
			monsterAttack(p)

		}
	}

	return p
}
