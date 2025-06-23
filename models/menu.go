package models

import (
	"fmt"
	// "strings"
)

type Menu struct {
	ID    string
	Name  string
	Price int
	Category string
}

var ListMenu []Menu

func (m Menu) RowMenu() {
	fmt.Printf("ID: %s, Price: %d, Category: %s, Name: %s\n", m.ID, m.Price, m.Category, m.Name)
}