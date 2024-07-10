package main

import (
	"fmt"
)

type (
	Player struct{}
	Item   struct {
		Name      string
		Abilities []Ability
	}
	Ability interface {
		Execute()
	}

	Heal        struct{}
	DamageBluff struct{}
	Healmana    struct{}
)

func (p *Player) Use(item Item) {
	fmt.Println("Use:", item.Name)

	for _, item := range item.Abilities {
		item.Execute()
	}
}

func (h Heal) Execute() {
	fmt.Println("Heal")
}

func (d DamageBluff) Execute() {
	fmt.Println("Increase Damage 100%")
}

func (d Healmana) Execute() {
	fmt.Println("Increase mana 100 unit")
}

func main() {
	p := &Player{}
	// steak := Item{
	// 	Name: "Steak",
	// 	Abilities: []Ability{
	// 		Heal{},
	// 		DamageBluff{},
	// 	},
	// }

	bottlehealnama := Item{
		Name: "Bottlehealnama",
		Abilities: []Ability{
			Healmana{},
		},
	}

	// p.Use(steak)
	p.Use(bottlehealnama)
}
