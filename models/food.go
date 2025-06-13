package models

import "fmt"

type Food struct {
	ID    string
	Name  string
	Price int
}

type FoodList struct {
	Foods []Food
}

func ShowFoodList(FoodList []Food) {
	for _, food := range FoodList {
		fmt.Printf("ID: %s, Name: %s, Price: %d\n", food.ID, food.Name, food.Price)
	}
}
