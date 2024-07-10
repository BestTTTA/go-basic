package main

import "fmt"


//same thing can be different in your mind.
type (
	Keyboard interface {
		Input()
	}

	MechanicalKeyboard struct {
		SwitchType string
		Size       string
		OS         string
	}

	NormalKeyboard struct {
		IscanPress bool
	}
)

func (m MechanicalKeyboard) Input() {
	fmt.Println("Data mechanicalkeyboard is: ", m.SwitchType, m.Size, m.OS)
}
func (m NormalKeyboard) Input() {
	fmt.Println("Data mechanicalkeyboard is: ", m.IscanPress)
}

func main() {

	mechanicalKeyboard := MechanicalKeyboard{
		SwitchType: "some-type",
		Size: "big",
		OS: "windows, macOS",
	}

	normalKeyboard := NormalKeyboard{
		IscanPress: true,
	}

	mechanicalKeyboard.Input()
	normalKeyboard.Input()


}
