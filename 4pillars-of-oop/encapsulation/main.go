package main

import "fmt"


//prevent call child from external 
type (
	PlayerClass interface {
		Attack()
	}

	Warrior struct {
		Weapon string
	}

	Mage struct {
		Spell string
	}
)

func (w Warrior) Attack() {
	fmt.Println("Warrior attack by: ", w.Weapon)
}

func (m Mage) Attack() {
	fmt.Println("Mage attack by: ", m.Spell)
}

func PlayerAttack(playerclass PlayerClass) {
	playerclass.Attack()
}

// same thing can be different thing but have similar methods

func main() {
	warrior := Warrior{
		Weapon: "Sword",
	}
	mage := Mage{
		Spell: "FireBall",
	}

	// warrior.Attack()
	// mage.Attack()

	PlayerAttack(warrior)
	PlayerAttack(mage)
}
