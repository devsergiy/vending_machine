package models

// The Dispenser describes item for sale
type Dispenser struct {
	products map[int32]*StorageItem
	*MoneyStorage
}

func NewDispenser() *Dispenser {
	return &Dispenser{make(map[int32]*StorageItem), NewMoneyStorage(100)}
}

func (m *Dispenser) RequestProduct(productID int) (product *Product, err error) {
	return nil, nil
}

func (d *Dispenser) AddProduct(product *Product, amount int) {

}
