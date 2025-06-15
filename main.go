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
	var history []models.Menu

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
			fmt.Println("M Y  C A R T")
			fmt.Println("======================")
			for _, item := range cart {
				fmt.Printf("%s \t %d\n", item.Name, item.Price)
			}
			fmt.Println("______________________")
			fmt.Println("[0] Back to menu")
			fmt.Println("[1] Checkout my cart")
			fmt.Print("Input: ")
			_, err := fmt.Scanln(&ci)
			if err != nil {
				fmt.Println("Invalid input")
			}

			if ci == "0" {
				continue
			} else if ci == "1" {
				checkout = append(checkout, cart...)
				fmt.Println("Checkout Success ✔")
				time.Sleep(1 * time.Second)
				continue
			} else {
				fmt.Println("Your Input Is Invalid!")
				time.Sleep(1 * time.Second)
				continue
			}
		case "3":
			fmt.Print("\033[H\033[2J")
			fmt.Println("======================")
			fmt.Println("C H E C K O U T")
			fmt.Println("======================")
			totalPrice := 0
			for _, item := range checkout {
				fmt.Printf("%s\t%d\n", item.Name, item.Price)
				totalPrice += item.Price
			}

			fmt.Println("______________________ +")
			fmt.Printf("Total Price: \t%d\n", totalPrice)
			fmt.Println("======================")
			fmt.Println("[0] Back to menu")
			fmt.Println("[1] Accept and pay")
			fmt.Print("Input: ")
			_, err := fmt.Scanln(&ci)
			if err != nil {
				fmt.Println("Invalid input")
			}

			if ci == "0" {
				continue
			} else if ci == "1" {
				history = append(history, checkout...)
				fmt.Println("payment success ✔")
				time.Sleep(1 * time.Second)
				continue
			} else {
				fmt.Println("your input is invalid!")
				time.Sleep(1 * time.Second)
				continue
			}
		case "4":
			fmt.Print("\033[H\033[2J")
			fmt.Println("======================")
			fmt.Println("H I S T O R Y")
			fmt.Println("======================")

			if len(history) == 0 {
				fmt.Println("No history found")
			}
			for _, item := range history {
				currentTime := time.Now().Format("2006-01-02 15:04:05")
				fmt.Printf("%s\t%s\t%s\t%d\n", currentTime, item.ID, item.Name, item.Price)
			}

			fmt.Println("______________________")
			fmt.Println("[0] Back to menu")
			fmt.Print("Input: ")
			_, err := fmt.Scanln(&ci)
			if err != nil {
				fmt.Println("Invalid input")
			}

			if ci != "0" {
				fmt.Println("your input is invalid!")
				time.Sleep(1 * time.Second)
				continue
			}

			case "5":
				fmt.Println("Thanks for using Warteg Apps. See You!")
				time.Sleep(2 * time.Second)
				os.Exit(0)
			default:
				fmt.Println("your input is invalid!")
				return
			}
		}
	}