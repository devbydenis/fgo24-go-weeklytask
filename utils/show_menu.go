package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	m "weeklytask/models"
)

func ShowMenu(ListMenu m.ListMenu, ch1 chan []m.Menu, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("--------------------------------------------------")
	fmt.Println("Showing food list...")
	
	for _, menu := range ListMenu {
		fmt.Printf("ID: %s, Price: %d, Category: %s, Name: %s\n", menu.ID, menu.Price, menu.Category, menu.Name)
	}
	
	fmt.Println("--------------------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Choose your favorite food: ")
    foodName, _ := reader.ReadString('\n')
    splitFN := strings.Split(foodName, ",")

    result := []m.Menu{}
    for _, name := range splitFN {
        for _, food := range ListMenu {
            if strings.Contains(strings.ToLower(name), strings.ToLower(food.Name)) {
                result = append(result, food)
            }
        }
    }
    
    ch1 <- result
    
	fmt.Println("success adding to cart")
	time.Sleep(1 * time.Second)
}

func ShowCart(cart []m.Menu, ci *string) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("my cart: ", cart)
	fmt.Println("back to menu      [0]")
	fmt.Println("checkout my chart [1]")
	
	fmt.Print("Input: ")
	_, err := fmt.Scanln(*ci)
	if err != nil {
		fmt.Println("Invalid input")
	}
}