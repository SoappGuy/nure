package sql_queries

import (
	"time"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type Student struct {
	StudentID          int64     `json:"student_id"`
	Firstname          string    `json:"firstname"`
	Middlename         string    `json:"middlename"`
	Lastname           string    `json:"lastname"`
	Gender             Gender    `json:"gender"`
	Birthday           time.Time `json:"birthday"`
	PersonalFileNumber string    `json:"personal_file_number"`
	Note               *string   `json:"note"`
}

type FoodType string

const (
	Vegetable FoodType = "Овочі"
	Fruit     FoodType = "Фрукти"
	Candy     FoodType = "Цукерки"
	Cookie    FoodType = "Печиво"
)

type Foodstuff struct {
	FoodstuffID int64    `json:"foodstuff_id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Type        FoodType `json:"type"`
}
