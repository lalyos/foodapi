package gofood

import "fmt"

type Food struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (f Food) String() string {
	return fmt.Sprintf("[name: %s, price:%d]", f.Name, f.Price)
}

type FoodRepo interface {
	Get(name string) (Food, bool)
	GetAllFoodList() []Food
	Add(food Food)
	Delete(name string)
	Update(food Food) bool
}
