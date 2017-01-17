package models

import "errors"

// The Dispenser describes item for sale
type Dispenser struct {
	products map[int]*StorageItem
	*MoneyStorage
}

var DispenserInstance = FillDispenser(NewDispenser())

func NewDispenser() *Dispenser {
	return &Dispenser{make(map[int]*StorageItem), NewMoneyStorage()}
}

func (d *Dispenser) RequestProduct(productID int) (float32, error) {
	currentDeposit := d.MoneyStorage.DepositAmount()

	storage, ok := d.products[productID]
	if !ok {
		return currentDeposit, errors.New("No such product")
	}

	if storage.Amount == 0 {
		return currentDeposit, errors.New("Not enough products")
	}

	amount, err := d.MoneyStorage.UseMoney(storage.Product.Price)

	if err != nil {
		return currentDeposit, err
	}

	return amount, nil
}

func (d *Dispenser) AddProduct(product *Product, amount int) {
	storage, ok := d.products[product.ID]

	if !ok {
		storage = &StorageItem{Product: product, Amount: 0}
	}

	storage.Amount += amount
	d.products[product.ID] = storage
}

func (d *Dispenser) ProductsList() []*StorageItem {
	results := make([]*StorageItem, 0)

	for _, storageItem := range d.products {
		results = append(results, storageItem)
	}

	return results
}
