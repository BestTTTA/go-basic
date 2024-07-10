package main

import "fmt"

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

// same thing can be different thing but have similar methods
func main() {
	warrior := Warrior{
		Weapon: "Sword",
	}
	mage := Mage{
		Spell: "FireBall",
	}

	warrior.Attack()
	mage.Attack()
}
