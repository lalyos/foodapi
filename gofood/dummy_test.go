package gofood

import "testing"

func TestConstructor(t *testing.T) {
	repo := NewDummyFoodRepo()

	if len(repo.GetAllFoodList()) < 3 {
		t.Error("constructor should create a lot")
	}
}

func TestAddFood(t *testing.T) {
	repo := NewDummyFoodRepo()
	beforeLen := len(repo.GetAllFoodList())

	f := Food{
		Name:  "pancake",
		Price: 400,
	}
	repo.AddFood(f)

	afterLen := len(repo.GetAllFoodList())
	if (afterLen - beforeLen) != 1 {
		t.Error("Add should increase length by 1")
	}

}
