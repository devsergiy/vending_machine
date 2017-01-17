package models

// The MoneyStorage describes item for sale
type MoneyStorage struct {
	cashBackAmount, moneyAmount, userDeposit, currentDeposit float32
}

func NewMoneyStorage(cashBackAmount float32) *MoneyStorage {
	return &MoneyStorage{cashBackAmount, 0, 0, 0}
}

// AddMoney allows user to deposit money
func (m *MoneyStorage) AddMoney(amount float32) {

}

func (m *MoneyStorage) DeductMoney(amount float32) error {
	return nil
}

func (m *MoneyStorage) MakeCashBack() (float32, error) {
	return 0, nil
}
