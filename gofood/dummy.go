package gofood

type InMemoryFoodRepo []Food

func (r *InMemoryFoodRepo) GetAllFoodList() []Food {
	return *r
}

func (r *InMemoryFoodRepo) AddFood(food Food) {
	*r = append(*r, food)
}

func NewDummyFoodRepo() *InMemoryFoodRepo {

	mylist := []Food{
		// Food{
		// 	Name:  "pacal",
		// 	Price: 550,
		// },
		Food{
			Name:  "pancake",
			Price: 400,
		},
		Food{
			Name:  "tortilla",
			Price: 1400,
		},
		Food{
			Name:  "pizza",
			Price: 1200,
		},
	}

	r := InMemoryFoodRepo(mylist)
	return &r
}
