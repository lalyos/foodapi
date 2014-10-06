package gofood

import "testing"

func TestConstructor(t *testing.T) {
	repo := NewDummyFoodRepo()

	if len(repo.GetAllFoodList()) < 3 {
		t.Error("constructor should create a lot")
	}
}

func getDummyFood() Food {
	return Food{
		Name:  "newDummyFood",
		Price: 400,
	}
}

func TestAddFood(t *testing.T) {
	repo := NewDummyFoodRepo()
	beforeLen := len(repo.GetAllFoodList())
	repo.AddFood(getDummyFood())

	afterLen := len(repo.GetAllFoodList())
	if (afterLen - beforeLen) != 1 {
		t.Error("Add should increase length by 1")
	}
}

func TestGetFood(t *testing.T) {
	repo := NewDummyFoodRepo()
	f, ok := repo.GetFood("pizza")

	if !ok {
		t.Error("GetFood should return a valid object")
	}

	if f.Price != 1200 {
		t.Error("GetFood should return a food with known price")
	}
}

func TestDeleteFood(t *testing.T) {
	repo := NewDummyFoodRepo()
	beforeLen := len(repo.GetAllFoodList())
	repo.DeleteFood("pizza")

	afterLen := len(repo.GetAllFoodList())
	if (beforeLen - afterLen) != 1 {
		t.Error("Delete should decrease length by 1")
	}
}

func TestAddFoodTwice(t *testing.T) {
	repo := NewDummyFoodRepo()
	beforeLen := len(repo.GetAllFoodList())

	repo.AddFood(getDummyFood())
	afterAddLen := len(repo.GetAllFoodList())

	if (afterAddLen - beforeLen) != 1 {
		t.Error("Add should increase length by 1")
	}

	repo.AddFood(getDummyFood())
	afterSecondAddLen := len(repo.GetAllFoodList())

	if (afterSecondAddLen - afterAddLen) != 0 {
		t.Error("Second Add should NOT increase the length")
	}
}
