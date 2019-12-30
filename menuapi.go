package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Base struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Size struct {
	Id    string  `json:"id"`
	Size  string  `json:"size"`
	Price float64 `json:"price"`
}

type Topping struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Menu struct {
	Bases    []Base    `json:"base"`
	Sizes    []Size    `json:"size"`
	Toppings []Topping `json:"topping"`
}

type menu Menu

func setMenu(w http.ResponseWriter, r *http.Request) {
	menu := Menu{
		[]Base{
			Base{Id: 1, Name: "Hand Tossed", Price: 0},
			Base{Id: 2, Name: "Thin Crust", Price: 20},
			Base{Id: 3, Name: "Cheese Burst", Price: 80},
			Base{Id: 4, Name: "Pan Pizza", Price: 40},
		},
		[]Size{
			Size{Id: "S", Size: "Small", Price: 230},
			Size{Id: "M", Size: "Medium", Price: 350},
			Size{Id: "L", Size: "Large", Price: 500},
		},
		[]Topping{
			Topping{Id: 1, Name: "Jalepino", Price: 20},
			Topping{Id: 2, Name: "Paprica", Price: 30},
			Topping{Id: 3, Name: "Corn", Price: 40},
			Topping{Id: 4, Name: "Capsicum", Price: 20},
			Topping{Id: 5, Name: "Paneer", Price: 40},
			Topping{Id: 6, Name: "Olives", Price: 30},
			Topping{Id: 7, Name: "Mushroom", Price: 30},
			Topping{Id: 8, Name: "BBQ Chicken", Price: 70},
			Topping{Id: 9, Name: "Chicken Sausage", Price: 80},
			Topping{Id: 10, Name: "Chicken Tikka", Price: 60},
		},
	}

	json.NewEncoder(w).Encode(menu)
}

func handleRequests() {
	http.HandleFunc("/menu", setMenu)
	log.Fatal(http.ListenAndServe(":9191", nil))
}

func main() {
	handleRequests()
}
