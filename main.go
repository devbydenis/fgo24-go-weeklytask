package main

import (
	// "fmt"
	"fmt"
	"os"
	"sync"
	"time"
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
	var checkout []models.Menu
	
	
	layout.WelcomeLayout()
	time.Sleep(2 * time.Second)
	
	for {
		
		wg.Add(1)
		fmt.Print("\033[H\033[2J")
		go layout.MenuLayout(&ci, &wg)
		wg.Wait()
		
		switch ci {
		case "1":
			wg.Add(1)
			go utils.ShowMenu(menu, ch1, &wg)
			cart = <-ch1
			wg.Wait()
		case "2":
			fmt.Print("\033[H\033[2J")
			fmt.Println("======================")
			fmt.Println("my cart: ", cart)
			fmt.Println("======================")
			fmt.Println("[0] back to menu")
			fmt.Println("[1] checkout my chart")
			fmt.Print("Input: ")
			_, err := fmt.Scanln(&ci)
			if err != nil {
				fmt.Println("Invalid input")
			}
			
			if ci == "0" {
				fmt.Println("you choose 0")
				continue
			} else if ci == "1" {
				checkout = append(checkout, cart...)
				fmt.Println("checkout success")
				time.Sleep(1 * time.Second)
				continue
			} else {
				fmt.Println("your input is invalid!")
				time.Sleep(1 * time.Second)
				continue
			}
		case "3":
			fmt.Print("\033[H\033[2J")
			fmt.Println("my checkout: ", checkout)
			fmt.Print("Input: ")
			_, err := fmt.Scanln(&ci)
			if err != nil {
				fmt.Println("Invalid input")
			}
		case "4":
			fmt.Println("Thanks for using Warteg Apps. See You!")
			time.Sleep(2 * time.Second)
			os.Exit(0)
		default:
			fmt.Println("your input is invalid!")
			return
		}
	}
}	