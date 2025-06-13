package main

import (
	// "fmt"
	"fmt"
	"os"
	"sync"
	"weeklytask/data"
	"weeklytask/layout"
	"weeklytask/models"
	"weeklytask/utils"
)

func main() {
	ci := data.CurrentInput
	menu := data.ListMenu
	wg := sync.WaitGroup{}
	ch1 := make(chan []models.Menu)
	var cart []models.Menu
	
	// for {
		layout.Main_Layout()
		
		fmt.Println("current cart: ",cart)
		
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("What can i help you?")
			fmt.Println("1. Show All Menu")
			fmt.Println("2. Cart")
			fmt.Println("3. Checkout")
			fmt.Println("4. Exit")
			fmt.Println("--------------------------------------------------")
			fmt.Print("Input: ")
			_, err := fmt.Scanln(&ci)
			if err != nil {
				fmt.Println("Invalid input")
			}
		}()
		wg.Wait()
		
		switch ci {
		case "1":
			wg.Add(1)
			go utils.ShowMenu(menu, ch1, &wg)
			cart = <-ch1
			wg.Wait()
		case "2":
			fmt.Println("my cart: ", cart)
		case "3":
			fmt.Println("my checkout")
		case "4":
			os.Exit(0)
		default:
			fmt.Println("your input is invalid!")
			return
		}
	// }
}	