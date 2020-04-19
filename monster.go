//Functions relating to the monsters in the game

package main

import (
	"fmt"
)

type monster struct {
	health       int
	species      string
	attackDamage int
}

//monster attack method
func (m monster) Attack(p player) player {
	fmt.Printf("The %v attacked you and did %v damage!\n", m.species, m.attackDamage)
	p.health -= m.attackDamage
	return p
}
