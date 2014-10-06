package gofood

type InMemoryFoodRepo map[string]Food

func (r *InMemoryFoodRepo) GetAllFoodList() []Food {
	list := make([]Food, 0, len(*r))
	for _, value := range *r {
		list = append(list, value)
	}
	return list
}

func (r *InMemoryFoodRepo) AddFood(food Food) {
	(*r)[food.Name] = food
}

func (r *InMemoryFoodRepo) DeleteFood(name string) {
	delete(*r, name)
}

func (r *InMemoryFoodRepo) GetFood(name string) (Food, bool) {
	f, ok := (*r)[name]
	return f, ok
}

func (r *InMemoryFoodRepo) UpdateFood(food Food) bool {
	_, ok := (*r)[food.Name]
	if !ok {
		return false
	}

	(*r)[food.Name] = food
	return true
}

func NewDummyFoodRepo() *InMemoryFoodRepo {

	r := InMemoryFoodRepo{}

	r.AddFood(Food{Name: "pancake", Price: 400})
	r.AddFood(Food{Name: "tortilla", Price: 1400})
	r.AddFood(Food{Name: "pizza", Price: 1200})

	return &r
}
