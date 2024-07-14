package main

import "fmt"

func main() {
	// for range 5{
	// 	func2()
	// 	func3()
	// 	fmt.Println()
	// }

	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}

	var value int = 2
	switch {
	case value == 1:
		fmt.Println("Hello")
	case value == 2:
		fmt.Println("Bonjour")
	case value == 3:
		fmt.Println("Namstay")
	default:
		fmt.Println("Invalid")
	}

	switch {
	case value == 1:
		fmt.Println("Hello")
	case value == 2:
		fmt.Println("Bonjour")
	case value == 3:
		fmt.Println("Namstay")
	default:
		fmt.Println("Invalid")
	}

}

func func1() {
	fmt.Println("Best")
}

func func2() {
	fmt.Println("jo pha")
}

func func3() {
	fmt.Println("mei")
}

//jo pha,mei,Best
