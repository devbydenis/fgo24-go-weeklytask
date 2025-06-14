package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
	m "weeklytask/models"
)


func filterByCategory(category string, ListMenu m.ListMenu) {
	for _, menu := range ListMenu {
		if strings.ToLower(menu.Category) == category {
			fmt.Printf("ID: %s, Price: %d, Category: %s, Name: %s\n", menu.ID, menu.Price, menu.Category, menu.Name)
		}
	}
}
func filterByCheapestPrice(ListMenu m.ListMenu) {
	sort.Slice(ListMenu, func(i, j int) bool {
		return ListMenu[i].Price < ListMenu[j].Price
	})
	for _, menu := range ListMenu {
		fmt.Printf("ID: %s, Price: %d, Category: %s, Name: %s\n", menu.ID, menu.Price, menu.Category, menu.Name)
	}
}
func filterByName(mode string, ListMenu m.ListMenu) {
	if mode == "ascending" {
		sort.Slice(ListMenu, func(i, j int) bool {
			return ListMenu[i].Name < ListMenu[j].Name
		})
	} 
	
	if mode == "descending" {
		sort.Slice(ListMenu, func(i, j int) bool {
			return ListMenu[i].Name > ListMenu[j].Name
		})
	}
	
	for _, menu := range ListMenu {
		fmt.Printf("ID: %s, Price: %d, Category: %s, Name: %s\n", menu.ID, menu.Price, menu.Category, menu.Name)
	}
}
func searchMenu(keyword string, ListMenu m.ListMenu, ch1 chan<- []m.Menu) {
	// var result []m.Menu
	for _, menu := range ListMenu {
		if strings.Contains(strings.ToLower(menu.Name), strings.ToLower(keyword)) {
			fmt.Printf("ID: %s, Price: %d, Category: %s, Name: %s\n", menu.ID, menu.Price, menu.Category, menu.Name)
			// result = append(result, menu)
		}
	}
	// ch1 <- result
	// close(ch1)
}

func filterMenu(filter int, ListMenu m.ListMenu){
	switch filter {
		case 1:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Filter by Food Category")
			filterByCategory("food", ListMenu)
		case 2:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Filter by Drink Category")
			filterByCategory("drink", ListMenu)
		case 3:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Sorted by cheapest price")
			filterByCheapestPrice(ListMenu)
		case 4:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Sorted by ascending name")
			filterByName("ascending", ListMenu)
		case 5:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Sorted by descending name")
			filterByName("descending", ListMenu)
		case 6:
			return
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Invalid input. Please input number between 1-6")
			time.Sleep(time.Second)
		}
	}


func ShowMenu(ListMenu m.ListMenu, ch1 chan []m.Menu, wg *sync.WaitGroup) {
	defer wg.Done()
	var input int
	var search string
	
	for {
		fmt.Println("  ==============================")
		fmt.Println("|| 1. Filter by food category   ||")
		fmt.Println("|| 2. Filter by drink category  ||")
		fmt.Println("|| 3. Filter by cheapest price  ||")
		fmt.Println("|| 4. Filter by ascending name  ||")
		fmt.Println("|| 5. Filter by descending name ||")
		fmt.Println("|| 6. Search menu               ||")
		fmt.Println("|| 7. Back to main menu         ||")
		fmt.Println("  ==============================")
		
		if input != 6 && input != 7 {
			fmt.Print("\nInput: ")
			_, err := fmt.Scan(&input)
			if err != nil {
				fmt.Println("Invalid input")
				return
			}
		}
		if input == 6 {
			// fmt.Print("\033[H\033[2J")
			fmt.Println("search dipannggil")
			fmt.Print("\nSearch Menu: ")
			_, err := fmt.Scan(&search)
			if err != nil {
				fmt.Println("Invalid input")
				break
			}
			searchMenu(search, ListMenu, ch1)
			// return
			// continue
			break
		}
		if input == 7 {	// back to menu
			break
		}
		
		filterMenu(input, ListMenu)
	}
	
	fmt.Println("Choose your favorite food, \nexample: nasi uduk, batagor, es dawet")
	fmt.Print("Input: ")
	reader := bufio.NewReader(os.Stdin)
    menuNames, _ := reader.ReadString('\n')
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
        fmt.Println("No food found")
        time.Sleep(1 * time.Second)
    } else {
	    ch1 <- result
	    fmt.Println("success adding to cart")
		time.Sleep(1 * time.Second)
    }
    
}