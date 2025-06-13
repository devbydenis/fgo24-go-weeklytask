package layout

import (
	"fmt"
)

func Main_Layout(currentInput *string) {
	fmt.Println("--------------------------------------------------")
	fmt.Println("--------------Welcome to the Warteg!--------------")
	fmt.Println("You can order food and drinks here by inputting \nyour order in the terminal.")
	fmt.Println("--------------------------------------------------")
	fmt.Println("What can i help you?")
	fmt.Println("1. Show me Food List")
	fmt.Println("2. Show me Drink List")
	fmt.Println("3. My cart")
	fmt.Println("4. Checkout")
	fmt.Println("5. Exit")
	fmt.Println("--------------------------------------------------")
	fmt.Print("Input: ")
	_, err := fmt.Scanln(currentInput)
	if err != nil {
		fmt.Println("Invalid input")
	}
	
}