package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	m "weeklytask/models"
)

func ShowMenu(ListMenu m.ListMenu, ch1 chan []m.Menu, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("--------------------------------------------------")
	fmt.Println("Showing food list...")
	
	for _, food := range ListMenu {
		fmt.Printf("ID: %s, Price: %d, Category: %s, Name: %s\n", food.ID, food.Price, food.Category, food.Name)
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
}