package layout

import (
	"fmt"
)

func WelcomeLayout() {
	// fmt.Print("\033[H\033[2J")
	fmt.Println("--------------------------------------------------")
	fmt.Println("------------Welcome to the Warteg Apps!-----------")
	fmt.Println("You can order food and drinks here by inputting \nyour order in the terminal.")
	fmt.Println("--------------------------------------------------")
	
}