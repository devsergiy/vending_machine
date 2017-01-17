package models

func FillDispenser(dispenser *Dispenser) *Dispenser {
	products := []*StorageItem{
		&StorageItem{&Product{1, "Coca-Cola", 2.25, "drink"}, 5},
		&StorageItem{&Product{2, "Nuts", 5.55, "chocolate"}, 10},
		&StorageItem{&Product{3, "Pepsi", 2.10, "drink"}, 2},
	}

	for _, item := range products {
		dispenser.AddProduct(item.Product, item.Amount)
	}

	return dispenser
}
