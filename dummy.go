package gofood

type InMemoryFoodRepo struct {
	list []Food
}

func (r InMemoryFoodRepo) GetAllFoodList() []Food {
	return r.list
}

func NewDummyFoodRepo() InMemoryFoodRepo {

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

	return InMemoryFoodRepo{
		list: mylist,
	}
}
