package layout

import (
	"fmt"
	"sync"
)

func MenuLayout(ci *string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("--------------------------------------------------")
	fmt.Println("What can i help you?")
	fmt.Println("--------------------------------------------------")
	fmt.Println("1. Show All Menu")
	fmt.Println("2. Cart")
	fmt.Println("3. Checkout")
	fmt.Println("4. History")
	fmt.Println("5. Exit")
	fmt.Println("--------------------------------------------------")
	fmt.Print("Input: ")
	_, err := fmt.Scanln(ci)
	if err != nil {
		fmt.Println("Invalid input")
	}
}