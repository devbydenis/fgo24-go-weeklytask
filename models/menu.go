package models

type Menu struct {
	ID    string
	Name  string
	Price int
	Category string
}

var ListMenu []Menu