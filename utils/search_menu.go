package utils

import (
	"fmt"
	"strings"
	m "weeklytask/models"
)

func searchMenu(keyword string, ListMenu []m.Menu, ch1 chan<- []m.Menu) {
	if keyword == "" {
		fmt.Println("Keyword cannot be empty")
		return
	}
	
	if len(ListMenu) == 0 {
		fmt.Println("Menu is not found")
		return
	}
	
	for _, menu := range ListMenu {
		if strings.Contains(strings.ToLower(menu.Name), strings.ToLower(keyword)) {
			fmt.Printf("ID: %s, Price: %d, Category: %s, Name: %s\n", menu.ID, menu.Price, menu.Category, menu.Name)
		}		
	}
	
}