package main

import "fmt"

//same class could have attribute and methods
type Player struct {
	Username string
	Level    uint
	Status   string
	Class    string
}

func (p *Player) LevelUp() {
	p.Level++
	fmt.Println(p.Level)
}

func main() {
	player1 := Player{
		Username: "player1",
		Level:    0,
		Status:   "active",
		Class:    "warrior",
	}

	for range 10 {
		player1.LevelUp()
	}

}
