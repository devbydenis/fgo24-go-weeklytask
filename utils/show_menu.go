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
var ListMenu []m.Menu

func ShowMenu(ListMenu []m.Menu, ch1 chan []m.Menu, wg *sync.WaitGroup) {
	defer wg.Done()
	var input int
	var search string
	
	for {
		if input == 0 {
			fmt.Println("Show All Menu")
			for _, item := range ListMenu {
				fmt.Printf("ID: %s, Price: %d, Category: %s, Name: %s\n", item.ID, item.Price, item.Category, item.Name)
			}
		}
		
		fmt.Println("  ==============================")
		fmt.Println("|| 1. Filter by food category   ||")
		fmt.Println("|| 2. Filter by drink category  ||")
		fmt.Println("|| 3. Filter by cheapest price  ||")
		fmt.Println("|| 4. Sort by ascending name    ||")
		fmt.Println("|| 5. Sort by descending name   ||")
		fmt.Println("|| 6. Search menu               ||")
		fmt.Println("  ==============================")
		
		if input != 6 && input != 7 {
			fmt.Print("\nInput [1-6]: ")
			_, err := fmt.Scan(&input)
			if err != nil {
				fmt.Println("Invalid input ✗")
				return
			}
		}
		if input == 6 {
			// fmt.Print("\033[H\033[2J")
			fmt.Print("\nSearch Menu: ")
			_, err := fmt.Scan(&search)
			if err != nil {
			fmt.Println("Invalid input ✗")
				break
			}
			searchMenu(search, ListMenu, ch1)
			break
		}
		if input == 7 {	// back to menu
			break
		}
		
		filterMenu(input, ListMenu)
	}
	fmt.Println("==========================================")
	fmt.Println("||Choose your favorite food.            ||")
	fmt.Println("||Example: nasi uduk, batagor, es dawet.||")
	fmt.Println("==========================================")
	fmt.Print("Input: ")
	reader := bufio.NewReader(os.Stdin)
    menuNames, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err.Error())
        return
    }
    splitedMN := strings.Split(menuNames, ",")

    result := []m.Menu{}
    for _, name := range splitedMN {
        for _, food := range ListMenu {
            if strings.Contains(strings.ToLower(name), strings.ToLower(food.Name)) {
                result = append(result, food)
            }
        }
    }
    
    if len(result) == 0 {
    	ch1 <- result
        fmt.Println("Menu Not Found")
        time.Sleep(2 * time.Second)
    } else {
	    ch1 <- result
	    fmt.Println("Success adding to cart ✔")
		time.Sleep(2 * time.Second)
    }
    
}




