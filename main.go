package main

import (
	// "fmt"
	"weeklytask/data"
	"weeklytask/layout"
	"weeklytask/models"
)

func main() {
	ci := data.CurrentInput
	fl := data.FoodList
	
	layout.Main_Layout(&ci)
	
	switch ci {
	case "1":
		models.ShowFoodList(fl)
	case "output":
		// layout.Output_Layout(&ci)
	default:
		// layout.Error_Layout(&ci)
	}
	
}	