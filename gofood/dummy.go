package gofood

type InMemoryFoodRepo map[string]Food

func (r *InMemoryFoodRepo) GetAllFoodList() []Food {
	list := make([]Food, 0, len(*r))
	for _, value := range *r {
		list = append(list, value)
	}
	return list
}

func (r *InMemoryFoodRepo) Add(food Food) {
	(*r)[food.Name] = food
}

func (r *InMemoryFoodRepo) Delete(name string) {
	delete(*r, name)
}

func (r *InMemoryFoodRepo) Get(name string) (Food, bool) {
	f, ok := (*r)[name]
	return f, ok
}

func (r *InMemoryFoodRepo) Update(food Food) bool {
	_, ok := (*r)[food.Name]
	if !ok {
		return false
	}

	(*r)[food.Name] = food
	return true
}

func NewDummyFoodRepo() *InMemoryFoodRepo {

	r := InMemoryFoodRepo{}

	r.Add(Food{Name: "pancake", Price: 400})
	r.Add(Food{Name: "tortilla", Price: 1400})
	r.Add(Food{Name: "pizza", Price: 1200})

	return &r
}
