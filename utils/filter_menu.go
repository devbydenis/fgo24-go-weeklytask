package utils

import (
	"fmt"
	"sort"
	"strings"
	"time"
	m "weeklytask/models"
)


func filterByCategory(category string, ListMenu []m.Menu) {
	for _, menu := range ListMenu {
		if strings.ToLower(menu.Category) == category {
			fmt.Printf("ID: %s, Price: %d, Category: %s, Name: %s\n", menu.ID, menu.Price, menu.Category, menu.Name)
		}
	}
}

func filterByCheapestPrice(ListMenu []m.Menu) {
	sort.Slice(ListMenu, func(i, j int) bool {
		return ListMenu[i].Price < ListMenu[j].Price
	})
	for _, menu := range ListMenu {
		fmt.Printf("ID: %s, Price: %d, Category: %s, Name: %s\n", menu.ID, menu.Price, menu.Category, menu.Name)
	}
}

func filterByName(mode string, ListMenu []m.Menu) {
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

func filterMenu(filter int, ListMenu []m.Menu){
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
			fmt.Println("Invalid input ✗. \nPlease input number between 1-6")
			time.Sleep(time.Second)
		}
	}